package opkg

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

func isInstalledStatus(value string) bool {

	chunks := strings.Fields(value)

	if len(chunks) > 0 {
		return chunks[len(chunks)-1] == "installed"
	}

	return false

}

func splitRelations(value string) []string {

	relations := make([]string, 0)
	chunks := strings.Split(value, ",")

	for c := 0; c < len(chunks); c++ {

		chunk := strings.TrimSpace(chunks[c])

		if chunk != "" {
			relations = append(relations, chunk)
		}

	}

	return relations

}
