package rpm

import "tholian-core/matchers"
import "strings"

func ParseConflict(value string, result *matchers.Package) bool {

	var verified bool = false

	if strings.Contains(value, "(") && strings.Contains(value, ")") {

		name := strings.TrimSpace(value[0:strings.Index(value, "(")])
		unknown := strings.TrimSpace(value[strings.Index(value, "(")+1 : strings.Index(value, ")")])
		rest := strings.TrimSpace(value[strings.Index(value, ")")+1:])

		if unknown == "x86-32" || unknown == "x86-64" {

			result.Parse(name + " " + rest)
			result.SetArchitecture(unknown)
			verified = true

		} else if unknown == "" {

			result.Parse(name + " " + rest)
			verified = true

		} else if strings.ToLower(unknown) != unknown {

			// "C_HEADER_SYMBOL"
			verified = false

		} else if rest != "" {

			result.Parse(name + "(" + unknown + ") " + rest)
			verified = true

		} else {

			result.SetName(name + "(" + unknown + ")")
			result.SetVersion("any")
			verified = true

		}

	} else {

		result.Parse(value)
		verified = true

	}

	return verified

}
