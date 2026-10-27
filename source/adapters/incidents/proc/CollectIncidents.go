package proc

import "tholian-core/structs"
import "tholian-scout/adapters/programs/proc"

func CollectIncidents(system string) []structs.Incident {

	var result []structs.Incident

	if SUPPORTED == true {

		proc.CollectPrograms()

		// for _, program := range proc.Programs.Programs {

			// TODO: Analyze environment variables (e.g. LD_PRELOAD)
			// TODO: Analyze injected shared objects (e.g. glibc so files)
			// TODO: Analyze suspicious process commands and arguments (e.g. netcat)

		// }

	}

	return result

}
