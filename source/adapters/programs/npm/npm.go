package npm

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/bin/npm")
	_, err2 := os.Stat("/usr/lib/node_modules")

	if err1 == nil && err2 == nil {
		SUPPORTED = true
	}

}
