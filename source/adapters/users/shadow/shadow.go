package shadow

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/etc/passwd")
	_, err2 := os.Stat("/etc/group")

	if err1 == nil && err2 == nil {
		SUPPORTED = true
	}

}
