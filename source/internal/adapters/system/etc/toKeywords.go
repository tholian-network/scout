package etc

import "strings"

func toKeywords(buffer []byte) map[string]string {

	keywords := make(map[string]string, 0)
	lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if strings.Contains(line, "=") {

			chunks := strings.Split(line, "=")

			if len(chunks) == 2 {

				key := chunks[0]
				val := chunks[1]

				if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
					val = strings.TrimSpace(val[1 : len(val)-1])
				}

				if key != "" && val != "" {
					keywords[key] = val
				}

			}

		}

	}

	return keywords

}
