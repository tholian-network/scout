package dropbear

import "strings"

func parseAuth(message string) (string, string) {

	username := ""
	host := ""

	if strings.Contains(message, " for '") && strings.Contains(message, "' from ") {

		tmp := strings.SplitN(message, " for '", 2)

		if len(tmp) == 2 && strings.Contains(tmp[1], "' from ") {

			chunks := strings.SplitN(tmp[1], "' from ", 2)

			username = strings.TrimSpace(chunks[0])
			host = parseHost(chunks[1])

		}

	} else if strings.Contains(message, " from ") {

		tmp := strings.SplitN(message, " from ", 2)
		host = parseHost(tmp[1])

	}

	return username, host

}

func parseHost(value string) string {

	value = strings.TrimSpace(value)

	if value == "" {
		return ""
	}

	if strings.HasPrefix(value, "[") && strings.Contains(value, "]:") {
		return value[1:strings.Index(value, "]:")]
	} else if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		return value[1 : len(value)-1]
	} else if strings.Count(value, ":") == 1 {
		return value[0:strings.Index(value, ":")]
	} else if strings.Count(value, ":") > 1 {
		return value
	} else if strings.Contains(value, " ") {
		return value[0:strings.Index(value, " ")]
	}

	return value

}
