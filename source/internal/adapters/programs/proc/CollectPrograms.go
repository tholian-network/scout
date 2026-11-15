package proc

import "tholian-core/structs"

func CollectPrograms() []structs.Program {

	var collected []structs.Program

	if SUPPORTED == true {

		refresh()

		for _, program := range Programs.Programs {

			if program.IsProgram() {
				collected = append(collected, *program)
			}

		}

	}

	return collected

}
