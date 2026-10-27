package pacman

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdates() []structs.Update {

	var collected []structs.Update

	if SUPPORTED == true && OPTIMIZED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		update_index := make(map[string]bool, 0)
		cmd1 := exec.Command("pacman", "-Qu", "--noconfirm")
		buffer1, err1 := cmd1.Output()

		if err1 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer1)), "\n")

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if strings.HasSuffix(line, "[ignored]") {
					line = strings.TrimSpace(line[0 : len(line)-9])
				}

				if strings.Contains(line, " ") && strings.Contains(line, " -> ") {

					// "package 1.2.3 -> 1.2.4"

					name := line[0:strings.Index(line, " ")]
					update_index[name] = true

				}

			}

		}

		cmd := exec.Command("pacman", "-Si", "--noconfirm")
		buffer, err := cmd.Output()

		if err == nil {

			blocks := strings.Split("\n\n"+strings.TrimSpace(string(buffer)), "\n\nRepository")

			for b := 0; b < len(blocks); b++ {

				block := strings.TrimSpace(blocks[b])

				if block != "" {

					update := structs.NewUpdate("pacman")
					ParseUpdate("Repository "+block, &update)

					if update.Name != "" && update.Version.IsValid() {

						if strings.HasPrefix(update.Name, "lib32-") {
							update.SetArchitecture("x86")
						}

						_, ok := update_index[update.Name]

						if ok == true {
							collected = append(collected, update)
						}

					}

				}

			}

		}

	} else if SUPPORTED == true {

		cmd1 := exec.Command("pacman", "-Sy", "--noconfirm")
		_, err1 := cmd1.Output()

		if err1 == nil {

			cmd2 := exec.Command("pacman", "-Qu", "--noconfirm")
			buffer2, err2 := cmd2.Output()

			if err2 == nil {

				lines := strings.Split(strings.TrimSpace(string(buffer2)), "\n")

				for l := 0; l < len(lines); l++ {

					line := strings.TrimSpace(lines[l])

					if strings.HasSuffix(line, "[ignored]") {
						line = strings.TrimSpace(line[0 : len(line)-9])
					}

					if strings.Contains(line, " ") && strings.Contains(line, " -> ") {

						// "package 1.2.3 -> 1.2.4"

						name := line[0:strings.Index(line, " ")]
						update := CollectUpdate(name)

						if update.Name != "" && update.Version.IsValid() {

							if strings.HasPrefix(update.Name, "lib32-") {
								update.SetArchitecture("x86")
							}

							collected = append(collected, update)

						}

					} else if strings.Contains(line, " ") {

						// "package 1.2.3"

						name := line[0:strings.Index(line, " ")]
						update := CollectUpdate(name)

						if update.Name != "" && update.Version.IsValid() {

							if strings.HasPrefix(update.Name, "lib32-") {
								update.SetArchitecture("x86")
							}

							collected = append(collected, update)

						}

					}

				}

			}

		}

	}

	return collected

}
