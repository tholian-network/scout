package versions

import "strings"

func HasWildcard(value string) bool {

	chunks := strings.Split(strings.TrimPrefix(strings.TrimSpace(value), "v"), ".")

	for c := 0; c < len(chunks); c++ {

		if chunks[c] == "x" || chunks[c] == "X" || chunks[c] == "*" {
			return true
		}

	}

	return false

}
