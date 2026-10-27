package sys

import "os"
import "strings"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/sys")
	_, err2 := os.Stat("/sys/class/dmi")

	if err1 == nil && err2 == nil {
		SUPPORTED = true
	}

}

func fixVersion(version string, release string) string {

	if strings.HasSuffix(version, " )") {
		version = strings.TrimSpace(version[0:len(version)-2] + ")")
	}

	var suffix = "(" + release + ")"

	if strings.HasSuffix(version, suffix) {
		version = strings.TrimSpace(version[0 : len(version)-len(suffix)])
	}

	return version

}

func readDMI(file string) string {

	var result string

	buffer, err := os.ReadFile("/sys/class/dmi/id/" + file)

	if err == nil {

		var str = strings.TrimSpace(string(buffer))

		if str != "Not Defined" && str != "To be filled by O.E.M." && str != "" {
			result = str
		}

	}

	return result

}
