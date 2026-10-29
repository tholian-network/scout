package dropbear

import "tholian-core/structs"
import "strings"

func Classify(system string, datetime string, hostname string, service string, message string) structs.Incident {

	var incident structs.Incident

	if datetime != "" && service == "dropbear" {

		if strings.HasPrefix(message, "Bad password attempt for") {

			username, host := parseAuth(message)

			if host != "" {
				incident = toReconIncident(system, datetime, host, username)
			}

		} else if strings.HasPrefix(message, "Login attempt for nonexistent user") {

			_, host := parseAuth(message)

			if host != "" {
				incident = toReconIncident(system, datetime, host, "")
			}

		} else if strings.HasPrefix(message, "Bad protocol version identification") {

			_, host := parseAuth(message)

			if host != "" {
				incident = toReconIncident(system, datetime, host, "")
			}

		} else if strings.HasPrefix(message, "Password auth succeeded for") {

			username, host := parseAuth(message)

			if host != "" {
				incident = toAccessIncident(system, datetime, host, username)
			}

		} else if strings.HasPrefix(message, "Pubkey auth succeeded for") {

			username, host := parseAuth(message)

			if host != "" {
				incident = toAccessIncident(system, datetime, host, username)
			}

		}

	}

	return incident

}
