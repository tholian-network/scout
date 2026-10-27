package postfix

import "tholian-core/structs"
import "strings"

func Classify(system string, datetime string, hostname string, service string, message string) structs.Incident {

	var incident structs.Incident

	if datetime != "" && service == "postfix" {

	} else if datetime != "" && service == "postfix/anvil" {

		// Do Nothing

	} else if datetime != "" && service == "postfix/bounce" {

		// Do Nothing

	} else if datetime != "" && service == "postfix/cleanup" {

		// Do Nothing

	} else if datetime != "" && service == "postfix/lmtp" {

		// Do Nothing

	} else if datetime != "" && service == "postfix/master" {

		// Do Nothing

	} else if datetime != "" && service == "postfix/pickup" {

		// Do Nothing

	} else if datetime != "" && service == "postfix/postmap" {

		// Do Nothing

	} else if datetime != "" && service == "postfix/qmgr" {

		// Do Nothing

	} else if datetime != "" && service == "postfix/scache" {

		// Do Nothing

	} else if datetime != "" && service == "postfix/smtpd" {

		if strings.HasPrefix(message, "warning: hostname") && strings.Contains(message, "does not resolve to address") {

			tmp := splitLine(message[17:], " ")
			host := parseHost(tmp[len(tmp)-1])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if strings.HasPrefix(message, "warning: non-SMTP command from") {

			tmp := splitLine(message[30:], " ")
			host := parseHost(tmp[0])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if strings.HasPrefix(message, "SSL_accept error from") {

			tmp := splitLine(message[21:], " ")
			host := parseHost(tmp[0])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if isLostConnection(message) {

			tmp := splitLine(message[21:], " ")
			host := parseHost(tmp[len(tmp)-1])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if strings.HasPrefix(message, "disconnect from") && (strings.Contains(message, "commands=0/0") || strings.Contains(message, "unknown=0/1") || strings.Contains(message, "auth=0/1")) {

			tmp := splitLine(message[15:], " ")
			host := parseHost(tmp[0])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		}

	} else if datetime != "" && service == "postfix/smtps/smtpd" {

		if strings.HasPrefix(message, "warning: hostname") && strings.Contains(message, "does not resolve to address") {

			tmp := splitLine(message[17:], " ")
			host := parseHost(tmp[len(tmp)-1])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if strings.HasPrefix(message, "warning:") && strings.Contains(message, "SASL LOGIN authentication failed") {

			tmp := splitLine(message[8:], " ")
			host := parseHost(tmp[0])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if strings.HasPrefix(message, "NOQUEUE: reject: RCPT from") && strings.Contains(message, "Relay access denied") {

			tmp := splitLine(message[26:], " ")
			host := parseHost(tmp[0])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if strings.HasPrefix(message, "SSL_accept error from") {

			tmp := splitLine(message[21:], " ")
			host := parseHost(tmp[0])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if strings.HasPrefix(message, "improper command pipelining after") && strings.Contains(message, "from") {

			tmp := splitLine(message[33:], " ")
			host := parseHost(tmp[len(tmp)-1])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if isLostConnection(message) {

			tmp := splitLine(message[21:], " ")
			host := parseHost(tmp[len(tmp)-1])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		} else if strings.HasPrefix(message, "disconnect from") && (strings.Contains(message, "commands=0/0") || strings.Contains(message, "unknown=0/1") || strings.Contains(message, "auth=0/1")) {

			tmp := splitLine(message[15:], " ")
			host := parseHost(tmp[0])

			if host != "" {
				incident = toReconIncident(system, datetime, host)
			}

		}

	}

	return incident

}
