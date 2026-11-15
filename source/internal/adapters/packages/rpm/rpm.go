package rpm

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/usr/bin/rpm")

	if err == nil {
		SUPPORTED = true
	}

}
