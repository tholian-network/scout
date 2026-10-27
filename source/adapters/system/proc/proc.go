package proc

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/proc")

	if err1 == nil {
		SUPPORTED = true
	}

}
