package versions

import "strings"

func Constraint(operator string, value string) (string, bool) {

	version, ok := Normalize(value)

	if ok == false {
		return "", false
	}

	operator = strings.TrimSpace(operator)

	switch operator {

	case "", "=", "==", "===":
		return "= " + version, true

	case ">=", ">", "<=", "<":
		return operator + " " + version, true

	}

	return "", false

}
