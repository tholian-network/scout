package systemd

import "strconv"
import "strings"

func toGID(value string) uint16 {

	var result uint16

	if strings.Contains(value, ":") {

		num, err := strconv.ParseUint(value[strings.Index(value, ":")+1:], 10, 16)

		if err == nil {
			result = uint16(num)
		}

	} else {

		num, err := strconv.ParseUint(value, 10, 16)

		if err == nil {
			result = uint16(num)
		}

	}

	return result

}
