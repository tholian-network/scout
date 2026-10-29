package actions

import "tholian-core/structs"
import "tholian-scout/adapters/devices/pciconf"
import "tholian-scout/adapters/devices/sys"
import "tholian-scout/adapters/devices/usbconfig"
import "strconv"

func CollectDevices(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectDevices")

	collected := make([]structs.Device, 0)
	candidates := make([]structs.Device, 0)

	tmp1 := sys.CollectPCI()

	if len(tmp1) > 0 {

		console.Info("sys.CollectPCI(): Found " + strconv.Itoa(len(tmp1)) + " Devices")
		candidates = append(candidates, tmp1...)

	} else {

		console.Warn("sys.CollectPCI(): Found 0 Devices")

	}

	tmp2 := sys.CollectUSB()

	if len(tmp2) > 0 {

		console.Info("sys.CollectUSB(): Found " + strconv.Itoa(len(tmp2)) + " Devices")
		candidates = append(candidates, tmp2...)

	} else {

		console.Warn("sys.CollectUSB(): Found 0 Devices")

	}

	tmp3 := pciconf.CollectPCI()

	if len(tmp3) > 0 {

		console.Info("pciconf.CollectPCI(): Found " + strconv.Itoa(len(tmp3)) + " Devices")
		candidates = append(candidates, tmp3...)

	} else {

		console.Warn("pciconf.CollectPCI(): Found 0 Devices")

	}

	tmp4 := usbconfig.CollectUSB()

	if len(tmp4) > 0 {

		console.Info("usbconfig.CollectUSB(): Found " + strconv.Itoa(len(tmp4)) + " Devices")
		candidates = append(candidates, tmp4...)

	} else {

		console.Warn("usbconfig.CollectUSB(): Found 0 Devices")

	}

	for _, device := range candidates {

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

	if len(collected) > 0 {
		system.SetDevices(collected)
		result = true
	}

	console.Log("Collected " + strconv.Itoa(len(system.Devices)) + "/" + strconv.Itoa(len(collected)) + " Devices")
	console.GroupEnd("actions/CollectDevices")

	return result

}
