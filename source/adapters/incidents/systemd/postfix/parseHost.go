package postfix

import "strings"

func parseHost(chunk string) string {

	var result string

	if strings.Contains(chunk, "[") && strings.HasSuffix(chunk, "]:") {

		tmp := strings.TrimSpace(chunk[strings.Index(chunk, "[")+1 : len(chunk)-2])

		if tmp != "unknown" && tmp != "" {
			result = tmp
		}

	} else if strings.Contains(chunk, "[") && strings.HasSuffix(chunk, "]") {

		tmp := strings.TrimSpace(chunk[strings.Index(chunk, "[")+1 : len(chunk)-1])

		if tmp != "unknown" && tmp != "" {
			result = tmp
		}

	}

	return result

}
