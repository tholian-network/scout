package pacman

import "os"

var OPTIMIZED bool = false
var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/usr/bin/pacman")

	if err == nil {
		OPTIMIZED = true
		SUPPORTED = true
	}

}
