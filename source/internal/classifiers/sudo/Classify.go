package sudo

import "tholian-core/matchers"
import "tholian-core/structs"
import utils_os "tholian-scout/internal/utils/os"
import "strings"

func Classify(system string, datetime string, hostname string, service string, message string) structs.Incident {

	var incident structs.Incident

	if datetime != "" && service == "sudo" {

		username := strings.TrimSpace(strings.Split(message, " : ")[0])
		message := strings.TrimSpace(strings.Join(strings.Split(message, " : ")[1:], " : "))

		if strings.Contains(message, "incorrect password attempts") || strings.Contains(message, "incorrect password attempt") {

			sudoexec := NewSuperUserExecution(message)

			if sudoexec.User == "root" && utils_os.IsDaemonUsername(username) {

				incident = toPersistIncident(system, datetime, username)

				program := matchers.NewProgram()
				program.SetName(sudoexec.ToProgramName())
				program.SetCommand(sudoexec.Command)
				program.SetArguments(sudoexec.Arguments)

			}

		}

	} else if datetime != "" && service == "su" {

		if strings.HasPrefix(message, "FAILED SU") {

			message := strings.TrimSpace(message[9:])

			if strings.HasPrefix(message, "(to ") && strings.Contains(message, ") ") {

				username := strings.TrimSpace(message[strings.Index(message, ") ")+2:])

				if strings.Contains(username, " ") {
					username = username[0:strings.Index(username, " ")]
				}

				suexec_username := strings.TrimSpace(message[3:strings.Index(message, ") ")])

				if suexec_username == "root" && utils_os.IsDaemonUsername(username) {
					incident = toPersistIncident(system, datetime, username)
				}

			}

		}

	}

	return incident

}
