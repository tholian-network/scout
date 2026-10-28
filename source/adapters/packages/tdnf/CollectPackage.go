package tdnf

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-scout/adapters/packages/rpm"
import "os"
import "os/exec"
import "strings"

func CollectPackage(name string) structs.Package {

	var result structs.Package = structs.NewPackage("tdnf")

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("tdnf", "--noplugins", "info", "installed", name)
		buffer, err := cmd.Output()

		if err == nil {

			ParsePackage(string(buffer), &result)

			cmd = exec.Command("tdnf", "--noplugins", "repoquery", "--installed", "--list", name)
			buffer, err = cmd.Output()

			if err == nil {

				lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

				for l := 0; l < len(lines); l++ {

					line := strings.TrimSpace(lines[l])

					if strings.HasPrefix(line, "/") {
						result.AddFilesystem(line)
					}

				}

			}

			cmd = exec.Command("tdnf", "--noplugins", "repoquery", "--installed", "--requires", name)
			buffer, err = cmd.Output()

			if err == nil {

				lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

				for l := 0; l < len(lines); l++ {

					line := strings.TrimSpace(lines[l])

					if line != "" {

						dependency := matchers.NewPackage()
						dependency.SetManager("tdnf")

						if rpm.ParseRequire(line, &dependency) == true {
							result.AddDependency(dependency)
						}

					}

				}

			}

			cmd = exec.Command("tdnf", "--noplugins", "repoquery", "--installed", "--provides", name)
			buffer, err = cmd.Output()

			if err == nil {

				lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

				for l := 0; l < len(lines); l++ {

					line := strings.TrimSpace(lines[l])

					if line != "" {

						provide := matchers.NewPackage()
						provide.SetManager("tdnf")

						if rpm.ParseProvide(line, &provide) == true {
							result.AddProvide(provide)
						}

					}

				}

			}

			cmd = exec.Command("tdnf", "--noplugins", "repoquery", "--installed", "--obsoletes", name)
			buffer, err = cmd.Output()

			if err == nil {

				lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

				for l := 0; l < len(lines); l++ {

					line := strings.TrimSpace(lines[l])

					if line != "" {

						conflict := matchers.ToPackage(line)
						conflict.SetManager("tdnf")
						result.AddConflict(conflict)

					}

				}

			}

			cmd = exec.Command("tdnf", "--noplugins", "repoquery", "--installed", "--conflicts", name)
			buffer, err = cmd.Output()

			if err == nil {

				lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

				for l := 0; l < len(lines); l++ {

					line := strings.TrimSpace(lines[l])

					if line != "" {

						conflict := matchers.NewPackage()
						conflict.SetManager("tdnf")

						if rpm.ParseConflict(line, &conflict) == true {
							result.AddConflict(conflict)
						}

					}

				}

			}

		}

	}

	return result

}
