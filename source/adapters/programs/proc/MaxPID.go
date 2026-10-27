package proc

import "os"
import "strconv"
import "strings"

var MaxPID uint64

func init() {

	buffer, err1 := os.ReadFile("/proc/sys/kernel/pid_max")

	if err1 == nil {

		num, err2 := strconv.ParseUint(strings.TrimSpace(string(buffer)), 10, 64)

		if err2 == nil {
			MaxPID = uint64(num)
		} else {
			MaxPID = uint64(32768)
		}

	} else {
		MaxPID = uint64(32768)
	}

}
