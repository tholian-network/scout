package tdnf

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdates() []structs.Update {

	var collected []structs.Update

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("tdnf", "--noplugins", "list", "updates")
		buffer, err := cmd.Output()

		if err == nil {
			collected = parseUpdates(string(buffer), "")
		}

	}

	return collected

}

func parseUpdates(buffer string, name string) []structs.Update {

	var result []structs.Update

	index := make(map[string]int)

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	for l := 0; l < len(lines); l++ {

		chunks := strings.Fields(strings.TrimSpace(lines[l]))

		if len(chunks) >= 2 {

			namearch := chunks[0]
			version := chunks[1]
			dot := strings.LastIndex(namearch, ".")

			if dot > 0 {

				package_name := strings.TrimSpace(namearch[0:dot])
				architecture := strings.TrimSpace(namearch[dot+1:])

				if name == "" || package_name == name {

					update := structs.NewUpdate("tdnf")
					update.SetName(package_name)
					update.SetVersion(version)
					update.SetArchitecture(architecture)

					if update.Name != "" && update.Version.IsValid() {

						key := update.Name + "." + update.Architecture.String()

						if i, ok := index[key]; ok {

							if update.Version.IsAfter(result[i].Version) {
								result[i] = update
							}

						} else {

							index[key] = len(result)
							result = append(result, update)

						}

					}

				}

			}

		}

	}

	return result

}
