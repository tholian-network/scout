package proc

import "strings"

type processMap struct {
	Start       string
	End         string
	Permissions string
	Offset      string
	Device      string
	Inode       string
	Path        string
}

func parseProcessMaps(buffer []byte) []processMap {

	var result []processMap

	lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if line != "" {

			entry := processMap{}
			fields := make([]string, 0, 5)
			remainder := line

			for f := 0; f < 5; f++ {

				index := strings.IndexAny(remainder, " \t")

				if index != -1 {
					fields = append(fields, remainder[0:index])
					remainder = strings.TrimLeft(remainder[index:], " \t")
				} else {
					break
				}

			}

			address := fields[0]

			if index := strings.Index(address, "-"); index != -1 {
				entry.Start = address[0:index]
				entry.End = address[index+1:]
			} else {
				entry.Start = address
			}

			entry.Permissions = fields[1]
			entry.Offset = fields[2]
			entry.Device = fields[3]
			entry.Inode = fields[4]
			entry.Path = strings.TrimSpace(remainder)

			result = append(result, entry)

		}

	}

	return result

}


