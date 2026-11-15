package coreutils

import "os/exec"
import "strings"

func CollectKernel() string {

	var collected string

	if SUPPORTED == true {

		cmd := exec.Command("uname")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {
			collected = strings.TrimSpace(string(buffer))
		}

	}

	return collected

}
