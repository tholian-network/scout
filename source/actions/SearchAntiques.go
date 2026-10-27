package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchAntiques(console *structs.Console, system *structs.System, search matchers.Antique) []structs.Antique {

	console.Group("actions/SearchAntiques")

	var filtered []structs.Antique

	for a := 0; a < len(system.Antiques); a++ {

		antique := system.Antiques[a]

		if search.Matches(antique.Name, antique.Version.String(), antique.Architecture.String(), antique.Manager.String(), antique.Vendor, antique.Service) {
			filtered = append(filtered, antique)
		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Antiques)) + " Antiques")
	console.GroupEnd("actions/SearchAntiques")

	return filtered

}
