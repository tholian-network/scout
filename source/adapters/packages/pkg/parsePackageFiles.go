package pkg

import "strings"

func parsePackageFiles(buffer string, parent string) []string {

	var collected []string

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		fields := strings.Split(strings.TrimSpace(lines[l]), "\t")
		file := ""

		if len(fields) >= 2 && (parent == "" || strings.TrimSpace(fields[0]) == parent) {
			file = strings.TrimSpace(fields[1])
		} else if len(fields) == 1 && parent != "" {
			file = strings.TrimSpace(fields[0])
		}

		if strings.HasPrefix(file, "/") {
			collected = append(collected, file)
		}

	}

	return collected

}
