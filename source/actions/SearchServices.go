package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchServices(console *structs.Console, system *structs.System, search matchers.Program) []structs.Program {

	console.Group("actions/SearchServices")

	var filtered []structs.Program

	for s := 0; s < len(system.Services); s++ {

		service := system.Services[s]

		if search.Matches(service.Name, service.Command, service.Arguments) {
			filtered = append(filtered, service)
		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Services)) + " Services")
	console.GroupEnd("actions/SearchServices")

	return filtered

}
