package kenv

import "tholian-core/structs"

func CollectBIOS() structs.Device {

	collected := structs.NewDevice("other")

	if SUPPORTED == true {

		environment := CollectKernelEnvironment()

		vendor := environment["smbios.bios.vendor"]
		version := environment["smbios.bios.version"]
		release := environment["smbios.bios.reldate"]

		if release == "" {
			release = version
		}

		if version == "" {
			version = release
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
