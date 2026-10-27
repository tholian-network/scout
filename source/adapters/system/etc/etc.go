package etc

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/etc/os-release")
	_, err2 := os.Stat("/etc/machine-id")

	if err1 == nil && err2 == nil {
		SUPPORTED = true
	}

}
