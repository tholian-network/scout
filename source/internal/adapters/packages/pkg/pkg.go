package pkg

import "os"
import "os/exec"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/sbin/pkg")
	_, err2 := os.Stat("/usr/local/sbin/pkg")
	_, err3 := os.Stat("/var/db/pkg")
	_, err4 := exec.LookPath("pkg")

	if (err1 == nil || err2 == nil || err4 == nil) && err3 == nil {
		SUPPORTED = true
	}

}
