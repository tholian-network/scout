package datetime

func IsISODate(value string) bool {

	if len(value) < 10 {
		return false
	}

	if value[4] != '-' || value[7] != '-' {
		return false
	}

	return IsYear(value[0:4]) && IsISOMonth(value[5:7]) && IsDay(value[8:10])

}
