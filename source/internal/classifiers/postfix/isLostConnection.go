package postfix

import "strings"

func isLostConnection(message string) bool {

	var result bool = false

	commands := []string{
		"CONNECT",
		"STARTTLS",
		"EHLO",
		"AUTH",
		"UNKNOWN",
	}

	for c := 0; c < len(commands); c++ {

		if strings.HasPrefix(message, "lost connection after "+commands[c]+" from") {
			result = true
			break
		}

	}

	return result

}
