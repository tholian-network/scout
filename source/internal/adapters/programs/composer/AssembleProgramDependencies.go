package composer

import "tholian-core/structs"
import "strings"

func AssembleProgramDependencies(result *structs.Program) {

	if SUPPORTED == true && result.Folder != "" && strings.HasPrefix(result.Folder, "/") {

		if isFile(result.Folder + "/composer.json") {

			manifest := NewManifest(result.Folder + "/composer.json")

			for _, dependency := range manifest.ToPackage().Dependencies {
				result.AddDependency(dependency)
			}

		}

	}

}
