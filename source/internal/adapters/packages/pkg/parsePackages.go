package pkg

import "tholian-core/structs"
import "strings"

func parsePackages(buffer string) []structs.Package {

	var collected []structs.Package

	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if line == "" {
			continue
		}

		pkg := structs.NewPackage("pkg")

		ParsePackage(line, &pkg)

		if pkg.Name != "" && pkg.Version.IsValid() {
			collected = append(collected, pkg)
		}

	}

	return collected

}
