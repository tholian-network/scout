package kenv

import "tholian-core/structs"

func CollectBoard() structs.Device {

	collected := structs.NewDevice("other")

	if SUPPORTED == true {

		environment := CollectKernelEnvironment()

		name := environment["smbios.planar.product"]
		vendor := environment["smbios.planar.maker"]
		version := environment["smbios.planar.version"]
		serial := environment["smbios.planar.serial"]

		if name == "" {
			name = environment["smbios.system.product"]
		}

		if vendor == "" {
			vendor = environment["smbios.system.maker"]
		}

		if version == "" {
			version = environment["smbios.system.version"]
		}

		if serial == "" {
			serial = environment["smbios.system.serial"]
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

