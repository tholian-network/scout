package pkg

import "tholian-core/types"
import "strings"

func parseInstalledPackages(buffer string) map[string]types.Version {

	collected := make(map[string]types.Version)
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		fields := strings.Split(strings.TrimSpace(lines[l]), "\t")

		if len(fields) >= 2 && strings.TrimSpace(fields[0]) != "" {

			tmp := parsePackageVersion(fields[1])

			if tmp != nil {
				collected[strings.TrimSpace(fields[0])] = *tmp
			}

		}

	}

	return collected

}
