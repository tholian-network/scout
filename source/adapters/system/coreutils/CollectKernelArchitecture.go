package coreutils

import "os/exec"
import "strings"

func CollectKernelArchitecture() string {

	var collected string

	if SUPPORTED == true {

		cmd := exec.Command("uname", "-m")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {
			collected = strings.TrimSpace(string(buffer))
		}

	}

	return collected

}
