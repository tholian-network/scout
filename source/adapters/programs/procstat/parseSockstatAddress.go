package procstat

import "strconv"
import "strings"

func parseSockstatAddress(value string) (string, uint16) {

	value = strings.TrimSpace(value)

	if value == "*:*" {

		return "any", 0

	} else if value != "" {

		index := strings.LastIndex(value, ":")

		if index == -1 {
			return value, 0
		}

		host := strings.TrimSpace(value[0:index])
		port_value := strings.TrimSpace(value[index+1:])

		if host == "*" {
			host = "any"
		}

		host = strings.TrimPrefix(host, "[")
		host = strings.TrimSuffix(host, "]")

		if port_value != "*" {

			port, err := strconv.ParseUint(port_value, 10, 16)

			if err == nil {
				return host, uint16(port)
			}

		} else {
			return host, 0
		}

	}

	return "", 0

}
