package pkg

import "tholian-core/types"
import "strings"

type remote_package struct {
	Version      types.Version
	Architecture string
}

func parseRemotePackages(buffer string) map[string]remote_package {

	collected := make(map[string]remote_package)
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		fields := strings.Split(strings.TrimSpace(lines[l]), "\t")

		if len(fields) >= 2 && strings.TrimSpace(fields[0]) != "" {

			architecture := "any"

			if len(fields) >= 3 && strings.TrimSpace(fields[2]) != "" {
				architecture = strings.TrimSpace(fields[2])
			}

			collected[strings.TrimSpace(fields[0])] = remote_package{
				Version:      *parsePackageVersion(fields[1]),
				Architecture: architecture,
			}

		}

	}

	return collected

}
