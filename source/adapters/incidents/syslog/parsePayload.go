package syslog

import "strings"

func parsePayload(payload string) (string, string) {

	payload = strings.TrimSpace(payload)

	for payload != "" {

		fields := strings.Fields(payload)

		if len(fields) == 0 || isPriority(fields[0]) == false {
			break
		}

		payload = strings.TrimSpace(strings.TrimPrefix(payload, fields[0]))

	}

	index := strings.Index(payload, ": ")

	if index != -1 {

		service := strings.TrimSpace(strings.ToLower(payload[0:index]))
		message := strings.TrimSpace(payload[index+2:])

		if strings.Contains(service, "[") {
			service = strings.TrimSpace(service[0:strings.Index(service, "[")])
		}

		if strings.HasPrefix(service, "(") && strings.HasSuffix(service, ")") {
			service = strings.TrimSpace(service[1 : len(service)-1])
		}

		return service, message

	}

	return "", ""

}
