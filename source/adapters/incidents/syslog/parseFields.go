package syslog

import "strings"
import utils_datetime "tholian-scout/internal/utils/datetime"

func parseFields(line string) (string, string, string) {

	fields := strings.Fields(line)

	if len(fields) >= 3 {

		index := 0
		timestamp := ""

		if utils_datetime.IsWeekday(fields[0]) && len(fields) >= 7 && utils_datetime.IsMonth(fields[1]) && utils_datetime.IsDay(fields[2]) && utils_datetime.IsTime(fields[3]) && utils_datetime.IsYear(fields[4]) {

			timestamp = strings.Join(fields[0:5], " ")
			index = 5

		} else if utils_datetime.IsMonth(fields[0]) && len(fields) >= 6 && utils_datetime.IsDay(fields[1]) && utils_datetime.IsTime(fields[2]) && utils_datetime.IsYear(fields[3]) {

			timestamp = strings.Join(fields[0:4], " ")
			index = 4

		} else if utils_datetime.IsMonth(fields[0]) && utils_datetime.IsDay(fields[1]) && utils_datetime.IsTime(fields[2]) {

			timestamp = strings.Join(fields[0:3], " ")
			index = 3

		} else if utils_datetime.IsISODateTime(fields[0]) {

			timestamp = fields[0]
			index = 1

		} else if utils_datetime.IsISODate(fields[0]) && len(fields) >= 4 && utils_datetime.IsTime(fields[1]) {

			timestamp = fields[0] + " " + fields[1]
			index = 2

		}

		if index < len(fields) {

			hostname := fields[index]
			payload := strings.Join(fields[index+1:], " ")

			return timestamp, hostname, payload

		}

	}

	return "", "", ""

}
