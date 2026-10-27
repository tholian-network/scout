package sys

import "tholian-core/matchers"
import "tholian-insights/devices"

func toPCISystemName(id_vendor string, id_device string) string {

	var result string

	candidates := devices.Devices.Query(matchers.Device{
		Name: "any",
		Bus: "pci",
		Vendor: id_vendor,
		Device: id_device,
	})

	if len(candidates) > 0 {

		for _, device := range candidates {

			if device.System != nil {
				result = device.System.Name
				break
			}

		}

	}

	return result

}

