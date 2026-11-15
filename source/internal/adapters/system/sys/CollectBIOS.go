package sys

import "tholian-core/structs"

func CollectBIOS() structs.Device {

	collected := structs.NewDevice("other")

	if SUPPORTED == true {

		var date = readDMI("bios_date")
		var release = readDMI("bios_release")
		var vendor = readDMI("bios_vendor")
		var version = readDMI("bios_version")

		version = fixVersion(version, release)

		if date != "" && version != "" {
			version = version + "(" + date + ")"
		}

		if release != "" {
			collected.SetName(release)
		}

		if vendor != "" && release != "" && version != "" {
			collected.SetSystem(vendor, release, version)
		}

	}

	return collected

}
