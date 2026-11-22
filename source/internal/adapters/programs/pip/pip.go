package pip

import "os"
import "os/exec"

var SUPPORTED bool = false

func init() {

	_, err1 := exec.LookPath("python3")
	_, err2 := exec.LookPath("python")
	_, err3 := exec.LookPath("pip3")
	_, err4 := exec.LookPath("pip")
	_, err5 := os.Stat("/usr/bin/python3")
	_, err6 := os.Stat("/usr/local/bin/python3")

	if err1 == nil || err2 == nil || err3 == nil || err4 == nil || err5 == nil || err6 == nil {
		SUPPORTED = true
	}

}
