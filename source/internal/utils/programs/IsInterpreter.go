package programs

import "tholian-core/structs"

var interpreter_names = []string{
	"ash",
	"bash",
	"dash",
	"lua",
	"node",
	"nodejs",
	"perl",
	"php",
	"python",
	"python2",
	"python3",
	"ruby",
	"sh",
	"zsh",
}

func IsInterpreter(program structs.Program) bool {

	for _, value := range toProgramNames(program) {

		for n := 0; n < len(interpreter_names); n++ {

			if value == interpreter_names[n] {
				return true
			}

		}

	}

	return false

}
