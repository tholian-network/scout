package programs

import "tholian-core/structs"

var backdoor_names = []string{
	"nc",
	"ncat",
	"netcat",
	"socat",
}

func IsBackdoor(program structs.Program) bool {

	for _, value := range toProgramNames(program) {

		for n := 0; n < len(backdoor_names); n++ {

			if value == backdoor_names[n] {
				return true
			}

		}

	}

	return false

}
