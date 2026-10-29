package pkg

import "tholian-core/structs"
import "os"
import "os/exec"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true {

		cmd0 := exec.Command("pkg", "query", "-a", "%n\t%v\t%q\t%o\t%w\t%m\t%p\t%t\t%R")
		cmd0.Env = append(os.Environ(), "LC_ALL=C")

		buffer, err0 := cmd0.Output()

		if err0 == nil && len(buffer) > 0 {

			collected = parsePackages(string(buffer))

			cmd1 := exec.Command("pkg", "query", "-a", "%n\t%dn\t%dv")
			cmd1.Env = append(os.Environ(), "LC_ALL=C")

			buffer_dependencies, err1 := cmd1.Output()

			cmd2 := exec.Command("pkg", "query", "-a", "%n\t%Fp")
			cmd2.Env = append(os.Environ(), "LC_ALL=C")

			buffer_files, err2 := cmd2.Output()

			for c := 0; c < len(collected); c++ {

				if err1 == nil && len(buffer_dependencies) > 0 {

					for _, dependency := range parsePackageDependencies(string(buffer_dependencies), collected[c].Name) {
						collected[c].AddDependency(dependency)
					}

				}

				if err2 == nil && len(buffer_files) > 0 {

					for _, file := range parsePackageFiles(string(buffer_files), collected[c].Name) {
						addFilesystem(&collected[c], file)
					}

				}

			}

		}

	}

	return collected

}
