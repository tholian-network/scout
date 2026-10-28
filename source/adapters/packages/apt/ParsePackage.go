package apt

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "strings"
import "time"

func ParsePackage(buffer string, result *structs.Package) {

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	var key string
	var val string

	for l := 0; l < len(lines); l++ {

		line := lines[l]

		if strings.HasPrefix(line, " ") {

			val = val + "\n " + strings.TrimSpace(line)

		} else if strings.Contains(line, ": ") {

			tmp := strings.Split(line, ": ")

			key = strings.TrimSpace(tmp[0])
			val = strings.TrimSpace(tmp[1])

		}

		if key == "Architecture" {

			result.SetArchitecture(val)

		} else if key == "Depends" {

			chunks := strings.Split(val, ", ")

			for c := 0; c < len(chunks); c++ {

				chunk := strings.TrimSpace(chunks[c])

				if strings.Contains(chunk, " | ") {

					unresolved := matchers.ToUnresolved(chunk)
					result.Unresolved = append(result.Unresolved, unresolved)

				} else if chunk != "" {

					dependency := matchers.ToPackage(chunk)
					dependency.SetManager("apt")
					result.AddDependency(dependency)

				}

			}

		} else if key == "Homepage" {

			result.SetURL(val)

		} else if key == "Maintainer" || key == "Original-Maintainer" {

			result.AddMaintainer(types.ToMaintainer(val))

		} else if key == "Package" {

			result.SetName(val)

		} else if key == "Provides" {

			chunks := strings.Split(val, ", ")

			for c := 0; c < len(chunks); c++ {

				chunk := strings.TrimSpace(chunks[c])

				if chunk != "" {

					provide := matchers.ToPackage(chunk)
					provide.SetManager("apt")
					result.AddProvide(provide)

				}

			}

		} else if key == "Replaces" {

			chunks := strings.Split(val, ", ")

			for c := 0; c < len(chunks); c++ {

				chunk := strings.TrimSpace(chunks[c])

				if chunk != "" {

					replace := matchers.ToPackage(chunk)
					replace.SetManager("apt")
					result.AddReplace(replace)

				}

			}

		} else if key == "Version" {

			result.SetVersion(val)

		}

	}

	if result.Datetime.IsValid() == false {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

}
