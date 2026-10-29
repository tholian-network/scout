package kldstat

import "os/exec"
import "sort"
import "strconv"
import "strings"

func CollectKernelModules() []string {

	var collected []string

	if SUPPORTED == true {

		cmd := exec.Command("kldstat", "-v")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			contains := false
			lines := strings.Split(strings.ReplaceAll(string(buffer), "\r\n", "\n"), "\n")

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if line == "" {
					continue
				}

				if strings.Contains(line, "Contains modules:") {
					contains = true
					continue
				}

				fields := strings.Fields(line)

				if len(fields) == 0 || fields[0] == "Id" {
					continue
				}

				name := ""

				if contains == true {

					if len(fields) >= 2 {
						name = fields[len(fields)-1]
					}

				} else if _, err := strconv.ParseUint(fields[0], 10, 64); err == nil && len(fields) >= 5 {
					name = fields[4]
				}

				name = strings.TrimSuffix(strings.TrimSpace(name), ".ko")

				if name != "" && name != "kernel" {

					found := false

					for c := 0; c < len(collected); c++ {
						if collected[c] == name {
							found = true
							break
						}
					}

					if found == false {
						collected = append(collected, name)
					}

				}

			}

			sort.Strings(collected)

		}

	}

	return collected

}
