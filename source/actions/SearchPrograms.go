package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchPrograms(console *structs.Console, system *structs.System, search matchers.Program) []structs.Program {

	if system == nil {
		system = initSystem(console)
	}

	CollectPrograms(console, system)
	CollectPackages(console, system)
	CollectUsers(console, system)

	console.Group("actions/SearchPrograms")

	filtered := make([]structs.Program, 0)

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
