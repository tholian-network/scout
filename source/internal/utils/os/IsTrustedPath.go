package os

import "strings"

var trusted_prefixes = []string{
	"/bin",
	"/lib",
	"/lib32",
	"/lib64",
	"/sbin",
	"/usr/bin",
	"/usr/lib",
	"/usr/lib32",
	"/usr/lib64",
	"/usr/libexec",
	"/usr/local/bin",
	"/usr/local/lib",
	"/usr/local/lib32",
	"/usr/local/lib64",
	"/usr/local/sbin",
	"/usr/sbin",
	"/usr/share",
}

func IsTrustedPath(value string) bool {

	if value != "" {

		if strings.HasPrefix(value, "memfd:") {

			return false

		} else {

			normalized := NormalizePath(value)

			if normalized != "" && normalized != "." {

				for t := 0; t < len(trusted_prefixes); t++ {

					prefix := trusted_prefixes[t]

					if normalized == prefix || strings.HasPrefix(normalized, prefix+"/") {
						return true
					}

				}

			}

		}

	}

	return false

}
