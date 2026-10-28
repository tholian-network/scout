package dpkg

import "strings"

func parseControlFields(buffer string) map[string]string {

	fields := make(map[string]string)
	lines := strings.Split(strings.ReplaceAll(buffer, "\r\n", "\n"), "\n")
	key := ""

	for l := 0; l < len(lines); l++ {

		line := lines[l]

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {

			if key != "" {
				fields[key] = strings.TrimSpace(fields[key] + " " + strings.TrimSpace(line))
			}

		} else if index := strings.Index(line, ":"); index > 0 {

			key = strings.TrimSpace(line[0:index])
			fields[key] = strings.TrimSpace(line[index+1:])

		}

	}

	return fields

}

