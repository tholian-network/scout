package zypper

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdates() []structs.Update {

	var collected []structs.Update

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("zypper", "list-updates")
		buffer, err := cmd.Output()

		if err == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

			if len(lines) > 4 {

				check1 := strings.TrimSpace(lines[0])
				check2 := strings.TrimSpace(lines[1])
				check3 := strings.TrimSpace(lines[2])
				check4 := strings.TrimSpace(lines[3])

				if strings.Contains(check1, "Loading repository data...") &&
					strings.Contains(check2, "Reading installed packages...") &&
					strings.Contains(check3, "|") &&
					strings.HasPrefix(check4, "--") && strings.HasSuffix(check4, "---") {

					for l := 4; l < len(lines); l++ {

						line := strings.TrimSpace(lines[l])

						if line != "" {

							if strings.Contains(line, "---") {

								// Do Nothing

							} else if strings.Contains(line, "|") {

								chunks := strings.Split(line, "|")

								if len(chunks) == 6 {

									status := strings.TrimSpace(chunks[0])
									// repo := strings.TrimSpace(chunks[1])
									name := strings.TrimSpace(chunks[2])
									// current_version := strings.TrimSpace(chunks[3])
									version := strings.TrimSpace(chunks[4])
									architecture := strings.TrimSpace(chunks[5])

									if status == "v" {

										update := structs.NewUpdate("zypper")
										update.SetName(name)
										update.SetVersion(version)
										update.SetArchitecture(architecture)

										if update.Name != "" && update.Version.IsValid() {
											collected = append(collected, update)
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
