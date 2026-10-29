package sshd

import "tholian-core/structs"
import "strings"

func Classify(system string, datetime string, hostname string, service string, message string) structs.Incident {

	var incident structs.Incident

	if datetime != "" && service == "sshd" {

		if strings.HasPrefix(message, "Failed password for invalid user") && strings.Contains(message, " from ") {

			tmp := strings.Split(strings.TrimSpace(message[32:]), " ")

			if len(tmp) >= 3 && tmp[1] == "from" {
				incident = toReconIncident(system, datetime, tmp[2], tmp[0])
			}

		} else if strings.HasPrefix(message, "Failed password for") {

			tmp := strings.Split(strings.TrimSpace(message[19:]), " ")

			if len(tmp) >= 3 && tmp[1] == "from" {
				incident = toReconIncident(system, datetime, tmp[2], tmp[0])
			}

		} else if strings.HasPrefix(message, "Invalid user") && strings.Contains(message, " from ") {

			tmp := strings.Split(strings.TrimSpace(message[12:]), " ")

			if len(tmp) >= 3 && tmp[1] == "from" {
				incident = toReconIncident(system, datetime, tmp[2], tmp[0])
			} else if len(tmp) >= 2 && tmp[0] == "from" {
				incident = toReconIncident(system, datetime, tmp[1], "")
			}

		} else if strings.HasPrefix(message, "error: maximum authentication attempts exceeded for invalid user") {

			tmp := strings.Split(strings.TrimSpace(message[64:]), " ")

			if len(tmp) >= 3 && tmp[1] == "from" {
				incident = toReconIncident(system, datetime, tmp[2], tmp[0])
			}

		} else if strings.HasPrefix(message, "error: maximum authentication attempts exceeded for") {

			tmp := strings.Split(strings.TrimSpace(message[51:]), " ")

			if len(tmp) >= 3 && tmp[1] == "from" {
				incident = toReconIncident(system, datetime, tmp[2], tmp[0])
			}

		} else if strings.HasPrefix(message, "fatal: Unable to negotiate with") && strings.Contains(message, "no matching cipher found") {

			tmp := strings.Split(strings.TrimSpace(message[31:]), " ")

			if len(tmp) >= 1 {
				incident = toReconIncident(system, datetime, tmp[0], "")
			}

		} else if strings.HasPrefix(message, "Unable to negotiate with") && strings.Contains(message, "no matching host key type found") {

			tmp := strings.Split(strings.TrimSpace(message[24:]), " ")

			if len(tmp) >= 1 {
				incident = toReconIncident(system, datetime, tmp[0], "")
			}

		} else if strings.HasPrefix(message, "Unable to negotiate with") && strings.Contains(message, "no matching key exchange method found") {

			tmp := strings.Split(strings.TrimSpace(message[24:]), " ")

			if len(tmp) >= 1 {
				incident = toReconIncident(system, datetime, tmp[0], "")
			}

		} else if strings.HasPrefix(message, "Bad protocol version identification") && strings.Contains(message, " from ") {

			index := strings.LastIndex(message, "'")

			if index != -1 {

				tmp := strings.Split(strings.TrimSpace(message[index+1:]), " ")

				if len(tmp) >= 2 && tmp[0] == "from" {
					incident = toReconIncident(system, datetime, tmp[1], "")
				}

			}

		} else if strings.HasPrefix(message, "Accepted publickey for") {

			tmp := strings.Split(strings.TrimSpace(message[22:]), " ")

			if len(tmp) >= 3 && tmp[1] == "from" {
				incident = toAccessIncident(system, datetime, tmp[2], tmp[0])
			}

		} else if strings.HasPrefix(message, "Accepted password for") {

			tmp := strings.Split(strings.TrimSpace(message[21:]), " ")

			if len(tmp) >= 3 && tmp[1] == "from" {
				incident = toAccessIncident(system, datetime, tmp[2], tmp[0])
			}

		}

	}

	return incident

}
