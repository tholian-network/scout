package systemd

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/usr/bin/journalctl")

	if err == nil {
		SUPPORTED = true
	}

}
