package programs

import "tholian-core/structs"

var scanner_names = []string{
	"hashcat",
	"hydra",
	"john",
	"masscan",
	"medusa",
	"ncrack",
	"nmap",
	"sqlmap",
	"zmap",
}

func IsScanner(program structs.Program) bool {

	for _, value := range toProgramNames(program) {

		for n := 0; n < len(scanner_names); n++ {

			if value == scanner_names[n] {
				return true
			}

		}

	}

	return false

}

