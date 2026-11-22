package pip

import "tholian-core/structs"
import "path/filepath"
import "strings"

func IsProgram(program structs.Program) bool {

	var result bool = false

	if program.Folder != "" && strings.HasPrefix(program.Folder, "/") {

		has_manifest := isFile(program.Folder+"/requirements.txt") ||
			isFile(program.Folder+"/pyproject.toml") ||
			isFile(program.Folder+"/setup.py") ||
			isFile(program.Folder+"/setup.cfg") ||
			isFile(program.Folder+"/Pipfile")

		if has_manifest {

			name := strings.ToLower(filepath.Base(program.Command))

			has_interpreter := strings.HasPrefix(name, "python") || strings.HasPrefix(name, "pypy")
			has_venv := isVenv(program.Folder) ||
				isVenv(program.Folder+"/.venv") ||
				isVenv(program.Folder+"/venv") ||
				isVenv(program.Folder+"/env")
			has_environment := program.Environment["VIRTUAL_ENV"] != ""

			if has_interpreter == true || has_venv == true || has_environment == true {
				result = true
			}

		}

	}

	return result

}

func isVenv(folder string) bool {

	if isFile(folder + "/pyvenv.cfg") {
		return true
	}

	if isFile(folder+"/bin/python") || isFile(folder+"/bin/python3") {
		return true
	}

	if isFile(folder + "/Scripts/python.exe") {
		return true
	}

	return false

}
