package procstat

import "tholian-core/structs"
import "strings"

func CollectProgram(name string) []structs.Program {

	var collected []structs.Program

	for _, program := range CollectProcesses() {

		if program.Name == name || program.Command == name || strings.HasSuffix(program.Command, "/"+name) {
			collected = append(collected, program)
		}

	}

	return collected

}
