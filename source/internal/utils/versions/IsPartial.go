package versions

import "strings"

func IsPartial(value string) bool {

	value = strings.TrimSpace(value)

	if strings.HasPrefix(strings.ToLower(value), "v") && len(value) > 1 {
		value = value[1:]
	}

	if strings.ContainsAny(value, "xX*") {
		return true
	}

	if strings.Contains(value, "-") || strings.Contains(value, "+") || strings.Contains(value, ":") {
		return false
	}

	chunks := strings.Split(value, ".")
	count := 0

	for c := 0; c < len(chunks); c++ {

		if isNumber(chunks[c]) {
			count++
		} else {
			return false
		}

	}

	return count < 3

}
