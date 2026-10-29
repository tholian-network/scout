package datetime

import "strings"

func IsISODateTime(value string) bool {

	if strings.Contains(value, "T") == false {
		return false
	}

	if len(value) < 19 {
		return false
	}

	return IsISODate(value[0:10]) && IsTime(value[11:19])

}
