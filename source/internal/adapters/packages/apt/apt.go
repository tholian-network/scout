package apt

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/bin/apt")
	_, err2 := os.Stat("/usr/bin/apt-cache")

	if err1 == nil && err2 == nil {
		SUPPORTED = true
	}

}
