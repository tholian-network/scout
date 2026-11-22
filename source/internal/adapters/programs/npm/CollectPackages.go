package npm

import "tholian-core/structs"
import "os"
import "sort"
import "strings"

func CollectPackages(folder string) []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true && folder != "" && strings.HasPrefix(folder, "/") {
		collected = collectPackages(folder)
	}

	return collected

}

func collectPackages(folder string) []structs.Package {

	collected := make([]structs.Package, 0)

	if folder != "" && strings.HasPrefix(folder, "/") {

		modules := folder + "/node_modules"

		if isDirectory(modules) {

			entries, err := os.ReadDir(modules)

			if err == nil {

				for e := 0; e < len(entries); e++ {

					name := entries[e].Name()

					if strings.HasPrefix(name, ".") {
						continue
					}

					if strings.HasPrefix(name, "@") {

						if isDirectory(modules + "/" + name) {

							subentries, err := os.ReadDir(modules + "/" + name)

							if err == nil {

								for s := 0; s < len(subentries); s++ {

									subname := subentries[s].Name()

									if strings.HasPrefix(subname, ".") {
										continue
									}

									path := modules + "/" + name + "/" + subname

									if isDirectory(path) {
										manifest := NewManifest(path + "/package.json")
										appendPackage(&collected, manifest.ToPackage())
									}

								}

							}

						}

					} else {

						path := modules + "/" + name

						if isDirectory(path) {
							manifest := NewManifest(path + "/package.json")
							appendPackage(&collected, manifest.ToPackage())
						}

					}

				}

			}

		}

	}

	sort.Slice(collected, func(a int, b int) bool {

		if collected[a].Name == collected[b].Name {
			return collected[a].Version.String() < collected[b].Version.String()
		}

		return collected[a].Name < collected[b].Name

	})

	return collected

}

func appendPackage(collected *[]structs.Package, pkg structs.Package) {

	if pkg.Name != "" && pkg.Version.IsValid() {

		found := false

		for c := 0; c < len(*collected); c++ {

			if (*collected)[c].Equal(pkg) {
				found = true
				break
			}

		}

		if found == false {
			*collected = append(*collected, pkg)
		}

	}

}
