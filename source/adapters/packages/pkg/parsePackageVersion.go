package pkg

import "tholian-core/types"
import "strconv"
import "strings"

func parsePackageVersion(value string) *types.Version {

	var result *types.Version = nil

	value = strings.TrimSpace(value)

	var epoche uint = 0
	var revision string = ""

	if index := strings.LastIndex(value, ","); index != -1 {

		if number, err := strconv.ParseUint(strings.TrimSpace(value[index+1:]), 10, 64); err == nil {
			epoche = uint(number)
			value = strings.TrimSpace(value[0:index])
		}

	}

	if index := strings.LastIndex(value, "_"); index != -1 {

		if _, err := strconv.ParseUint(strings.TrimSpace(value[index+1:]), 10, 64); err == nil {
			revision = strings.TrimSpace(value[index+1:])
			value = strings.TrimSpace(value[0:index])
		}

	}

	tmp := types.ToVersion(value)

	if tmp.IsValid() {

		result = &tmp

		if epoche > 0 {
			result.Epoche = epoche
		}

		if revision != "" {
			result.Revision = revision
		}

	}

	return result

}
