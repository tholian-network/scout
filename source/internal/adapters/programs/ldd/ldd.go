package ldd

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/usr/bin/ldd")

	if err == nil {
		SUPPORTED = true
	}

}
