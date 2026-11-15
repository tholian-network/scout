package sys

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-insights/devices"
import "os"

func CollectUSB() []structs.Device {

	var collected []structs.Device

	if SUPPORTED == true {

		entries, err1 := os.ReadDir("/sys/bus/usb/devices")

		if err1 == nil {

			for e := 0; e < len(entries); e++ {

				reference := entries[e].Name()

				vendor, device := readUSB(reference, "uevent")

				if vendor != "" && device != "" {

					candidate := findSystemDevice(devices.Devices.Query(matchers.Device{
						Name:   "any",
						Bus:    "pci",
						Vendor: vendor,
						Device: device,
					}), vendor, device)

					if candidate != nil {

						dev := structs.NewDevice("usb")
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

	return collected

}
