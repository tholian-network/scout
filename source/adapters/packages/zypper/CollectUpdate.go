package zypper

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdate(name string) structs.Update {

	var result structs.Update = structs.NewUpdate("zypper")

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("zypper", "info", name)
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			block := strings.TrimSpace(string(buffer))

			// "Information for package xyz:"
			// "-----------------------------\n"
			if strings.Contains(block, "----\n") {
				block = strings.TrimSpace(block[strings.Index(block, "----\n")+5:])
			}

			if block != "" {
				ParseUpdate(block, &result)
			}

		}

	}

	return result

}
