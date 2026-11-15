package proc

import "os"
import "strconv"
import "strings"

func getBootTime() uint64 {

	buffer, err := os.ReadFile("/proc/stat")

	if err == nil {

		lines := strings.Split(string(buffer), "\n")

		for l := 0; l < len(lines); l++ {

			line := strings.TrimSpace(lines[l])

			if strings.HasPrefix(line, "btime ") {

				value, err := strconv.ParseUint(strings.TrimSpace(line[6:]), 10, 64)

				if err == nil {
					return value
				}

			}

		}

	}

	return 0

}
