package syslog

import "tholian-core/structs"
import "tholian-scout/internal/classifiers"
import "tholian-core/types"
import "strings"

func ParseIncidents(buffer []byte, system string) []structs.Incident {

	var result []structs.Incident

	lines := strings.Split(strings.ReplaceAll(string(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(strings.TrimRight(lines[l], "\r"))

		if line != "" {

			if strings.HasPrefix(line, "<") && strings.Contains(line, ">") {
				line = strings.TrimSpace(line[strings.Index(line, ">")+1:])
			}

			timestamp, hostname, payload := parseFields(line)

			if timestamp != "" && hostname != "" && payload != "" {

				service, message := parsePayload(payload)

				if service != "" && message != "" {

					datetime := types.ToDatetime(timestamp)

					if datetime.IsValid() {

						classify, ok := classifiers.Classifiers[service]

						if ok == true {

							incident := classify(system, datetime.String(), hostname, service, message)

							if incident.IsValid() == true {
								result = append(result, incident)
							}

						}

					}

				}

			}

		}

	}

	return result

}
