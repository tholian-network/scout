package pam

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import utils_os "tholian-scout/internal/utils/os"
import "strings"

func Classify(system string, datetime string, hostname string, service string, message string) structs.Incident {

	var incident structs.Incident

	if datetime != "" && service == "pam" {

		if strings.HasPrefix(message, "pam_unix(") && strings.Contains(message, "):") {

			program_name := strings.TrimSpace(strings.Split(message[9:], "):")[0])
			message := strings.TrimSpace(strings.Split(message[9:], "):")[1])

			if strings.HasSuffix(program_name, ":auth") {
				program_name = strings.Split(program_name, ":")[0]
			} else if strings.HasSuffix(program_name, ":session") {
				program_name = strings.Split(program_name, ":")[0]
			}

			// su-l has a special program name
			if program_name == "su-l" {
				program_name = "su"
			}

			if strings.HasPrefix(message, "authentication failure;") {

				auth_failure := NewAuthFailure(message)

				if auth_failure.RemoteHost != "" && auth_failure.RemoteUser != "" {

					if types.IsIPv4(auth_failure.RemoteHost) {
						// Do Nothing
					} else {
						incident = toReconIncident(system, datetime, auth_failure.RemoteHost, auth_failure.RemoteUser)
					}

				} else if auth_failure.TTY != "" {

					if auth_failure.UID != 0 && auth_failure.EUID == 0 {

						if utils_os.IsDaemonUsername(auth_failure.User) {

							incident = toPersistIncident(system, datetime, auth_failure.User)

							program := matchers.NewProgram()
							program.SetName(program_name)
							incident.AddProgram(program)

						} else if auth_failure.User == "" {

							incident = toPersistIncident(system, datetime, "root")

							program := matchers.NewProgram()
							program.SetName(program_name)
							incident.AddProgram(program)

						}

					} else if auth_failure.UID == 0 && auth_failure.EUID == 0 {

						if utils_os.IsDaemonUsername(auth_failure.User) {

							incident = toPersistIncident(system, datetime, auth_failure.User)

							program := matchers.NewProgram()
							program.SetName(program_name)
							incident.AddProgram(program)

						} else {

							incident = toPersistIncident(system, datetime, "root")

							program := matchers.NewProgram()
							program.SetName(program_name)
							incident.AddProgram(program)

						}

					}

				}

			} else if strings.HasPrefix(message, "check pass; user unknown") {

				// Do Nothing

			} else if strings.HasPrefix(message, "conversation failed") {

				// Do Nothing

			} else if strings.HasPrefix(message, "auth could not identify password for") {

				// Do Nothing

			} else if strings.HasPrefix(message, "session opened for user") {

				// Do Nothing

			} else if strings.HasPrefix(message, "session closed for user") {

				// Do Nothing

			}

		}

	}

	return incident

}
