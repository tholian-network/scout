package sys

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-insights/devices"
import "os"

func CollectPCI() []structs.Device {

	var collected []structs.Device

	if SUPPORTED == true {

		entries, err1 := os.ReadDir("/sys/bus/pci/devices")

		if err1 == nil {

			for e := 0; e < len(entries); e++ {

				reference := entries[e].Name()

				vendor := readPCI(reference, "vendor")
				device := readPCI(reference, "device")
				subsystem_vendor := readPCI(reference, "subsystem_vendor")
				subsystem_device := readPCI(reference, "subsystem_device")

				if vendor != "" && device != "" {

					if subsystem_vendor != "" && subsystem_device != "" {

						candidate := findSubsystemDevice(devices.Devices.QueryBySubsystem(matchers.Device{
							Name: "any",
							Bus: "pci",
							Vendor: subsystem_vendor,
							Device: subsystem_device,
						}), vendor, device, subsystem_vendor, subsystem_device)

						if candidate != nil {

							dev := structs.NewDevice("pci")
							dev.SetName(candidate.Name)
							dev.SetSystem(candidate.System.Vendor, candidate.System.Device, candidate.System.Name)
							dev.SetSubsystem(candidate.Subsystem.Vendor, candidate.Subsystem.Device, candidate.Subsystem.Name)

							if dev.Name == "" {

								if dev.Subsystem.Name == "" && dev.System.Name != "" {
									dev.Name = dev.System.Name
								} else {
									dev.Name = "Unknown Device"
								}

							}

							collected = append(collected, dev)

						}

					} else {

						candidate := findSystemDevice(devices.Devices.QueryBySubsystem(matchers.Device{
							Name: "any",
							Bus: "pci",
							Vendor: vendor,
							Device: device,
						}), vendor, device)

						if candidate != nil {

							dev := structs.NewDevice("pci")
							dev.SetName(candidate.Name)
							dev.SetSystem(candidate.System.Vendor, candidate.System.Device, candidate.System.Name)

							if dev.Name == "" {

								if dev.System.Name != "" {
									dev.Name = dev.System.Name
								} else {
									dev.Name = "Unknown Device"
								}

							}

							collected = append(collected, dev)

						}

					}

				}

			}

		}

	}

	return collected

}
