package procstat

import "os"
import "os/exec"
import "strconv"
import "strings"
import "time"

func CollectProcessStartTimes() map[uint]string {

	collected := make(map[uint]string)

	if SUPPORTED == true {

		cmd := exec.Command("ps", "-ax", "-o", "pid,lstart")
		cmd.Env = append(os.Environ(), "LC_ALL=C")

		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(buffer)), "\r\n", "\n"), "\n")

			for l := 0; l < len(lines); l++ {

				fields := strings.Fields(lines[l])

				if fields[0] != "PID" && len(fields) >= 2 {

					pid, err := strconv.ParseUint(fields[0], 10, 64)

					if err == nil {

						value := strings.Join(fields[1:], " ")

						if parsed, err := time.Parse("Mon Jan _2 15:04:05 2006", value); err == nil {
							value = parsed.Format("2006-01-02 15:04:05")
						}

						collected[uint(pid)] = value

					}

				}

			}

		}

	}

	return collected

}
