package procstat

import "tholian-core/structs"

func CollectServices() []structs.Program {

	var collected []structs.Program

	for _, program := range CollectProcesses() {

		if program.IsService() {
			collected = append(collected, program)
		}

	}

	return collected

}
