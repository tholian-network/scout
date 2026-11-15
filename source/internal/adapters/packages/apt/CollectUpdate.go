package apt

import "tholian-core/structs"
import "tholian-core/types"
import "os"
import "os/exec"
import "strings"

func CollectUpdate(name string, version string) structs.Update {

	var result structs.Update = structs.NewUpdate("apt")

	if version == "" {
		version = "*"
	}

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("apt-cache", "show", name)
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			blocks := strings.Split(strings.TrimSpace(string(buffer)), "\n\n")

			if len(blocks) > 0 {

				var block string
				var block_version = types.ToVersion(version)

				if version == "*" {

					block = strings.TrimSpace(blocks[0])

				} else {

					for b := 0; b < len(blocks); b++ {

						lines := strings.Split(strings.TrimSpace(blocks[b]), "\n")

						for l := 0; l < len(lines); l++ {

							line := strings.TrimSpace(lines[l])

							if strings.HasPrefix(line, "Version: ") {

								tmp_version := types.ToVersion(strings.TrimSpace(line[9:]))

								if tmp_version.String() == block_version.String() {
									block = strings.TrimSpace(blocks[b])
									break
								}

							}

							if block != "" {
								break
							}

						}

						if block != "" {
							break
						}

					}

				}

				if block != "" {
					ParseUpdate(block, &result)
				}

			}

		}

	}

	return result

}
