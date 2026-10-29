package os

import "path"
import "strings"

func NormalizePath(value string) string {

	value = strings.TrimSpace(value)
	value = strings.TrimSuffix(value, " (deleted)")
	value = strings.TrimSpace(value)

	if strings.HasPrefix(value, "/overlay/upper/") {
		value = "/" + strings.TrimPrefix(value, "/overlay/upper/")
	} else if strings.HasPrefix(value, "/overlay/") {
		value = "/" + strings.TrimPrefix(value, "/overlay/")
	} else if strings.HasPrefix(value, "/rom/") {
		value = "/" + strings.TrimPrefix(value, "/rom/")
	}

	if value == "" {
		return value
	}

	return path.Clean(value)

}

