package zypper

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

			if key == "Provides" || key == "Requires" || key == "Conflicts" || key == "Obsoletes" {

				// Ignore counters for multi-line output
				if strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]") {
					val = ""
				}

			}

		} else {

			val = strings.TrimSpace(line)

		}

		if key == "Arch" {

			if val != "" {
				result.SetArchitecture(val)
			}

		} else if key == "Name" {

			if val != "" {
				result.SetName(val)
			}

		} else if key == "Upstream URL" {

			if val != "" {
				result.SetURL(val)
			}

		} else if key == "Version" {

			if val != "" {
				result.SetVersion(val)
			}

		}

	}

}
