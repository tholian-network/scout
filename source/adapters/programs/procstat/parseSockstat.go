package procstat

import "strconv"
import "strings"

type socket_info struct {
	PID         uint
	Protocol    string
	Local       string
	LocalPort   uint16
	Foreign     string
	ForeignPort uint16
}

func parseSockstat(buffer string) []socket_info {

	var collected []socket_info

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		fields := strings.Fields(lines[l])

		if len(fields) < 7 || fields[0] == "USER" {
			continue
		}

		pid, err := strconv.ParseUint(fields[2], 10, 64)

		if err != nil {
			continue
		}

		protocol := strings.ToLower(fields[4])

		if protocol == "tcp4" || protocol == "tcp6" || protocol == "tcp46" {
			protocol = "tcp"
		} else if protocol == "udp4" || protocol == "udp6" || protocol == "udp46" {
			protocol = "udp"
		} else {
			continue
		}

		local, local_port := parseSockstatAddress(fields[5])
		foreign, foreign_port := parseSockstatAddress(fields[6])

		collected = append(collected, socket_info{
			PID:         uint(pid),
			Protocol:    protocol,
			Local:       local,
			LocalPort:   local_port,
			Foreign:     foreign,
			ForeignPort: foreign_port,
		})

	}

	return collected

}

