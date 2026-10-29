package datetime

func IsWeekday(value string) bool {

	switch value {
	case "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun":
		return true
	}

	return false

}
