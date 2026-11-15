package zypper

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/usr/bin/zypper")

	if err == nil {
		SUPPORTED = true
	}

}
