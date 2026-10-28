package tdnf

import "tholian-scout/adapters/packages/rpm"
import "tholian-core/matchers"
import "tholian-core/structs"
import "strings"
import "time"

func ParsePackage(buffer string, result *structs.Package) {

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	var key string
	var val string

	var epoch string
	var version string
	var release string

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

		if key == "Arch" || key == "Architecture" {

			result.SetArchitecture(val)

		} else if key == "Buildtime" {

			result.SetDatetime(val)

		} else if key == "Conflicts" {

			if val != "" {

				conflict := matchers.NewPackage()
				conflict.SetManager("tdnf")

				parse_result := rpm.ParseConflict(val, &conflict)

				if parse_result == true {
					result.AddConflict(conflict)
				}

			}

		} else if key == "Epoch" {

			epoch = val

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
				conflict.SetManager("tdnf")
				result.AddConflict(conflict)

			}

		} else if key == "Provides" {

			if val != "" {

				provide := matchers.NewPackage()
				provide.SetManager("tdnf")

				parse_result := rpm.ParseProvide(val, &provide)

				if parse_result == true {
					result.AddProvide(provide)
				}

			}

		} else if key == "Release" {

			release = val

		} else if key == "Requires" {

			if val != "" {

				dependency := matchers.NewPackage()
				dependency.SetManager("tdnf")

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

			version = val

		}

	}

	if version != "" {

		if release != "" {

			evr := version + "-" + release

			if epoch != "" && epoch != "0" {
				evr = epoch + ":" + evr
			}

			result.SetVersion(evr)

		} else {

			result.SetVersion(version)

		}

	}

	if result.Datetime.IsValid() == false {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

}
