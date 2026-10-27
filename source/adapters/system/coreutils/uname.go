package coreutils

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/usr/bin/uname")

	if err == nil {
		SUPPORTED = true
	}

}
