package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchPrograms(console *structs.Console, system *structs.System, search matchers.Program) []structs.Program {

	console.Group("actions/SearchPrograms")

	var filtered []structs.Program

	for p := 0; p < len(system.Programs); p++ {

		program := system.Programs[p]

		if search.Matches(program.Name, program.Command, program.Arguments) {
			filtered = append(filtered, program)
		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Programs)) + " Programs")
	console.GroupEnd("actions/SearchPrograms")

	return filtered

}
