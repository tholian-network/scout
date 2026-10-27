package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchDevices(console *structs.Console, system *structs.System, search matchers.Device) []structs.Device {

	console.Group("actions/SearchDevices")

	var filtered []structs.Device

	for d := 0; d < len(system.Devices); d++ {

		device := system.Devices[d]
		matches_system := false
		matches_subsystem := false

		if device.System != nil {

			if search.Matches(device.Name, device.System.Vendor, device.System.Device, device.Bus) {
				matches_system = true
			}

		}

		if device.Subsystem != nil {

			if search.Matches(device.Name, device.Subsystem.Vendor, device.Subsystem.Device, device.Bus) {
				matches_subsystem = true
			}

		}

		if matches_system || matches_subsystem {
			filtered = append(filtered, device)
		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Devices)) + " Devices")
	console.GroupEnd("actions/SearchDevices")

	return filtered

}
