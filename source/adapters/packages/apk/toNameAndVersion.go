package apk

import "strings"

func toNameAndVersion(value string) (string, string) {

	var name string
	var version string

	tmp := strings.Split(value, "-")

	var mode = "name"

	for t := 0; t < len(tmp); t++ {

		if mode == "name" {

			chr := string(tmp[t][0])

			if t > 0 && chr >= "0" && chr <= "9" {

				mode = "version"
				version = tmp[t]

			} else {

				if name != "" {
					name = name + "-" + tmp[t]
				} else {
					name = tmp[t]
				}

			}

		} else if mode == "version" {

			if version != "" {
				version = version + "-" + tmp[t]
			} else {
				version = tmp[t]
			}

		}

	}

	return name, version

}
