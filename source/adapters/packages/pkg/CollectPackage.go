package pkg

import "tholian-core/structs"
import utils_path "tholian-core/utils/path"
import "os"
import "os/exec"
import "strings"

func CollectPackage(name string) structs.Package {

	name = strings.TrimSpace(name)

	var result structs.Package = structs.NewPackage("pkg")

	if SUPPORTED == true {

		if name != "" {

			cmd0 := exec.Command("pkg", "query", "%n\t%v\t%q\t%o\t%w\t%m\t%p\t%t\t%R", name)
			cmd0.Env = append(os.Environ(), "LC_ALL=C")

			buffer, err0 := cmd0.Output()

			if err0 == nil && len(buffer) > 0 {

				collected := parsePackages(string(buffer))

				if len(collected) > 0 {

					// First package is the installed one
					result = collected[0]

					cmd1 := exec.Command("pkg", "query", "%n\t%dn\t%dv", name)
					cmd1.Env = append(os.Environ(), "LC_ALL=C")

					buffer_dependencies, err1 := cmd1.Output()

					cmd2 := exec.Command("pkg", "query", "%n\t%Fp", name)
					cmd2.Env = append(os.Environ(), "LC_ALL=C")

					buffer_files, err2 := cmd2.Output()

					if err1 == nil && len(buffer_dependencies) > 0 {

						for _, dependency := range parsePackageDependencies(string(buffer_dependencies), name) {
							result.AddDependency(dependency)
						}

					}

					if err2 == nil && len(buffer_files) > 0 {

						for _, file := range parsePackageFiles(string(buffer_files), name) {

							if utils_path.IsWatchedFile(file) {

								if result.HasFilesystem(file) == false {
									result.Filesystem = append(result.Filesystem, file)
								}

							}

						}

					}

				}

			}

		}

	}

	return result

}
