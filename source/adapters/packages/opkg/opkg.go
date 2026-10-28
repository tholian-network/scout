package opkg

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/bin/opkg")
	_, err2 := os.Stat("/bin/opkg")
	_, err3 := os.Stat("/usr/lib/opkg/status")

	if (err1 == nil || err2 == nil) && err3 == nil {
		SUPPORTED = true
	}

}
