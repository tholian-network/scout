package composer

import "tholian-core/structs"
import "strings"

func IsProgram(program structs.Program) bool {

	var result bool = false

	if program.Folder != "" && strings.HasPrefix(program.Folder, "/") {

		if isFile(program.Folder+"/composer.json") && isDirectory(program.Folder+"/vendor") {
			result = true
		}

	}

	return result

}
