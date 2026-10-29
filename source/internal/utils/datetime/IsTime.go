package datetime

import "strings"

func IsTime(value string) bool {

	if strings.Count(value, ":") != 2 {
		return false
	}

	chunks := strings.Split(value, ":")

	if len(chunks) != 3 {
		return false
	}

	for c := 0; c < len(chunks); c++ {

		if len(chunks[c]) == 0 || len(chunks[c]) > 2 {
			return false
		}

	}

	return true

}
