package pacman

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true && OPTIMIZED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		packages := make(map[string]*structs.Package)
		cmd1 := exec.Command("pacman", "-Qi", "--noconfirm")
		buffer1, err1 := cmd1.Output()

		if err1 == nil {

			blocks := strings.Split("\n\n"+strings.TrimSpace(string(buffer1)), "\n\nName")

			for b := 0; b < len(blocks); b++ {

				block := strings.TrimSpace(blocks[b])

				if block != "" {

					pkg := structs.NewPackage("pacman")
					ParsePackage("Name "+block, &pkg)

					if pkg.Name != "" && pkg.Version.IsValid() {

						if strings.HasPrefix(pkg.Name, "lib32-") {
							pkg.SetArchitecture("x86")
						}

						packages[pkg.Name] = &pkg

					}

				}

			}

		}

		cmd2 := exec.Command("pacman", "-Ql", "--noconfirm")
		buffer2, err2 := cmd2.Output()

		if err2 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer2)), "\n")

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if strings.Contains(line, " ") {

					name := strings.TrimSpace(line[0:strings.Index(line, " ")])
					file := strings.TrimSpace(line[strings.Index(line, " ")+1:])

					pkg, ok := packages[name]

					if ok == true {

						if strings.HasPrefix(file, "/") {
							pkg.AddFilesystem(file)
						}

					}

				}

			}

		}

		for _, pkg := range packages {
			collected = append(collected, *pkg)
		}

	} else if SUPPORTED == true {

		cmd := exec.Command("pacman", "-Q", "--noconfirm")
		buffer, err := cmd.Output()

		if err == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if strings.Contains(line, " ") {

					// "package 1.2.3"

					name := line[0:strings.Index(line, " ")]
					pkg := CollectPackage(name)

					if pkg.Name != "" && pkg.Version.IsValid() {

						if strings.HasPrefix(pkg.Name, "lib32-") {
							pkg.SetArchitecture("x86")
						}

						collected = append(collected, pkg)

					}

				}

			}

		}

		for c := 0; c < len(collected); c++ {
			collected[c].ResolveDependencies(collected)
		}

	}

	return collected

}
