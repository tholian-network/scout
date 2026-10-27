package rpm

import "strings"

func IsPackageFunction(value string) bool {

	var result bool = false

	functions := []string{
		"metainfo",
		"alternative",
		"application",
		"base-module",
		"bundled",
		"config",
		"dnf-command",
		"fedora-repos",
		"flavor",
		"font",
		"firmware",
		"group",
		"installonlypkg",
		"kmod",
		"ksym",
		"locale",
		"mimehandler",
		"multiversion",
		"pkgconfig",
		"postscriptdriver",
		"product-url",
		"rpm_macro",
		"system-release",
		"typelib",
		"user",
		"variant_config",
		"weakremover",
		"zypper",
	}

	for f := 0; f < len(functions); f++ {

		if strings.HasPrefix(value, functions[f]+"(") && strings.Contains(value, ")") {
			result = true
			break
		}

	}

	return result

}
