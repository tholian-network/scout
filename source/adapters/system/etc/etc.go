package etc

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/etc/os-release")

	if err == nil {
		SUPPORTED = true
	}

}
