package datetime

func IsMonth(value string) bool {

	switch value {
	case "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec":
		return true
	}

	return false

}

