package npm

import "tholian-core/structs"
import "os"
import "strings"

func CollectPackages(folder string) []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true && folder != "" && strings.HasPrefix(folder, "/") {

		stat1, err1 := os.Stat(folder + "/node_modules")

		if err1 == nil && stat1.IsDir() == true {

			entries, err2 := os.ReadDir(folder + "/node_modules")

			if err2 == nil {

				for e := 0; e < len(entries); e++ {

					name := entries[e].Name()

					if entries[e].IsDir() {

						if strings.HasPrefix(name, ".") {
							// Do Nothing
						} else if strings.HasPrefix(name, "@") {

							subentries, err3 := os.ReadDir(folder + "/node_modules/" + name)

							if err3 == nil {

								for s := 0; s < len(subentries); s++ {

									name := entries[e].Name() + "/" + subentries[s].Name()
									manifest := NewManifest(folder + "/node_modules/" + name + "/package.json")
									pkg := manifest.ToPackage()

									if pkg.Name != "" && pkg.Version.IsValid() {
										collected = append(collected, pkg)
									}

								}

							}

						} else {

							manifest := NewManifest(folder + "/node_modules/" + name + "/package.json")
							pkg := manifest.ToPackage()

							if pkg.Name != "" && pkg.Version.IsValid() {
								collected = append(collected, pkg)
							}

						}

					}

				}

			}

		}

	}

	return collected

}
