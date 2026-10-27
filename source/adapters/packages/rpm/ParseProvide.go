package rpm

import "tholian-core/matchers"
import "strings"

func ParseProvide(value string, result *matchers.Package) bool {

	var verified bool = false

	if IsPackageFunction(value) {

		// Red Hat is weird
		verified = false

	} else if strings.Contains(value, ".so.") || strings.Contains(value, ".so(") {

		shared_library := matchers.NewPackage()
		ParseSharedLibrary(value, &shared_library)

		if shared_library.Name != "" {
			result.SetName(shared_library.Name)
			verified = true
		}

		if shared_library.Version != "" {
			result.SetVersion(shared_library.Version)
		}

		if shared_library.Architecture != "" {
			result.SetArchitecture(shared_library.Architecture)
		}

	} else if strings.Contains(value, "(") && strings.Contains(value, ")") {

		name := strings.TrimSpace(value[0:strings.Index(value, "(")])
		unknown := strings.TrimSpace(value[strings.Index(value, "(")+1 : strings.Index(value, ")")])
		rest := strings.TrimSpace(value[strings.Index(value, ")")+1:])

		if unknown == "x86-32" || unknown == "x86_32" || unknown == "x86-64" || unknown == "x86_64" {

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

	} else if strings.HasPrefix(value, "/") {

		result.SetName(value)
		verified = true

	} else {

		result.Parse(value)
		verified = true

	}

	return verified

}
