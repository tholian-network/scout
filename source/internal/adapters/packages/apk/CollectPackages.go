package apk

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true && OPTIMIZED == true {

		buffer, err := os.ReadFile("/lib/apk/db/installed")

		if err == nil {

			blocks := strings.Split("\n\n"+strings.TrimSpace(string(buffer)), "\n\nC:")

			for b := 0; b < len(blocks); b++ {

				block := strings.TrimSpace(blocks[b])

				if block != "" {

					pkg := structs.NewPackage("apk")
					parsePackageIndex("C:"+block, &pkg)

					if pkg.Name != "" && pkg.Version.IsValid() {
						collected = append(collected, pkg)
					}

				}

			}

		}

	} else if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("apk", "list", "--installed")
		buffer, err := cmd.Output()

		if err == nil {

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

					_, is_installed := flags["installed"]

					if is_installed {

						name, version := toNameAndVersion(chunks[0])
						architecture := strings.TrimSpace(chunks[1])
						pkg := CollectPackage(name)
						pkg.SetVersion(version)
						pkg.SetArchitecture(architecture)

						if pkg.Name != "" && pkg.Version.IsValid() {
							collected = append(collected, pkg)
						}

					}

				}

			}

		}

	}

	return collected

}
