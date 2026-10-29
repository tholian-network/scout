package main

import "flag"
import "fmt"
import "os"
import "time"
import "tholian-scout-e2e/freebsd/internal/qemu"

func main() {

	root := flag.String("root", ".", "Path to the e2e-tests module")
	test_binary := flag.String("test-binary", "", "Prebuilt FreeBSD test binary")
	helper_binary := flag.String("helper-binary", "", "Prebuilt FreeBSD helper service binary")
	setup := flag.String("setup", "", "VM setup script")
	reset := flag.String("reset", "", "VM reset script")
	keep := flag.Bool("keep", false, "Keep the VM run directory")
	diagnose := flag.Bool("diagnose", false, "Print FreeBSD collector command output and exit")
	test_run := flag.String("test-run", "", "Only run matching Go tests")
	timeout := flag.Duration("timeout", 360*time.Second, "SSH boot timeout")
	test_timeout := flag.Duration("test-timeout", 10*time.Minute, "Go test timeout inside the VM")

	flag.Parse()

	err := qemu.Run(qemu.Config{
		Root:         *root,
		TestBinary:   *test_binary,
		HelperBinary: *helper_binary,
		SetupScript:  *setup,
		ResetScript:  *reset,
		Keep:         *keep,
		Diagnose:     *diagnose,
		TestRun:      *test_run,
		Timeout:      *timeout,
		TestTimeout:  *test_timeout,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

}
