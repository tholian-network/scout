package etc

import "tholian-core/structs"

func CollectSystem() structs.System {

	var collected structs.System

	if SUPPORTED == true {

		var system = toSystem()

		if system.Name != "" {
			collected = system
		}

	}

	return collected

}
