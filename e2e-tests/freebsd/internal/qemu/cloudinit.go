package qemu

import "fmt"
import "os"
import "os/exec"
import "path/filepath"

func WriteSeed(dir string, publicKey string) (string, error) {

	meta := "instance-id: scout-freebsd-15.1\nlocal-hostname: scout-freebsd\n"

	user := fmt.Sprintf(`#cloud-config
disable_root: false
ssh_pwauth: false
users:
  - default
  - name: freebsd
    groups: wheel
    shell: /bin/sh
    ssh_authorized_keys:
      - %s
runcmd:
  - mkdir -p /root/.ssh
  - sh -c 'printf "%%s\n" "%s" > /root/.ssh/authorized_keys'
  - chmod 700 /root/.ssh
  - chmod 600 /root/.ssh/authorized_keys
  - sed -i '' 's/^#*PermitRootLogin.*/PermitRootLogin prohibit-password/' /etc/ssh/sshd_config
  - sh -c 'grep -q "^PermitRootLogin prohibit-password" /etc/ssh/sshd_config || echo "PermitRootLogin prohibit-password" >> /etc/ssh/sshd_config'
  - sysrc sshd_enable=YES
  - service sshd restart
`, publicKey, publicKey)

	meta_file := filepath.Join(dir, "meta-data")
	user_file := filepath.Join(dir, "user-data")

	if err := os.WriteFile(meta_file, []byte(meta), 0644); err != nil {
		return "", err
	}

	if err := os.WriteFile(user_file, []byte(user), 0644); err != nil {
		return "", err
	}

	seed := filepath.Join(dir, "seed.iso")
	tool, err := findTool("genisoimage", "mkisofs", "xorriso")

	if err != nil {
		return "", err
	}

	var cmd *exec.Cmd

	if tool == "xorriso" {
		cmd = exec.Command(tool, "-as", "mkisofs", "-output", seed, "-volid", "cidata", "-joliet", "-rock", meta_file, user_file)
	} else {
		cmd = exec.Command(tool, "-output", seed, "-volid", "cidata", "-joliet", "-rock", meta_file, user_file)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("create seed ISO: %w", err)
	}

	return seed, nil

}

func findTool(names ...string) (string, error) {

	for _, name := range names {
		if path, err := exec.LookPath(name); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("missing required tool: %s", names[0])

}
