package dnf

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/usr/bin/dnf")

	if err == nil {
		SUPPORTED = true
	}

}
