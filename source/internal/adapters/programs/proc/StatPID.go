package proc

import "os"
import "strconv"
import "strings"

func StatPID() uint64 {

	var result uint64 = 0

	buffer, err1 := os.ReadFile("/proc/stat")

	if err1 == nil {

		lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

		for l := 0; l < len(lines); l++ {

			line := lines[l]

			if strings.HasPrefix(line, "processes ") {

				num, err2 := strconv.ParseUint(strings.TrimSpace(line[10:]), 10, 64)

				if err2 == nil {
					result = num
					break
				}

			}

		}

	}

	return result

}
