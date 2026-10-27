package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchNetworks(console *structs.Console, system *structs.System, search matchers.Network) []structs.Network {

	console.Group("actions/SearchNetworks")

	var filtered []structs.Network

	for n := 0; n < len(system.Networks); n++ {

		network := system.Networks[n]

		for s := 0; s < len(network.Subnets); s++ {

			if search.Matches(network.Name, network.Subnets[s].Address) {
				filtered = append(filtered, network)
				break
			}

		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Networks)) + " Networks")
	console.GroupEnd("actions/SearchNetworks")

	return filtered

}
