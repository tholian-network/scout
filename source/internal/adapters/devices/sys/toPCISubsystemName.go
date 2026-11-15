package sys

import "tholian-core/matchers"
import "tholian-insights/devices"

func toPCISubsystemName(id_vendor string, id_device string) string {

	var result string

	candidates := devices.Devices.QueryBySubsystem(matchers.Device{
		Name:   "any",
		Bus:    "pci",
		Vendor: id_vendor,
		Device: id_device,
	})

	if len(candidates) > 0 {

		for _, device := range candidates {

			if device.Subsystem != nil {
				result = device.Subsystem.Name
				break
			}

		}

	}

	return result

}
