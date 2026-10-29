package pciconf

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/sbin/pciconf")
	_, err2 := os.Stat("/usr/bin/pciconf")
	_, err3 := os.Stat("/sbin/pciconf")

	if err1 == nil || err2 == nil || err3 == nil {
		SUPPORTED = true
	}

}
