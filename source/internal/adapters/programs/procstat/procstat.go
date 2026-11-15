package procstat

import "os"
import "os/exec"
import "strings"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/bin/procstat")
	_, err2 := os.Stat("/usr/sbin/procstat")
	_, err3 := os.Stat("/bin/procstat")

	if err1 == nil || err2 == nil || err3 == nil {

		cmd := exec.Command("uname", "-s")
		buffer, err := cmd.Output()

		if err == nil && strings.TrimSpace(string(buffer)) == "FreeBSD" {
			SUPPORTED = true
		}

	}

}
