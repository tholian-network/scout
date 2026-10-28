package tdnf

import "tholian-core/structs"
import "strings"

func ParseUpdate(buffer string, result *structs.Update) {

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	var key string
	var val string

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if strings.Contains(line, " : ") {

			key = strings.TrimSpace(line[0:strings.Index(line, " : ")])
			val = strings.TrimSpace(line[strings.Index(line, " : ")+3:])

		} else if strings.HasSuffix(line, " :") {

			key = strings.TrimSpace(line[0:strings.Index(line, " :")])
			val = ""

		} else {

			val = strings.TrimSpace(line)

		}

		if key == "Architecture" {

			result.SetArchitecture(val)

		} else if key == "Name" {

			result.SetName(val)

		} else if key == "URL" {

			if val != "" {
				result.SetURL(val)
			}

		} else if key == "Version" {

			result.SetVersion(val)

		}

	}

}
