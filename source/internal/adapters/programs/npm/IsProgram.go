package npm

import "tholian-core/structs"
import "strings"

func IsProgram(program structs.Program) bool {

	var result bool = false

	if program.Folder != "" && strings.HasPrefix(program.Folder, "/") {

		if isFile(program.Folder+"/package.json") && isDirectory(program.Folder+"/node_modules") {
			result = true
		}

	}

	return result

}
