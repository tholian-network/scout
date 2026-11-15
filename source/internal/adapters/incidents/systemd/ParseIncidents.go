package systemd

import "tholian-core/structs"
import "tholian-core/types"
import "tholian-scout/internal/classifiers"
import "strings"

func ParseIncidents(buffer []byte, system string) []structs.Incident {

	var result []structs.Incident

	lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

	for l := 0; l < len(lines); l++ {

		if strings.HasPrefix(lines[l], " ") {
			// Do Nothing
		} else if lines[l] == "-- No entries --" {
			// Do Nothing
		} else if strings.HasPrefix(lines[l], "-- Boot") {
			// Do Nothing
		} else if strings.HasPrefix(lines[l], "-- Journal") {
			// Do Nothing
		} else if strings.HasPrefix(lines[l], "-- ") {
			// Do Nothing
		} else if len(lines[l]) > 16 && strings.Contains(lines[l], " ") && strings.Contains(lines[l], ": ") {

			meta := lines[l][0:strings.Index(lines[l], ": ")]
			message := lines[l][strings.Index(lines[l], ": ")+2:]

			datetime := types.ToDatetime(meta[0:15])
			hostname := strings.TrimSpace(strings.Split(meta[16:], " ")[0])
			service := strings.Split(meta[16:], " ")[1]

			if datetime.IsValid() && datetime.IsFuture() {
				datetime.Year -= 1
			}

			if strings.Contains(service, "[") && strings.HasSuffix(service, "]") {
				service = service[0:strings.Index(service, "[")]
			}

			if strings.HasPrefix(service, "(") && strings.HasSuffix(service, ")") {
				service = service[1 : len(service)-1]
			}

			if strings.HasPrefix(message, "pam_unix(") && strings.Contains(message, "):") {
				service = "pam"
			} else if strings.HasPrefix(message, "pam_faillock(") && strings.Contains(message, "):") {
				service = "pam"
			} else if strings.HasPrefix(message, "pam_systemd_home(") && strings.Contains(message, "):") {
				service = "systemd-homed"
			}

			if l < len(lines)-1 && strings.HasPrefix(lines[l+1], " ") {

				seek := 1

				for strings.HasPrefix(lines[l+seek], " ") {

					if l+seek < len(lines)-1 {
						seek++
					} else {
						break
					}

				}

				message_lines := make([]string, 0)
				message_lines = append(message_lines, message)

				for s := l + 1; s < l+seek; s++ {
					message_line := strings.TrimSpace(strings.ReplaceAll(lines[s], "\t", " "))
					message_lines = append(message_lines, message_line)
				}

				message = strings.Join(message_lines, "\n")
				l += seek - 1

			}

			classify, ok1 := classifiers.Classifiers[service]

			if ok1 == true {

				incident := classify(system, datetime.String(), hostname, service, message)

				if incident.IsValid() {
					result = append(result, incident)
				}

			} else {

				// console.Warn("systemd: Cannot classify messages from service \"" + service + "\"")

			}

		}

	}

	return result

}
