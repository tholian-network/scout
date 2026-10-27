package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchUpdates(console *structs.Console, system *structs.System, search matchers.Update) []structs.Update {

	console.Group("actions/SearchUpdates")

	var filtered []structs.Update

	for u := 0; u < len(system.Updates); u++ {

		update := system.Updates[u]

		if search.Matches(update.Name, update.Version.String(), update.Architecture.String(), update.Manager.String(), update.Vendor) {
			filtered = append(filtered, update)
		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Updates)) + " Updates")
	console.GroupEnd("actions/SearchUpdates")

	return filtered

}
