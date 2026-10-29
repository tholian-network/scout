package datetime

func IsYear(value string) bool {

	if len(value) != 4 {
		return false
	}

	for y := 0; y < len(value); y++ {

		if value[y] < '0' || value[y] > '9' {
			return false
		}

	}

	return true

}
