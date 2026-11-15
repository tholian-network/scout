package apk

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdates() []structs.Update {

	var collected []structs.Update

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("apk", "list", "--upgradable")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

			for l := 0; l < len(lines); l++ {

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

				if strings.Contains(line, "(") && strings.HasSuffix(line, ")") {
					// Remove License
					line = strings.TrimSpace(line[0:strings.Index(line, "(")])
				}

				chunks := strings.Split(line, " ")

				if len(chunks) == 3 {

					_, is_upgradable := flags["upgradable from"]

					if is_upgradable {

						name, version := toNameAndVersion(chunks[0])
						architecture := strings.TrimSpace(chunks[1])

						update := structs.NewUpdate("apk")
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

	return collected

}
