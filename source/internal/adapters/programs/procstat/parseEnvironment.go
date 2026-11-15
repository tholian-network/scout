package procstat

import "strings"

func parseEnvironment(value string) map[string]string {

	collected := make(map[string]string)

	fields := strings.Fields(value)

	for f := 0; f < len(fields); f++ {

		index := strings.Index(fields[f], "=")

		if index > 0 {
			collected[fields[f][0:index]] = fields[f][index+1:]
		}

	}

	return collected

}
