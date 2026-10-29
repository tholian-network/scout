package pkg

import "tholian-core/structs"
import "os"
import "os/exec"
import "sort"

func CollectUpdates() []structs.Update {

	var collected []structs.Update

	if SUPPORTED == true {

		exec.Command("pkg", "update", "-q").Run()

		cmd1 := exec.Command("pkg", "query", "-a", "%n\t%v")
		cmd1.Env = append(os.Environ(), "LC_ALL=C")

		buffer_installed_packages, err1 := cmd1.Output()

		cmd2 := exec.Command("pkg", "rquery", "-a", "%n\t%v\t%q")
		cmd2.Env = append(os.Environ(), "LC_ALL=C")

		buffer_remote_packages, err2 := cmd2.Output()

		if err1 == nil && err2 == nil {

			installed := parseInstalledPackages(string(buffer_installed_packages))
			remote := parseRemotePackages(string(buffer_remote_packages))

			for name, version := range installed {

				if candidate, ok := remote[name]; ok == true {

					if candidate.Version.IsAfter(version) {

						update := structs.NewUpdate("pkg")
						update.SetName(name)
						update.Version = candidate.Version
						update.SetArchitecture(candidate.Architecture)

						if update.Name != "" && update.Version.IsValid() {
							collected = append(collected, update)
						}

					}

				}

			}

			sort.Slice(collected, func(a int, b int) bool {
				return collected[a].Name < collected[b].Name
			})

		}

	}

	return collected

}

