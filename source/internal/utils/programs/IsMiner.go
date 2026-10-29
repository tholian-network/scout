package programs

import "tholian-core/structs"

var miner_names = []string{
	"cpuminer",
	"ethminer",
	"minerd",
	"xmrig",
}

func IsMiner(program structs.Program) bool {

	for _, value := range toProgramNames(program) {

		for n := 0; n < len(miner_names); n++ {

			if value == miner_names[n] {
				return true
			}

		}

	}

	return false

}
