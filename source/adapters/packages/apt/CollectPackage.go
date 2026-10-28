package apt

import "tholian-core/structs"
import "tholian-core/types"
import "os"
import "os/exec"
import "strings"

func CollectPackage(name string, version string) structs.Package {

	var result structs.Package = structs.NewPackage("apt")

	if version == "" {
		version = "*"
	}

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("apt-cache", "-a", "show", name)
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
					ParsePackage(block, &result)
				}

			}

		}

		if result.Name != "" {

			cmd2 := exec.Command("dpkg-query", "-L", result.Name)
			buffer2, err2 := cmd2.Output()

			if err2 == nil {

				lines := strings.Split(strings.TrimSpace(string(buffer2)), "\n")

				for l := 0; l < len(lines); l++ {

					file := strings.TrimSpace(lines[l])

					if strings.HasPrefix(file, "/") {
						result.AddFilesystem(file)
					}

				}

			}

		}

		if strings.HasPrefix(result.Name, "lib32-") {
			result.Architecture = "x86"
		}

	}

	return result

}
