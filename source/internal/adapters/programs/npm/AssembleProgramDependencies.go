package npm

import "tholian-core/matchers"
import "tholian-core/structs"
import "os"
import "strings"

func AssembleProgramDependencies(result *structs.Program) {

	if SUPPORTED == true && strings.HasPrefix(result.Folder, "/") {

		stat1, err1 := os.Stat(result.Folder + "/package.json")

		if err1 == nil && stat1.IsDir() == false {

			manifest := NewManifest(result.Folder + "/package.json")

			if len(manifest.Dependencies) > 0 {

				for name, version := range manifest.Dependencies {

					dependency := matchers.NewPackage()
					dependency.SetName(name)
					dependency.SetVersion(version)
					dependency.SetArchitecture("any")
					dependency.SetManager("npm")
					dependency.SetVendor("npm")

					result.AddDependency(dependency)

				}

			}

			if len(manifest.DevDependencies) > 0 {

				for name, version := range manifest.DevDependencies {

					dependency := matchers.NewPackage()
					dependency.SetName(name)
					dependency.SetVersion(version)
					dependency.SetArchitecture("any")
					dependency.SetManager("npm")
					dependency.SetVendor("npm")

					result.AddDependency(dependency)

				}

			}

		}

	}

}
