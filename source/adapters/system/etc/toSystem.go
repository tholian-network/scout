package etc

import "tholian-insights/distributions"
import "tholian-core/structs"
import "encoding/hex"
import "math/rand"
import "os"
import "strings"

func toSystem() structs.System {

	var system structs.System = structs.NewSystem()

	buffer1, err1 := os.ReadFile("/etc/os-release")

	if err1 == nil {

		keywords := toKeywords(buffer1)
		distributions := distributions.Distributions.QueryByKeywords(keywords)

		if len(distributions) > 0 {

			candidate := distributions[0]

			distribution := structs.NewDistribution()
			distribution.SetName(candidate.Name)
			distribution.SetVersion(candidate.Version)
			distribution.SetManager(candidate.Manager)
			distribution.SetVendor(candidate.Vendor)

			system.SetDistribution(distribution)

		}

	}

	buffer2, err2 := os.ReadFile("/etc/machine-id")

	if err2 == nil {

		system.SetName(strings.TrimSpace(string(buffer2)))

	} else {

		seed := make([]byte, 16)
		rand.Read(seed)
		uuid := hex.EncodeToString(seed)

		err := os.WriteFile("/etc/machine-id", []byte(uuid), 0444)

		if err == nil {
			system.SetName(uuid)
		} else if hostname, hostname_err := os.Hostname(); hostname_err == nil && hostname != "" {
			system.SetName(hostname)
		}

	}

	timezone, err3 := os.Readlink("/etc/localtime")

	if err3 == nil {

		if strings.HasPrefix(timezone, "/usr/share/zoneinfo/") {
			timezone = timezone[20:]
		}

		system.SetTimezone(timezone)

	}

	locale_conf, err4 := os.ReadFile("/etc/locale.conf")

	if err4 == nil {

		var lines = strings.Split(strings.TrimSpace(string(locale_conf)), "\n")

		for l := 0; l < len(lines); l++ {

			var line = strings.TrimSpace(lines[l])

			if strings.HasPrefix(line, "LANG=") {

				locale := line[5:]

				if strings.HasSuffix(locale, ".UTF-8") {
					locale = locale[0 : len(locale)-6]
				}

				system.SetLocale(locale)

			}

		}

	}

	return system

}
