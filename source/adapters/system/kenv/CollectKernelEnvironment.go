package kenv

import "os/exec"
import "strings"

func CollectKernelEnvironment() map[string]string {

	collected := make(map[string]string)

	if SUPPORTED == true {

		cmd := exec.Command("kenv", "-q")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			lines := strings.Split(strings.ReplaceAll(string(buffer), "\r\n", "\n"), "\n")

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}

				index := strings.Index(line, "=")

				if index > 0 {

					key := strings.TrimSpace(line[0:index])
					value := strings.TrimSpace(line[index+1:])

					if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") && len(value) >= 2 {
						value = value[1 : len(value)-1]
					}

					if key != "" && value != "" {
						collected[key] = value
					}

				}

			}

		}

	}

	return collected

}
