package apk

import "os"

var OPTIMIZED bool = false
var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/sbin/apk")
	_, err2 := os.Stat("/lib/apk/db/installed")

	if err1 == nil && err2 == nil {
		SUPPORTED = true
		OPTIMIZED = true
	} else if err1 == nil {
		SUPPORTED = true
	}

}
