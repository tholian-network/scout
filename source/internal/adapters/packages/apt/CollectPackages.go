package apt

import "tholian-core/structs"
import "os/exec"
import "strings"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true {

		cmd := exec.Command("apt", "list")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")
			check := strings.TrimSpace(lines[0])

			if strings.Contains(check, "Listing...") {

				for l := 1; l < len(lines); l++ {

					line := strings.TrimSpace(lines[l])
					flags := make(map[string]string)

					if strings.Contains(line, "[") && strings.HasSuffix(line, "]") {

						tmp := strings.Split(strings.TrimSpace(line[strings.Index(line, "[")+1:strings.LastIndex(line, "]")]), ",")

						for t := 0; t < len(tmp); t++ {

							if strings.Contains(tmp[t], ": ") {

								key := strings.TrimSpace(tmp[t][0:strings.Index(tmp[t], ": ")])
								val := strings.TrimSpace(tmp[t][strings.Index(tmp[t], ": ")+2:])

								flags[key] = val

							} else {

								key := strings.TrimSpace(tmp[t])

								flags[key] = ""

							}

						}

						line = strings.TrimSpace(line[0:strings.Index(line, "[")])

					}

					chunks := strings.Split(line, " ")

					if len(chunks) == 3 {

						_, is_installed := flags["installed"]
						upgradable_version, is_upgradable := flags["upgradable from"]

						if is_installed {

							name := strings.TrimSpace(chunks[0][0:strings.Index(chunks[0], "/")])
							version := strings.TrimSpace(chunks[1])
							pkg := CollectPackage(name, version)

							if pkg.Name != "" && pkg.Version.IsValid() {
								collected = append(collected, pkg)
							}

						} else if is_upgradable {

							name := strings.TrimSpace(chunks[0][0:strings.Index(chunks[0], "/")])
							version := upgradable_version
							pkg := CollectPackage(name, version)

							if pkg.Name != "" && pkg.Version.IsValid() {
								collected = append(collected, pkg)
							}

						}

					}

				}

			}

		}

		for c := 0; c < len(collected); c++ {
			collected[c].ResolveDependencies(collected)
		}

	}

	return collected

}
