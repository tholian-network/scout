package actions

import "tholian-core/structs"
import "strconv"

func CollectAntiques(console *structs.Console, system *structs.System) bool {

	console.Group("actions/CollectAntiques")

	collected := make([]structs.Antique, 0)

	for s := 0; s < len(system.Services); s++ {
		system.Services[s].ResolveDependencies(system.Packages)
	}

	for u := 0; u < len(system.Updates); u++ {

		update := system.Updates[u]

		for s := 0; s < len(system.Services); s++ {

			service := system.Services[s]

			for p := 0; p < len(service.Packages); p++ {

				pkg := service.Packages[p]

				if pkg.Name == update.Name && isSamePackageManagerFamily(pkg.Manager.String(), update.Manager.String()) {

					antique := structs.NewAntique(update.Manager.String(), service.Name)
					antique.SetName(update.Name)
					antique.SetVersion(update.Version.String())
					antique.SetArchitecture(update.Architecture.String())
					antique.SetURL(update.URL)

					if antique.IsValid() {
						collected = append(collected, antique)
					}

				}

			}

			for d := 0; d < len(service.Dependencies); d++ {

				dependency := service.Dependencies[d]

				if dependency.Name == update.Name && isSamePackageManagerFamily(dependency.Manager, update.Manager.String()) {

					antique := structs.NewAntique(update.Manager.String(), service.Name)
					antique.SetName(update.Name)
					antique.SetVersion(update.Version.String())
					antique.SetArchitecture(update.Architecture.String())
					antique.SetURL(update.URL)

					if antique.IsValid() {
						collected = append(collected, antique)
					}

				}

			}

		}

	}

	if len(collected) > 0 {
		system.SetAntiques(collected)
	}

	console.Log("Collected " + strconv.Itoa(len(system.Antiques)) + "/" + strconv.Itoa(len(collected)) + " Antiques")
	console.GroupEnd("actions/CollectAntiques")

	return true

}
