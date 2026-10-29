package strings

func Truncate(value string, length int) string {

	if length <= 0 {
		return ""
	}

	if len(value) > length {
		return value[0:length]
	}

	return value

}
