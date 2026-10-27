package actions

import "tholian-core/structs"
import utils_fs "tholian-scout/internal/utils/fs"

func Install(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/Install")

	result = utils_fs.CopySelf("/usr/bin/tholian-scout")

	console.GroupEnd("actions/Install")

	return result

}
