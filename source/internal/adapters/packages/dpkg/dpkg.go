package dpkg

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/bin/dpkg")
	_, err2 := os.Stat("/bin/dpkg")
	_, err3 := os.Stat("/var/lib/dpkg/status")

	if (err1 == nil || err2 == nil) && err3 == nil {
		SUPPORTED = true
	}

}
