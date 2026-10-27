package actions

import "tholian-core/structs"
import "tholian-scout/adapters/devices/sys"
import "strconv"

func CollectDevices(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectDevices")

	collected := make([]structs.Device, 0)

	tmp1 := sys.CollectPCI()

	if len(tmp1) > 0 {

		console.Info("sys.CollectPCI(): Found " + strconv.Itoa(len(tmp1)) + " Devices")

		for _, device := range tmp1 {

			if device.IsValid() {

				found := false

				for c := 0; c < len(collected); c++ {

					if collected[c].Equal(device) {
						found = true
						break
					}

				}

				if found == false {
					collected = append(collected, device)
				}

			}

		}

	} else {

		console.Warn("sys.CollectPCI(): Found 0 Devices")

	}

	tmp2 := sys.CollectUSB()

	if len(tmp2) > 0 {

		console.Info("sys.CollectUSB(): Found " + strconv.Itoa(len(tmp2)) + " Devices")

		for _, device := range tmp2 {

			if device.IsValid() {

				found := false

				for c := 0; c < len(collected); c++ {

					if collected[c].Equal(device) {
						found = true
						break
					}

				}

				if found == false {
					collected = append(collected, device)
				}

			}

		}

	} else {

		console.Warn("sys.CollectUSB(): Found 0 Devices")

	}

	if len(collected) > 0 {
		system.SetDevices(collected)
		result = true
	}

	console.Log("Collected " + strconv.Itoa(len(system.Devices)) + "/" + strconv.Itoa(len(collected)) + " Devices")
	console.GroupEnd("actions/CollectDevices")

	return result

}
