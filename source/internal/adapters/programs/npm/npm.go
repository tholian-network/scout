package npm

import "os"
import "os/exec"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/lib/node_modules")
	_, err2 := os.Stat("/usr/local/lib/node_modules")
	_, err3 := exec.LookPath("npm")
	_, err4 := exec.LookPath("node")

	if err1 == nil || err2 == nil || err3 == nil || err4 == nil {
		SUPPORTED = true
	}

}
