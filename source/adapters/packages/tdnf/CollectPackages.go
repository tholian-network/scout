package tdnf

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("tdnf", "--noplugins", "info", "installed")
		buffer, err := cmd.Output()

		if err == nil {

			blocks := strings.Split("\n\n"+strings.TrimSpace(string(buffer)), "\n\nName")

			for b := 0; b < len(blocks); b++ {

				block := strings.TrimSpace(blocks[b])

				if block != "" {

					pkg := structs.NewPackage("tdnf")
					ParsePackage("Name "+block, &pkg)

					if pkg.Name != "" && pkg.Version.IsValid() {

						cmd2 := exec.Command("tdnf", "--noplugins", "repoquery", "--installed", "--list", pkg.Name)
						buffer2, err2 := cmd2.Output()

						if err2 == nil {

							lines := strings.Split(strings.TrimSpace(string(buffer2)), "\n")

							for l := 0; l < len(lines); l++ {

								line := strings.TrimSpace(lines[l])

								if strings.HasPrefix(line, "/") {
									pkg.AddFilesystem(line)
								}

							}

						}

						collected = append(collected, pkg)

					}

				}

			}

		}

	}

	return collected

}
