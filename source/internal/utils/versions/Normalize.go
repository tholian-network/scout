package versions

import "tholian-core/types"
import "strings"

func Normalize(value string) (string, bool) {

	value = strings.TrimSpace(value)

	if value == "" || value == "*" || value == "x" || value == "X" {
		return "", false
	}

	lower := strings.ToLower(value)

	if lower == "any" || lower == "all" || lower == "latest" || lower == "next" {
		return "", false
	}

	if strings.HasPrefix(lower, "v") && len(value) > 1 {
		value = value[1:]
	}

	// PEP 440 uses ! for epoches, tholian-core uses :
	value = strings.ReplaceAll(value, "!", ":")

	chunks := strings.Split(value, ".")

	for c := 0; c < len(chunks); c++ {

		if chunks[c] == "x" || chunks[c] == "X" || chunks[c] == "*" {
			chunks[c] = "0"
		}

	}

	value = strings.Join(chunks, ".")

	if types.IsVersion(value) {
		return value, true
	}

	return "", false

}
