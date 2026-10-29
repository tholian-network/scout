package procstat

import "strconv"
import "strings"

func parseProcstatBinary(buffer string) map[uint]string {

	collected := make(map[uint]string)
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		fields := strings.Fields(lines[l])

		if fields[0] != "PID" && len(fields) >= 4 {

			pid, err := strconv.ParseUint(fields[0], 10, 64)

			if err == nil {
				collected[uint(pid)] = strings.Join(fields[3:], " ")
			}

		}

	}

	return collected

}
