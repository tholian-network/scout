package actions

import "tholian-core/structs"

func Collect(console *structs.Console, system *structs.System) bool {

	console.Group("actions/Collect")

	CollectDrives(console, system)
	CollectDevices(console, system)
	CollectIncidents(console, system)
	CollectNetworks(console, system)
	CollectPrograms(console, system)
	CollectServices(console, system)
	CollectPackages(console, system)
	CollectUpdates(console, system)
	CollectUsers(console, system)

	for p := 0; p < len(system.Programs); p++ {
		system.Programs[p].ResolveDependencies(system.Packages)
	}

	CollectAntiques(console, system)

	console.GroupEnd("actions/Collect")

	return true

}
