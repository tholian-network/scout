package pkg

import "tholian-core/structs"
import "tholian-core/types"
import "strconv"
import "strings"
import "time"

func ParsePackage(buffer string, result *structs.Package) {

	fields := strings.Split(strings.TrimRight(buffer, "\n"), "\t")

	if len(fields) < 9 {
		return
	}

	result.SetManager("pkg")
	result.SetName(fields[0])

	version := parsePackageVersion(fields[1])

	if version != nil {
		result.SetVersion(version.String())
	}

	result.SetArchitecture(fields[2])
	result.SetVendor(fields[3])
	result.SetURL(fields[4])

	if strings.TrimSpace(fields[5]) != "" {

		maintainer := types.ToMaintainer(fields[5])

		if maintainer.Name == "" && strings.Contains(maintainer.Email, "@") {
			maintainer.Name = strings.TrimSpace(maintainer.Email[0:strings.Index(maintainer.Email, "@")])
		}

		if maintainer.IsValid() == true {
			result.AddMaintainer(maintainer)
		}

	}

	if number, err := strconv.ParseInt(strings.TrimSpace(fields[7]), 10, 64); err == nil {
		result.SetDatetime(time.Unix(number, 0).Format("2006-01-02 15:04:05"))
	}

	if result.Datetime.IsValid() == false {
		result.SetDatetime(time.Now().Format("2006-01-02 15:04:05"))
	}

}
