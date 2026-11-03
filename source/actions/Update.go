package actions

import "tholian-core/matchers"
import "tholian-core/structs"

// Update returns the remaining available updates
func Update(console *structs.Console, system *structs.System) []structs.Update {

	if system == nil {
		system = initSystem(console)
	}

	console.Group("actions/Update")

	CollectUpdates(console, system)

	remaining := make([]structs.Update, 0)
	searches := make([]matchers.Package, 0)

	for u := 0; u < len(system.Updates); u++ {

		update := system.Updates[u]
		search := matchers.NewPackage()

		search.SetName(update.Name)
		search.SetVersion(update.Version.String())
		search.SetArchitecture(update.Architecture.String())
		search.SetVendor("any")

		searches = append(searches, search)

	}

	residual := UpdatePackages(console, system, searches)

	if len(residual) > 0 {
		remaining = residual
	}

	console.GroupEnd("actions/Update")

	return remaining

}
