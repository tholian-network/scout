package sys

import "tholian-core/structs"

func findSubsystemDevice(candidates []*structs.Device, id_vendor string, id_device string, id_subsystem_vendor string, id_subsystem_device string) *structs.Device {

	var result *structs.Device = nil

	for _, other := range candidates {

		if other.System != nil &&
			other.System.Vendor == id_vendor &&
			other.System.Device == id_device &&
			other.Subsystem != nil &&
			other.Subsystem.Vendor == id_subsystem_vendor &&
			other.Subsystem.Device == id_subsystem_device {
			result = other
			break
		}

	}

	return result

}
