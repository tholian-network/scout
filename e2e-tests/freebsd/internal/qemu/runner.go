package qemu

import "fmt"
import "os"
import "os/exec"
import "path/filepath"
import "time"

type Config struct {
	Root         string
	TestBinary   string
	HelperBinary string
	SetupScript  string
	ResetScript  string
	Keep         bool
	Diagnose     bool
	TestRun      string
	Timeout      time.Duration
	TestTimeout  time.Duration
}

func Run(config Config) error {

	if config.Timeout == 0 {
		config.Timeout = 360 * time.Second
	}

	if config.TestTimeout == 0 {
		config.TestTimeout = 10 * time.Minute
	}

	if config.Root == "" {
		config.Root = "."
	}

	root, err := filepath.Abs(config.Root)

	if err != nil {
		return err
	}

	if config.SetupScript == "" {
		config.SetupScript = filepath.Join(root, "freebsd", "setup.sh")
	}

	if config.ResetScript == "" {
		config.ResetScript = filepath.Join(root, "freebsd", "reset.sh")
	}

	buildDir := filepath.Join(filepath.Dir(root), "build", "tests", "freebsd")

	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return err
	}

	if config.TestBinary == "" {
		config.TestBinary = filepath.Join(buildDir, "freebsd-amd64.test")
	}

	if config.HelperBinary == "" {
		config.HelperBinary = filepath.Join(buildDir, "e2e-scout-pkg")
	}

	for _, tool := range []string{"qemu-system-x86_64", "qemu-img", "ssh", "scp", "ssh-keygen", "xz"} {
		if _, err := exec.LookPath(tool); err != nil {
			return fmt.Errorf("missing required tool: %s", tool)
		}
	}

	if err := buildTest(root, config.TestBinary); err != nil {
		return err
	}

	if err := buildHelper(root, config.HelperBinary); err != nil {
		return err
	}

	image, err := EnsureImage(filepath.Join(buildDir, "cache"))

	if err != nil {
		return err
	}

	runDir := filepath.Join(buildDir, "runs", time.Now().Format("20060102-150405"))

	if err := os.MkdirAll(runDir, 0755); err != nil {
		return err
	}

	overlay := filepath.Join(runDir, "disk.qcow2")

	if err := CreateOverlay(image, overlay); err != nil {
		return err
	}

	if config.Keep == false {
		defer os.Remove(overlay)
	}

	privateKey, publicKey, err := GenerateSSHKey(runDir)

	if err != nil {
		return err
	}

	seed, err := WriteSeed(runDir, publicKey)

	if err != nil {
		return err
	}

	port, err := FreePort()

	if err != nil {
		return err
	}

	serial := filepath.Join(runDir, "serial.log")
	qemuLog := filepath.Join(runDir, "qemu.log")

	logFile, err := os.Create(qemuLog)

	if err != nil {
		return err
	}

	defer logFile.Close()

	arguments := []string{
		"-machine", "q35,accel=kvm:tcg",
		"-m", "8192",
		"-smp", "2",
		"-drive", "file=" + overlay + ",if=virtio,format=qcow2",
		"-drive", "file=" + seed + ",if=virtio,format=raw,readonly=on",
		"-netdev", "user,id=net0,hostfwd=tcp:127.0.0.1:" + port + "-:22",
		"-device", "virtio-net-pci,netdev=net0",
		"-display", "none",
		"-serial", "file:" + serial,
		"-monitor", "none",
	}

	cmd := exec.Command("qemu-system-x86_64", arguments...)
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	fmt.Println("Booting FreeBSD 15.1 ...")

	if err := cmd.Start(); err != nil {
		return err
	}

	finished := false

	defer func() {
		if finished == false && cmd.Process != nil {
			cmd.Process.Kill()
			cmd.Wait()
		}
	}()

	client, err := WaitForSSH(port, privateKey, config.Timeout)

	if err != nil {
		return fmt.Errorf("%w (serial log: %s, qemu log: %s)", err, serial, qemuLog)
	}

	if err := client.CopyAs(config.TestBinary, "/tmp/e2e"); err != nil {
		return err
	}

	if err := client.CopyAs(config.HelperBinary, "/tmp/e2e-scout-pkg"); err != nil {
		return err
	}

	if err := client.CopyAs(config.SetupScript, "/tmp/e2e-setup.sh"); err != nil {
		return err
	}

	if err := client.CopyAs(config.ResetScript, "/tmp/e2e-reset.sh"); err != nil {
		return err
	}

	if err := client.Run("chmod +x /tmp/e2e /tmp/e2e-scout-pkg /tmp/e2e-setup.sh /tmp/e2e-reset.sh"); err != nil {
		return err
	}

	if err := client.Run("sh /tmp/e2e-setup.sh"); err != nil {
		return fmt.Errorf("VM setup failed: %w (serial log: %s)", err, serial)
	}

	if config.Diagnose == true {

		commands := []string{
			"uname -a",
			"ps -ax -o pid,uid,comm",
			"procstat -a -h -b",
			"procstat -a -h -c",
			"procstat -a -h -e",
			"procstat -a -h -f",
			"procstat -a -h -v",
			"sockstat -n -4 -6",
			"pkg info -l e2e-scout-pkg",
			"pkg query -a '%n\t%v\t%q'",
		}

		for _, command := range commands {
			output, err := client.Output(command)
			fmt.Printf("\n$ %s\n%s\n", command, string(output))
			if err != nil {
				fmt.Printf("command error: %s\n", err.Error())
			}
		}

		client.Run("shutdown -p now")
		finished = true
		return nil

	}

	command := fmt.Sprintf("env E2E_PACKAGE=e2e-scout-pkg E2E_SERVICE=e2e-scout-pkg E2E_OLD_VERSION=1.0 E2E_PACKAGE_MANAGER=pkg E2E_RESET=/var/cache/e2e/reset.sh /tmp/e2e -test.v -test.timeout %s", config.TestTimeout)

	if config.TestRun != "" {
		command += " -test.run '" + config.TestRun + "'"
	}

	test := client.Command(command)
	test.Stdout = os.Stdout
	test.Stderr = os.Stderr

	testErr := test.Run()

	client.Run("shutdown -p now")

	done := make(chan error, 1)

	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		<-done
	}

	finished = true

	if testErr != nil {
		return fmt.Errorf("FreeBSD tests failed: %w (serial log: %s)", testErr, serial)
	}

	if config.Keep == false {
		fmt.Println("Serial log:", serial)
	}

	return nil

}

func buildTest(root string, output string) error {

	cmd := exec.Command("go", "test", "-tags", "scout", "-c", "-o", output, "./freebsd")
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOOS=freebsd", "GOARCH=amd64", "CGO_ENABLED=0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func buildHelper(root string, output string) error {

	cmd := exec.Command("go", "build", "-o", output, "./freebsd/service")
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOOS=freebsd", "GOARCH=amd64", "CGO_ENABLED=0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}
