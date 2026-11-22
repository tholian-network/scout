package pip

import "tholian-core/structs"
import "strings"

func AssembleProgramDependencies(result *structs.Program) {

	if SUPPORTED == true && result.Folder != "" && strings.HasPrefix(result.Folder, "/") {

		if isFile(result.Folder + "/requirements.txt") {

			manifest := NewManifest(result.Folder + "/requirements.txt")

			for _, dependency := range manifest.ToPackage().Dependencies {
				result.AddDependency(dependency)
			}

		}

	}

}
