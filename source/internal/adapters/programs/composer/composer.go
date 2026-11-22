package composer

import "os"
import "os/exec"

var SUPPORTED bool = false

func init() {

	_, err1 := exec.LookPath("composer")
	_, err2 := exec.LookPath("php")
	_, err3 := os.Stat("/usr/bin/php")
	_, err4 := os.Stat("/usr/local/bin/php")

	if err1 == nil || err2 == nil || err3 == nil || err4 == nil {
		SUPPORTED = true
	}

}
