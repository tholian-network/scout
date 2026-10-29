package kenv

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/bin/kenv")
	_, err2 := os.Stat("/usr/sbin/kenv")
	_, err3 := os.Stat("/bin/kenv")
	_, err4 := os.Stat("/sbin/kenv")

	if err1 == nil || err2 == nil || err3 == nil || err4 == nil {
		SUPPORTED = true
	}

}
