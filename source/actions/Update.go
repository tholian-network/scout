package actions

import "tholian-core/matchers"
import "tholian-core/structs"

func Update(console *structs.Console, system *structs.System) bool {

	console.Group("actions/Update")

	CollectUpdates(console, system)

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

	result := UpdatePackages(console, system, searches)

	if result == true {
		CollectPackages(console, system)
		CollectUpdates(console, system)
	}

	console.GroupEnd("actions/Update")

	return result

}
