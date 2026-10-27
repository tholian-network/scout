package actions

import "tholian-core/structs"
import "os"

func Uninstall(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/Uninstall")

	err := os.Remove("/usr/bin/tholian-scout")

	if err == nil {
		result = true
	}

	console.GroupEnd("actions/Uninstall")

	return result

}
