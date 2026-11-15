package npm

import "tholian-core/structs"
import "os"
import "strings"

func IsProgram(program structs.Program) bool {

	var result bool = false

	if program.Folder != "" && strings.HasPrefix(program.Folder, "/") {

		_, err1 := os.Stat(program.Folder + "/package.json")
		_, err2 := os.Stat(program.Folder + "/node_modules")

		if err1 == nil && err2 == nil {
			result = true
		}

	}

	return result

}
