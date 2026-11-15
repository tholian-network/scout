package sys

import "os"
import "path"
import "strings"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/sys")
	_, err2 := os.Stat("/sys/bus/pci/devices")
	_, err3 := os.Stat("/sys/bus/usb/devices")

	if err1 == nil {

		if err2 == nil || err3 == nil {
			SUPPORTED = true
		}

	}

}

func formatString(value string, length int) string {

	var result string

	if len(value) < length {

		var prefix strings.Builder

		for p := 0; p < length-len(value); p++ {
			prefix.WriteString("0")
		}

		result = prefix.String() + value

	} else {

		result = value

	}

	return result

}

func readPCI(file string, attribute string) string {

	var result string

	link, err := os.Readlink("/sys/bus/pci/devices/" + file)

	if err == nil {

		resolved := path.Join("/sys/bus/pci/devices", link)

		if strings.HasPrefix(resolved, "/sys/devices/") {

			buffer, err1 := os.ReadFile(resolved + "/" + attribute)

			if err1 == nil {

				var identifier = strings.TrimSpace(string(buffer))

				if strings.HasPrefix(identifier, "0x") {
					identifier = identifier[2:]
				}

				result = identifier

			}

		}

	}

	return result

}

func readUSB(file string, attribute string) (string, string) {

	var result_vendor string
	var result_device string

	link, err := os.Readlink("/sys/bus/usb/devices/" + file)

	if err == nil {

		resolved := path.Join("/sys/bus/usb/devices", link)

		if strings.HasPrefix(resolved, "/sys/devices/") {

			buffer, err1 := os.ReadFile(resolved + "/" + attribute)

			if err1 == nil {

				var lines = strings.Split(strings.TrimSpace(string(buffer)), "\n")

				for l := 0; l < len(lines); l++ {

					var line = strings.TrimSpace(lines[l])

					if strings.HasPrefix(line, "PRODUCT=") {

						tmp := strings.Split(line[8:], "/")

						if len(tmp) == 3 {
							result_vendor = formatString(tmp[0], 4)
							result_device = formatString(tmp[1], 4)
						}

					}

				}

			}

		}

	}

	return result_vendor, result_device

}
