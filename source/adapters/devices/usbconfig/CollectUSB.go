package usbconfig

import "tholian-core/structs"
import "fmt"
import "os/exec"
import "regexp"
import "strconv"
import "strings"

var usb_entry_pattern = regexp.MustCompile(`(ugen[0-9]+\.[0-9]+):\s+<([^>]*)>`)
var usb_vendor_pattern = regexp.MustCompile(`idVendor=([^\s]+)`)
var usb_product_pattern = regexp.MustCompile(`idProduct=([^\s]+)`)

type usb_entry struct {
	Address string
	Name    string
}

func CollectUSB() []structs.Device {

	var collected []structs.Device

	if SUPPORTED == true {

		cmd := exec.Command("usbconfig", "list")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			var usb_entries []usb_entry

			lines := strings.Split(strings.ReplaceAll(string(buffer), "\r\n", "\n"), "\n")

			for l := 0; l < len(lines); l++ {

				matches := usb_entry_pattern.FindStringSubmatch(lines[l])

				if len(matches) == 3 {

					usb_entries = append(usb_entries, usb_entry{
						Address: matches[1],
						Name:    strings.TrimSpace(matches[2]),
					})

				}

			}

			for e := 0; e < len(usb_entries); e++ {

				descriptor, err := exec.Command("usbconfig", "-d", usb_entries[e].Address, "dump_device_desc").Output()

				if err == nil {

					vendor := ""
					device := ""
					content := strings.ReplaceAll(string(descriptor), "\r\n", "\n")

					if matches := usb_vendor_pattern.FindStringSubmatch(content); len(matches) == 2 {

						tmp := strings.ToLower(strings.TrimSpace(matches[1]))

						if strings.HasPrefix(tmp, "0x") {

							if num, err := strconv.ParseUint(tmp, 16, 64); err == nil {
								vendor = fmt.Sprintf("%04x", num)
							}

						}

					}

					if matches := usb_product_pattern.FindStringSubmatch(content); len(matches) == 2 {

						tmp := strings.ToLower(strings.TrimSpace(matches[1]))

						if strings.HasPrefix(tmp, "0x") {

							if num, err := strconv.ParseUint(tmp, 16, 64); err == nil {
								device = fmt.Sprintf("%04x", num)
							}

						}

					}

					if vendor != "" && device != "" {

						name := strings.TrimSpace(usb_entries[e].Name)

						if name == "" {
							name = "Unknown Device"
						}

						entry := structs.NewDevice("usb")
						entry.SetName(name)
						entry.SetSystem(vendor, device, name)

						if entry.IsValid() {
							collected = append(collected, entry)
						}

					}

				}

			}

		}

	}

	return collected

}

