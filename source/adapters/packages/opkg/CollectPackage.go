package opkg

import "tholian-core/structs"
import "strings"

func CollectPackage(name string) structs.Package {

	var result structs.Package = structs.NewPackage("opkg")

	if SUPPORTED == true {

		buffer, err := readStatus()

		if err == nil && len(buffer) > 0 {

			collected := parseStatus(string(buffer))
			requested := strings.TrimSpace(name)

			for c := 0; c < len(collected); c++ {

				if collected[c].Name == requested {

					result = collected[c]
					attachFiles(&result)
					break

				}

			}

		}

	}

	return result

}
