package zypper

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("zypper", "search", "--installed-only")
		buffer, err := cmd.Output()

		if err == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

			if len(lines) > 4 {

				check1 := strings.TrimSpace(lines[0])
				check2 := strings.TrimSpace(lines[1])
				// Empty Line in between
				check3 := strings.TrimSpace(lines[3])
				check4 := strings.TrimSpace(lines[4])

				if strings.Contains(check1, "Loading repository data...") &&
					strings.Contains(check2, "Reading installed packages...") &&
					strings.Contains(check3, "|") &&
					strings.HasPrefix(check4, "---") && strings.HasSuffix(check4, "---") {

					for l := 5; l < len(lines); l++ {

						line := strings.TrimSpace(lines[l])

						if line != "" {

							if strings.Contains(line, "---") {

								// Do Nothing

							} else if strings.Contains(line, "|") {

								chunks := strings.Split(line, "|")

								if len(chunks) == 4 {

									status := strings.TrimSpace(chunks[0])
									name := strings.TrimSpace(chunks[1])
									typ := strings.TrimSpace(chunks[3])

									if typ == "package" {

										if status == "i" || status == "i+" {

											pkg := CollectPackage(name)

											if pkg.Name != "" && pkg.Version.IsValid() {
												collected = append(collected, pkg)
											}

										}

									}

								}

							}

						}

					}

				}

			}

		}

	}

	return collected

}
