package usbconfig

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/sbin/usbconfig")
	_, err2 := os.Stat("/usr/bin/usbconfig")
	_, err3 := os.Stat("/sbin/usbconfig")

	if err1 == nil || err2 == nil || err3 == nil {
		SUPPORTED = true
	}

}
