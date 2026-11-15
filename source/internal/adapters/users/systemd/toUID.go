package systemd

import "strconv"
import "strings"

func toUID(value string) uint16 {

	var result uint16

	if strings.Contains(value, ":") {

		num, err := strconv.ParseUint(value[0:strings.Index(value, ":")], 10, 16)

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
