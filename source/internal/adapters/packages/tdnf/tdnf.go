package tdnf

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/usr/bin/tdnf")

	if err == nil {
		SUPPORTED = true
	}

}
