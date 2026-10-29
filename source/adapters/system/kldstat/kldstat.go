package kldstat

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/bin/kldstat")
	_, err2 := os.Stat("/usr/sbin/kldstat")
	_, err3 := os.Stat("/sbin/kldstat")

	if err1 == nil || err2 == nil || err3 == nil {
		SUPPORTED = true
	}

}
