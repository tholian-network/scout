package dnf

import "tholian-scout/adapters/packages/rpm"
import "tholian-core/matchers"
import "tholian-core/structs"
import "strings"

func ParsePackage(buffer string, result *structs.Package) {

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	var key string
	var val string

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if strings.Contains(line, " : ") {

			key = strings.TrimSpace(line[0:strings.Index(line, " : ")])
			val = strings.TrimSpace(line[strings.Index(line, " : ")+3:])

		} else if strings.HasSuffix(line, " :") {

			key = strings.TrimSpace(line[0:strings.Index(line, " :")])
			val = ""

		} else {

			val = strings.TrimSpace(line)

		}

		if key == "Architecture" {

			result.SetArchitecture(val)

		} else if key == "Buildtime" {

			result.SetDatetime(val)

		} else if key == "Conflicts" {

			if val != "" {

				conflict := matchers.NewPackage()
				conflict.SetManager("dnf")

				parse_result := rpm.ParseConflict(val, &conflict)

				if parse_result == true {
					result.AddConflict(conflict)
				}

			}

		} else if key == "Files" {

			if val != "" {

				if strings.HasPrefix(val, "/") {
					result.AddFilesystem(val)
				}

			}

		} else if key == "Name" {

			result.SetName(val)

		} else if key == "Obsoletes" {

			if val != "" {

				conflict := matchers.ToPackage(val)
				conflict.SetManager("dnf")
				result.AddConflict(conflict)

			}

		} else if key == "Provides" {

			if val != "" {

				provide := matchers.NewPackage()
				provide.SetManager("dnf")

				parse_result := rpm.ParseProvide(val, &provide)

				if parse_result == true {
					result.AddProvide(provide)
				}

			}

		} else if key == "Requires" {

			if val != "" {

				dependency := matchers.NewPackage()
				dependency.SetManager("dnf")

				parse_result := rpm.ParseRequire(val, &dependency)

				if parse_result == true {
					result.AddDependency(dependency)
				}

			}

		} else if key == "URL" {

			if val != "" {
				result.SetURL(val)
			}

		} else if key == "Version" {

			result.SetVersion(val)

		}

	}

}
