package proc

import "tholian-core/types"
import "os"
import "strconv"
import "time"

func getProcessStartedAt(pid uint) *types.Datetime {

	var result *types.Datetime

	buffer, err := os.ReadFile("/proc/" + strconv.FormatUint(uint64(pid), 10) + "/stat")

	if err == nil {

		boot_time := getBootTime()
		process_ticks := getProcessTicks(buffer)

		if boot_time > 0 && process_ticks > 0 {

			seconds := (boot_time) + (process_ticks / 100)
			datetime := types.ToDatetime(time.Unix(int64(seconds), 0).Format("2006-01-02 15:04:05"))

			if datetime.IsValid() {
				result = &datetime
			}

		}

	}

	return result

}
