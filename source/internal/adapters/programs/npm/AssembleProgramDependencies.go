package npm

import "tholian-core/structs"
import "strings"

func AssembleProgramDependencies(result *structs.Program) {

	if SUPPORTED == true && result.Folder != "" && strings.HasPrefix(result.Folder, "/") {

		if isFile(result.Folder + "/package.json") {

			manifest := NewManifest(result.Folder + "/package.json")
			root := manifest.ToPackage()

			root.ResolveDependencies(result.Packages)

			for _, dependency := range root.Dependencies {
				result.AddDependency(dependency)
			}

			for _, unresolved := range root.Unresolved {
				result.AddUnresolved(unresolved)
			}

		}

	}

}
