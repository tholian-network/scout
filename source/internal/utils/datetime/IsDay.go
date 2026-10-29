package datetime

func IsDay(value string) bool {

	if len(value) == 0 || len(value) > 2 {
		return false
	}

	for d := 0; d < len(value); d++ {

		if value[d] < '0' || value[d] > '9' {
			return false
		}

	}

	return true

}

