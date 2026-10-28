package dpkg

import "tholian-core/types"
import "strings"

func stripArchitecture(value string) string {

	index := strings.LastIndex(value, ":")

	if index > 0 {

		suffix := value[index+1:]

		if suffix == "any" || types.ParseArchitecture(suffix) != nil {
			return value[0:index]
		}

	}

	return value

}
