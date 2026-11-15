package pacman

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectPackage(name string) structs.Package {

	var result structs.Package = structs.NewPackage("pacman")

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd1 := exec.Command("pacman", "-Qi", "--noconfirm", name)
		buffer1, err1 := cmd1.Output()

		if err1 == nil {
			ParsePackage(string(buffer1), &result)
		}

		cmd2 := exec.Command("pacman", "-Ql", "--noconfirm", name)
		buffer2, err2 := cmd2.Output()

		if err2 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer2)), "\n")

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if strings.Contains(line, " ") {

					file := strings.TrimSpace(line[strings.Index(line, " ")+1:])

					if file != "" {

						if strings.HasPrefix(file, "/") {
							result.AddFilesystem(file)
						}

					}

				}

			}

		}

	}

	return result

}
