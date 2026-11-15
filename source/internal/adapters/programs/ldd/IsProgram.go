package ldd

import "tholian-core/structs"
import "os"
import "strings"

func IsProgram(program structs.Program) bool {

	var result bool = false

	if program.Folder != "" && strings.HasPrefix(program.Folder, "/") {

		if program.Command != "" && strings.HasPrefix(program.Command, "/") {

			file, err1 := os.Open(program.Command)

			if err1 == nil {

				check := make([]byte, 4)
				_, err2 := file.ReadAt(check, 0)

				if err2 == nil &&
					check[0] == 0x7f &&
					check[1] == 'E' &&
					check[2] == 'L' &&
					check[3] == 'F' {
					result = true
				}

			}

		}

	}

	return result

}
