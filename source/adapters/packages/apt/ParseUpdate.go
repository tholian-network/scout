package apt

import "tholian-core/structs"
import "strings"

func ParseUpdate(buffer string, result *structs.Update) {

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	var key string
	var val string

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if strings.Contains(line, ": ") {

			key = strings.TrimSpace(line[0:strings.Index(line, ": ")])
			val = strings.TrimSpace(line[strings.Index(line, ": ")+2:])

		} else {

			val = val + "\n" + strings.TrimSpace(line)

		}

		if key == "Architecture" {

			result.SetArchitecture(val)

		} else if key == "Homepage" {

			result.SetURL(val)

		} else if key == "Package" {

			result.SetName(val)

		} else if key == "Version" {

			result.SetVersion(val)

		}

	}

}
