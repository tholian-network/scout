package rpm

import "tholian-core/matchers"
import "tholian-core/types"
import "strings"

func ParseSharedLibrary(value string, result *matchers.Package) bool {

	var verified bool = false

	if strings.HasPrefix(value, "(") && strings.Contains(value, " if ") && strings.HasSuffix(value, ")") {
		value = strings.TrimSpace(value[1:strings.Index(value, " if ")])
	}

	if strings.Contains(value, "(") && strings.HasSuffix(value, ")") {

		name := strings.TrimSpace(value[0:strings.Index(value, "(")])
		header := strings.TrimSpace(value[strings.Index(value, "(")+1 : strings.Index(value, ")")])
		architecture := strings.TrimSpace(value[strings.Index(value, ")")+1:])

		if strings.HasPrefix(architecture, "(64bit)") {

			if strings.Contains(header, "_") {

				tmp := types.ToVersion(header[strings.LastIndex(header, "_")+1:])

				result.SetName(name[0 : strings.Index(name, ".so")+3])
				result.SetVersion(tmp.String())
				result.SetArchitecture("x86_64")
				verified = true

			} else if strings.Contains(name, ".so") {

				result.SetName(name[0 : strings.Index(name, ".so")+3])
				result.SetVersion("any")
				result.SetArchitecture("x86_64")
				verified = true

			} else {

				result.SetName(name + "(" + header + ")")
				result.SetVersion("any")
				result.SetArchitecture("x86_64")
				verified = true

			}

		} else if strings.HasPrefix(architecture, "(32bit)") {

			if strings.Contains(header, "_") {

				tmp := types.ToVersion(header[strings.LastIndex(header, "_")+1:])

				result.SetName(name[0 : strings.Index(name, ".so")+3])
				result.SetVersion(tmp.String())
				result.SetArchitecture("x86")
				verified = true

			} else if strings.Contains(name, ".so") {

				result.SetName(name[0 : strings.Index(name, ".so")+3])
				result.SetVersion("any")
				result.SetArchitecture("x86")
				verified = true

			} else {

				// TODO: Does this ever happen?
				verified = false

			}

		} else {

			if strings.Contains(header, "_") {

				tmp := types.ToVersion(header[strings.LastIndex(header, "_")+1:])

				result.SetName(name[0 : strings.Index(name, ".so")+3])
				result.SetVersion(tmp.String())
				verified = true

			} else if strings.Contains(name, ".so") {

				result.SetName(name[0 : strings.Index(name, ".so")+3])
				result.SetVersion("any")
				verified = true

			} else {

				result.SetName(name + "(" + header + ")")
				result.SetVersion("any")
				result.SetArchitecture("x86_64")
				verified = true

			}

		}

	} else {

		result.SetName(value[0 : strings.Index(value, ".so")+3])
		result.SetVersion("any")
		verified = true

	}

	return verified

}
