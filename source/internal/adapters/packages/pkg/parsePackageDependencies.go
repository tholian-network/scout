package pkg

import "tholian-core/matchers"
import "strings"

func parsePackageDependencies(buffer string, parent string) []matchers.Package {

	var collected []matchers.Package

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		fields := strings.Split(strings.TrimSpace(lines[l]), "\t")
		name := ""
		version := ""

		if len(fields) >= 3 {

			if parent != "" && strings.TrimSpace(fields[0]) != parent {
				continue
			}

			name = strings.TrimSpace(fields[1])
			version = strings.TrimSpace(fields[2])

		} else if len(fields) == 2 && parent != "" {
			name = strings.TrimSpace(fields[0])
			version = strings.TrimSpace(fields[1])
		} else if len(fields) == 1 && parent != "" {
			name = strings.TrimSpace(fields[0])
		}

		if name != "" {

			dependency := matchers.NewPackage()
			dependency.SetName(name)
			dependency.SetVersion(version)
			dependency.SetArchitecture("any")
			dependency.SetManager("pkg")

			collected = append(collected, dependency)

		}

	}

	return collected

}
