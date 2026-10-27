package proc

import "tholian-core/structs"
import "os"
import "strconv"
import "strings"

func AssembleProgramFilesystem(result *structs.Program, process_id uint) {

	// Optional argument
	if process_id == 0 {
		process_id = result.PID
	}

	if process_id != 0 {

		pid := strconv.FormatUint(uint64(process_id), 10)
		addresses, err1 := os.ReadDir("/proc/" + pid + "/map_files")

		if err1 == nil {

			for a := 0; a < len(addresses); a++ {

				link, err2 := os.Readlink("/proc/" + pid + "/map_files/" + addresses[a].Name())

				if err2 == nil {

					file := strings.TrimSpace(string(link))

					if strings.HasPrefix(file, "/") {
						result.AddFilesystem(file)
					}

				}

			}

		}

	}

}
