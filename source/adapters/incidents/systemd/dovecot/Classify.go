package dovecot

import "tholian-core/structs"
import "tholian-core/types"
import "strings"

func Classify(system string, datetime string, hostname string, service string, message string) structs.Incident {

	var incident structs.Incident

	if datetime != "" && service == "dovecot" {

		if strings.HasPrefix(message, "auth-worker(") && strings.Contains(message, "):") {

			message := strings.TrimSpace(strings.Join(strings.Split(message[12:], "):")[1:], "):"))

			if strings.HasPrefix(message, "Debug:") {
				message = strings.TrimSpace(message[6:])
			}

			if strings.Contains(message, "auth-worker<") && strings.Contains(message, ">: ") {

				chunk := message[strings.Index(message, ">: ")+3:]

				if strings.HasPrefix(chunk, "pam(") && strings.Contains(chunk, "): ") && strings.Contains(chunk, "pam_authenticate() failed") {

					username := chunk[4:strings.Index(chunk, ",")]

					if types.IsIPv4(username) {

						// Do Nothing

					} else {

						ip := chunk[strings.Index(chunk, ",")+1 : strings.Index(chunk, "): ")]
						password := ""

						if strings.Contains(chunk, "(given password: ") && strings.HasSuffix(chunk, ")") {
							password = chunk[strings.Index(chunk, "(given password: ")+17 : len(chunk)-1]
						}

						incident = toReconIncident(system, datetime, ip, username, password)

					}

				} else if chunk == "Handling USER request" {

					// Do Nothing

				} else if chunk == "Handling PASSV request" {

					// Do Nothing

				} else if strings.HasPrefix(chunk, "Finished:") {

					// Do Nothing

				}

			}

		}

	}

	return incident

}
