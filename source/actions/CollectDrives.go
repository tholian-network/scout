package actions

import "tholian-core/structs"
import "tholian-scout/adapters/drives/df"
import "strconv"

func CollectDrives(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectDrives")

	collected := make([]structs.Drive, 0)

	if df.SUPPORTED == true {

		tmp := df.CollectDrives()

		if len(tmp) > 0 {

			console.Info("df.CollectDrives(): Found " + strconv.Itoa(len(tmp)) + " Drives")

			for _, drive := range tmp {

				if drive.IsValid() {
					collected = append(collected, drive)
				}

			}

		} else {

			console.Warn("df.CollectDrives(): Found 0 Drives")

		}

		system.SetDrives(collected)
		result = true

	} else {

		console.Warn("df.CollectDrives(): Unsupported")

	}

	console.Log("Collected " + strconv.Itoa(len(system.Drives)) + "/" + strconv.Itoa(len(collected)) + " Drives")
	console.GroupEnd("actions/CollectDrives")

	return result

}
