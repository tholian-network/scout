package df

import "os"
import "strconv"
import "strings"

var SUPPORTED bool = false

var suffixes map[string]uint64 = map[string]uint64{
	"KB": uint64(1000),
	"K":  uint64(1024),
	"MB": uint64(1000 * 1000),
	"M":  uint64(1024 * 1024),
	"GB": uint64(1000 * 1000 * 1000),
	"G":  uint64(1024 * 1024 * 1024),
	"TB": uint64(1000 * 1000 * 1000 * 1000),
	"T":  uint64(1024 * 1024 * 1024 * 1024),
	"PB": uint64(1000 * 1000 * 1000 * 1000 * 1000),
	"P":  uint64(1024 * 1024 * 1024 * 1024 * 1024),
	// Needs hi/low representation, golang does not have uint128
	// "EB": uint64(1000 * 1000 * 1000 * 1000 * 1000 * 1000),
	// "E": uint64(1024 * 1024 * 1024 * 1024 * 1024 * 1024),
	// "ZB": uint64(1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000),
	// "Z": uint64(1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024),
	// "YB": uint64(1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000),
	// "Y": uint64(1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024),
}

func init() {

	_, err1 := os.Stat("/usr/bin/df")
	_, err2 := os.Stat("/bin/df")

	if err1 == nil || err2 == nil {
		SUPPORTED = true
	}

}

func toUint64(value string) uint64 {

	var result uint64
	var factor uint64 = 1
	var suffix string

	for str, num := range suffixes {

		if strings.HasSuffix(value, str) {
			factor = num
			suffix = str
			break
		}

	}

	if suffix != "" {
		value = value[0 : len(value)-len(suffix)]
	}

	if strings.Contains(value, ".") {

		tmp, err := strconv.ParseFloat(value, 64)

		if err == nil {
			result = uint64(tmp * float64(factor))
		}

	} else {

		tmp, err := strconv.ParseUint(value, 10, 64)

		if err == nil {
			result = uint64(tmp * factor)
		}

	}

	return result

}
