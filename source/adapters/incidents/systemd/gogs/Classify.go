package gogs

import "tholian-core/structs"
import "strings"

func Classify(system string, datetime string, hostname string, service string, message string) structs.Incident {

	var incident structs.Incident

	if datetime != "" && service == "gogs" {

		if strings.Contains(message, "http: TLS handshake error from") {

			tmp := message[strings.Index(message, "http: TLS handshake error from")+30:]

			if strings.Contains(tmp, ": ") {

				chunks := strings.Split(tmp, ": ")

				if len(chunks) == 2 {

					ip := strings.TrimSpace(chunks[0])
					message := strings.TrimSpace(chunks[1])

					if strings.Contains(ip, ":") {
						ip = strings.Split(ip, ":")[0]
					}

					if message == "EOF" || message == "unexpected EOF" {
						incident = toReconIncident(system, datetime, ip)
					}

				} else if len(chunks) == 3 {

					ip := strings.TrimSpace(chunks[0])

					if strings.Contains(ip, ":") {
						ip = strings.Split(ip, ":")[0]
					}

					if chunks[1] == "tls" {

						message := strings.TrimSpace(chunks[2])

						if message == "first record does not look like a TLS handshake" {
							incident = toReconIncident(system, datetime, ip)
						} else if message == "no cipher suite supported by both client and server" {
							incident = toReconIncident(system, datetime, ip)
						} else if message == "unsupported SSLv2 handshake received" {
							incident = toReconIncident(system, datetime, ip)
						} else if message == "no ECDHE curve supported by both client and server" {
							incident = toReconIncident(system, datetime, ip)
						} else if strings.HasPrefix(message, "client requested unsupported application protocols") {
							incident = toReconIncident(system, datetime, ip)
						}

					}

				} else if len(chunks) == 4 {

					ip := strings.TrimSpace(chunks[0])

					if strings.Contains(ip, ":") {
						ip = strings.Split(ip, ":")[0]
					}

					if chunks[1] == "tls" {

						message := strings.TrimSpace(chunks[2])

						if message == "client offered only unsupported versions" {
							incident = toReconIncident(system, datetime, ip)
						}

					} else if strings.HasPrefix(chunks[1], "read tcp") {

						message := strings.TrimSpace(chunks[3])

						if message == "connection reset by peer" {
							incident = toReconIncident(system, datetime, ip)
						} else if message == "connection timed out" {
							incident = toReconIncident(system, datetime, ip)
						}

					}

				}

			}

		}

	}

	return incident

}
