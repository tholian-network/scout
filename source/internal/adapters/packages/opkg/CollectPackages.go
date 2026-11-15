package opkg

import "tholian-core/structs"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true {

		buffer, err := readStatus()

		if err == nil && len(buffer) > 0 {

			collected = parseStatus(string(buffer))

			for c := 0; c < len(collected); c++ {
				attachFiles(&collected[c])
			}

		}

	}

	return collected

}
