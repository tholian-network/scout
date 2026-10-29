package shadow

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/etc/passwd")
	_, err2 := os.Stat("/etc/group")
	_, err3 := os.Stat("/etc/shadow")

	if err1 == nil && err2 == nil && err3 == nil {
		SUPPORTED = true
	}

}
