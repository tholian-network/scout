package proc

import "tholian-core/structs"
import "sort"

// CollectProcesses returns every valid process in the cache, including
// processes that listen on sockets. CollectPrograms() intentionally filters
// services out, so incident analysis must use this function instead.
func CollectProcesses() []structs.Program {

	var collected []structs.Program

	if SUPPORTED == true {

		refresh()

		for _, program := range Programs.Programs {

			if program.IsValid() {
				collected = append(collected, *program)
			}

		}

	}

	sort.Slice(collected, func(a int, b int) bool {
		return collected[a].PID < collected[b].PID
	})

	return collected

}
