package npm

import "tholian-core/structs"
import "os"
import "strings"

func AssembleProgramPackages(result *structs.Program) {

	if SUPPORTED == true && result.Folder != "" && strings.HasPrefix(result.Folder, "/") {

		stat1, err1 := os.Stat(result.Folder + "/node_modules")

		if err1 == nil && stat1.IsDir() == true {

			entries, err2 := os.ReadDir(result.Folder + "/node_modules")

			if err2 == nil {

				for e := 0; e < len(entries); e++ {

					name := entries[e].Name()

					if entries[e].IsDir() {

						if strings.HasPrefix(name, ".") {
							// Do Nothing
						} else if strings.HasPrefix(name, "@") {

							subentries, err3 := os.ReadDir(result.Folder + "/node_modules/" + name)

							if err3 == nil {

								for s := 0; s < len(subentries); s++ {

									name := entries[e].Name() + "/" + subentries[s].Name()
									manifest := NewManifest(result.Folder + "/node_modules/" + name + "/package.json")
									pkg := manifest.ToPackage()

									if pkg.Name != "" && pkg.Version.IsValid() {
										result.Packages = append(result.Packages, pkg)
									}

								}

							}

						} else {

							manifest := NewManifest(result.Folder + "/node_modules/" + name + "/package.json")
							pkg := manifest.ToPackage()

							if pkg.Name != "" && pkg.Version.IsValid() {
								result.Packages = append(result.Packages, pkg)
							}

						}

					}

				}

			}

		}

	}

}
