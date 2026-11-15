package syslog

import "strings"

func isPriority(value string) bool {

	if strings.Contains(value, ".") == false {
		return false
	}

	chunks := strings.SplitN(value, ".", 2)

	if len(chunks) != 2 {
		return false
	}

	levels := []string{"alert", "crit", "debug", "emerg", "err", "error", "info", "notice", "warning", "warn"}

	for l := 0; l < len(levels); l++ {
		if chunks[1] == levels[l] {
			return true
		}
	}

	return false

}
