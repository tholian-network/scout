package proc

import "os"
import "sort"
import "strings"

func CollectKernelModules() []string {

	var collected []string

	if SUPPORTED == true {

		buffer, err := os.ReadFile("/proc/modules")

		if err == nil && len(buffer) > 0 {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

			for l := 0; l < len(lines); l++ {

				line := splitLine(strings.TrimSpace(lines[l]), " ")

				if len(line) == 6 {

					var name = line[0]
					// var references = line[2]
					var dependants []string

					if line[3] != "-" {
						dependants = splitLine(line[3], ",")
					}

					if name != "" && len(dependants) > 0 {

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

			}

			sort.Strings(collected)

		}

	}

	return collected

}
