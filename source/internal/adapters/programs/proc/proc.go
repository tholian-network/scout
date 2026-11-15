package proc

import "os"
import "os/exec"
import "strings"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/proc")

	cmd := exec.Command("uname", "-s")
	buffer, _ := cmd.Output()

	kernel := strings.TrimSpace(string(buffer))

	if err1 == nil && kernel == "Linux" {
		SUPPORTED = true
	}

}
