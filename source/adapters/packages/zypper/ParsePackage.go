package zypper

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-scout/adapters/packages/rpm"
import "strings"
import "time"

func ParsePackage(buffer string, result *structs.Package) {

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	var key string
	var val string

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if strings.Contains(line, " : ") {

			key = strings.TrimSpace(line[0:strings.Index(line, " : ")])
			val = strings.TrimSpace(line[strings.Index(line, " : ")+3:])

			if key == "Provides" || key == "Requires" || key == "Conflicts" || key == "Obsoletes" {

				// Ignore counters for multi-line output
				if strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]") {
					val = ""
				}

			}

		} else {

			val = strings.TrimSpace(line)

		}

		if key == "Arch" {

			if val != "" {
				result.SetArchitecture(val)
			}

		} else if key == "Conflicts" {

			if val != "" && val != "---" {

				conflict := matchers.NewPackage()
				conflict.SetManager("zypper")

				parse_result := rpm.ParseConflict(val, &conflict)

				if parse_result == true {
					result.AddConflict(conflict)
				}

			}

		} else if key == "Name" {

			if val != "" {
				result.SetName(val)
			}

		} else if key == "Vendor" {

			if val != "" {
				result.SetVendor(strings.ToLower(val))
			}

		} else if key == "Obsoletes" {

			if val != "" && val != "---" {

				conflict := matchers.ToPackage(val)
				conflict.SetManager("zypper")
				result.AddConflict(conflict)

			}

		} else if key == "Provides" {

			if val != "" && val != "---" {

				provide := matchers.NewPackage()
				provide.SetManager("zypper")

				parse_result := rpm.ParseProvide(val, &provide)

				if parse_result == true {
					result.AddProvide(provide)
				}

			}

		} else if key == "Requires" {

			if val != "" && val != "---" {

				dependency := matchers.NewPackage()
				dependency.SetManager("zypper")

				parse_result := rpm.ParseRequire(val, &dependency)

				if parse_result == true {
					result.AddDependency(dependency)
				}

			}

		} else if key == "Status" {

			if val != "" {

				if val == "up-to-date" {

					// Do Nothing, already uses correct package version

				} else if strings.HasPrefix(val, "out-of-date (version ") && strings.HasSuffix(val, " installed)") {

					// Set the correct _local_ package version
					version := strings.TrimSpace(val[21 : len(val)-11])

					if version != "" {
						result.SetVersion(version)
					}

				}

			}

		} else if key == "Upstream URL" {

			if val != "" {
				result.SetURL(val)
			}

		} else if key == "Version" {

			if val != "" {
				result.SetVersion(val)
			}

		}

	}

	if result.Datetime.IsValid() == false {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

}
