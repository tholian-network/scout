package opkg

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "strconv"
import "strings"
import "time"

func ParsePackage(buffer string, result *structs.Package) {

	fields := parseControlFields(buffer)

	result.SetManager("opkg")

	if value, ok := fields["Package"]; ok {
		result.SetName(value)
	}

	if value, ok := fields["Version"]; ok {
		result.SetVersion(value)
	}

	if value, ok := fields["Architecture"]; ok {
		result.SetArchitecture(value)
	}

	if value, ok := fields["Homepage"]; ok {
		result.SetURL(value)
	}

	if value, ok := fields["URL"]; ok {
		result.SetURL(value)
	}

	if value, ok := fields["Maintainer"]; ok {
		result.AddMaintainer(types.ToMaintainer(value))
	}

	if value, ok := fields["Depends"]; ok {

		relations := splitRelations(value)

		for r := 0; r < len(relations); r++ {

			relation := relations[r]

			if strings.Contains(relation, " | ") {

				result.AddUnresolved(matchers.ToUnresolved(relation))

			} else {

				dependency := matchers.ToPackage(relation)
				dependency.SetManager("opkg")
				result.AddDependency(dependency)

			}

		}

	}

	if value, ok := fields["Provides"]; ok {

		relations := splitRelations(value)

		for r := 0; r < len(relations); r++ {

			provide := matchers.ToPackage(relations[r])
			provide.SetManager("opkg")
			result.AddProvide(provide)

		}

	}

	if value, ok := fields["Replaces"]; ok {

		relations := splitRelations(value)

		for r := 0; r < len(relations); r++ {

			replace := matchers.ToPackage(relations[r])
			replace.SetManager("opkg")
			result.AddReplace(replace)

		}

	}

	if value, ok := fields["Conflicts"]; ok {

		relations := splitRelations(value)

		for r := 0; r < len(relations); r++ {

			conflict := matchers.ToPackage(relations[r])
			conflict.SetManager("opkg")
			result.AddConflict(conflict)

		}

	}

	if value, ok := fields["Installed-Time"]; ok {

		number, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64)

		if err == nil {
			result.SetDatetime(time.Unix(number, 0).Format(time.RFC3339))
		}

	}

	if result.Datetime.IsValid() == false {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

}
