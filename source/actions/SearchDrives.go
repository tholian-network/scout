package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchDrives(console *structs.Console, system *structs.System, search matchers.Drive) []structs.Drive {

	console.Group("actions/SearchDrives")

	var filtered []structs.Drive

	for d := 0; d < len(system.Drives); d++ {

		drive := system.Drives[d]

		if search.Matches(drive.Name, drive.Mountpoint, drive.Type) {
			filtered = append(filtered, drive)
		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Drives)) + " Drives")
	console.GroupEnd("actions/SearchDrives")

	return filtered

}
