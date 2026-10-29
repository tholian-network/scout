package procstat

import "strconv"
import "strings"

type process_info struct {
	PID  uint
	UID  uint16
	Name string
}

func parseProcessStatus(buffer string) []process_info {

	var collected []process_info

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		fields := strings.Fields(lines[l])

		if fields[0] != "PID" && len(fields) >= 3 {

			pid, err1 := strconv.ParseUint(fields[0], 10, 64)
			uid, err2 := strconv.ParseUint(fields[1], 10, 16)

			if err1 == nil && err2 == nil {

				collected = append(collected, process_info{
					PID:  uint(pid),
					UID:  uint16(uid),
					Name: strings.Join(fields[2:], " "),
				})

			}

		}

	}

	return collected

}
