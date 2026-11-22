package npm

import "tholian-core/structs"
import "strings"

func AssembleProgramDependencies(result *structs.Program) {

	if SUPPORTED == true && result.Folder != "" && strings.HasPrefix(result.Folder, "/") {

		if isFile(result.Folder + "/package.json") {

			manifest := NewManifest(result.Folder + "/package.json")

			for _, dependency := range manifest.ToPackage().Dependencies {
				result.AddDependency(dependency)
			}

		}

	}

}
