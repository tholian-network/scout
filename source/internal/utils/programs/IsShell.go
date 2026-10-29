package programs

import "tholian-core/structs"

var shell_names = []string{
	"ash",
	"bash",
	"dash",
	"sh",
	"zsh",
}

func IsShell(program structs.Program) bool {

	for _, value := range toProgramNames(program) {

		for n := 0; n < len(shell_names); n++ {

			if value == shell_names[n] {
				return true
			}

		}

	}

	return false

}
