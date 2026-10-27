package sys

import "tholian-core/structs"

func CollectBoard() structs.Device {

	collected := structs.NewDevice("other")

	if SUPPORTED == true {

		var name = readDMI("board_name")
		var serial = readDMI("board_serial")
		var vendor = readDMI("board_vendor")
		var version = readDMI("board_version")

		if version == "Not Defined" {
			version = ""
		}

		if name != "" {
			collected.SetName(name)
		}

		if vendor != "" || serial != "" {
			collected.SetSystem(vendor, serial, name)
		}

		if version != "" {
			collected.SetSubsystem(vendor, version, name)
		}

	}

	return collected

}
