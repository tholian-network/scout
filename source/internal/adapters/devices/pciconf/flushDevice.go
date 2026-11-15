package pciconf

import "tholian-core/structs"

func flushDevice(current *structs.Device, names map[string]string, ids map[string]string) (structs.Device, bool) {

	if current == nil {
		return structs.Device{}, false
	}

	vendor := ids["vendor"]
	device := ids["device"]

	if vendor == "" || device == "" {
		return structs.Device{}, false
	}

	name := names["device"]

	if name == "" {
		name = names["vendor"]
	}

	if name == "" {
		name = "Unknown Device"
	}

	current.SetName(name)
	current.SetSystem(vendor, device, name)

	subvendor := ids["subvendor"]
	subdevice := ids["subdevice"]

	if subvendor != "" && subdevice != "" && subvendor != "0000" && subdevice != "0000" {
		current.SetSubsystem(subvendor, subdevice, name)
	}

	return *current, true

}
