package sys

import "tholian-core/structs"

func findSystemDevice(candidates []*structs.Device, id_vendor string, id_device string) *structs.Device {

	var result *structs.Device = nil

	for _, other := range candidates {

		if other.System != nil &&
			other.System.Vendor == id_vendor &&
			other.System.Device == id_device {
			result = other
			break
		}

	}

	return result

}

