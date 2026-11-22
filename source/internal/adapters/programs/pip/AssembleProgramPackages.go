package pip

import "tholian-core/structs"
import "strings"

func AssembleProgramPackages(result *structs.Program) {

	if SUPPORTED == true && result.Folder != "" && strings.HasPrefix(result.Folder, "/") {

		for _, pkg := range collectPackages(result.Folder) {
			result.AddPackage(pkg)
		}

	}

}
