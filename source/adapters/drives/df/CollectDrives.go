package df

import "tholian-core/structs"
import utils_strings "tholian-scout/internal/utils/strings"
import "os/exec"
import "strings"

func CollectDrives() []structs.Drive {

	var result []structs.Drive

	if SUPPORTED == true {

		cmd := exec.Command("df", "-h")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")
			headers := utils_strings.Split(lines[0], " ")

			for l := 1; l < len(lines); l++ {

				chunks := utils_strings.Split(lines[l], " ")
				drive := structs.NewDrive(chunks[0], "local")

				for c := 0; c < len(chunks); c++ {

					if headers[c] == "Filesystem" {
						drive.SetName(chunks[c])
					} else if headers[c] == "Size" {
						drive.SetSize(toUint64(chunks[c]))
					} else if headers[c] == "Used" {
						drive.SetUsed(toUint64(headers[c]))
					} else if headers[c] == "Avail" || headers[c] == "Available" {
						drive.SetFree(toUint64(chunks[c]))
					} else if headers[c] == "Use%" {
						// Do Nothing
					} else if headers[c] == "Mounted" {
						drive.SetMountpoint(chunks[c])
					}

				}

				result = append(result, drive)

			}

		}

	}

	return result

}
