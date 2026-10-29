package procstat

import "tholian-core/structs"

func CollectPrograms() []structs.Program {

	var collected []structs.Program

	for _, program := range CollectProcesses() {

		if program.IsProgram() {
			collected = append(collected, program)
		}

	}

	return collected

}

