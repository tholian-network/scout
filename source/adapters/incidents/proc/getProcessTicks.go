package proc

import "strconv"
import "strings"

func getProcessTicks(buffer []byte) uint64 {

	stat := strings.TrimSpace(string(buffer))
	index := strings.LastIndex(stat, ")")

	if index != -1 {

		fields := strings.Fields(stat[index+1:])

		// Field 22 of /proc/<pid>/stat. Field 3 is the first field after the
		// process name, so starttime is at index 19 here.
		if len(fields) > 19 {

			value, err := strconv.ParseUint(fields[19], 10, 64)

			if err == nil {
				return value
			}

		}

	}

	return 0

}
