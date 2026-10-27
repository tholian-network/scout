package systemd

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/usr/lib/sysusers.d")

	if err == nil {
		SUPPORTED = true
	}

}
