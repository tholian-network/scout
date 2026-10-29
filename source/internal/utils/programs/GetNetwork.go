package programs

import "tholian-core/structs"
import "tholian-core/types"
import "strings"

func GetNetwork(program structs.Program) string {

	for a := 0; a < len(program.Arguments); a++ {

		arg := strings.TrimSpace(program.Arguments[a])

		for _, prefix := range []string{"/dev/tcp/", "/dev/udp/"} {

			if index := strings.Index(arg, prefix); index != -1 {
				arg = arg[index+len(prefix):]
				break
			}

		}

		fields := strings.FieldsFunc(arg, func(character rune) bool {

			if character == ' ' ||
				character == '\t' ||
				character == '\'' ||
				character == '"' ||
				character == '`' ||
				character == ';' ||
				character == ',' {
				return true
			}

			return false

		})

		for f := 0; f < len(fields); f++ {

			host := strings.Trim(fields[f], "[](){}")

			host = strings.TrimSuffix(host, ".")

			if strings.Contains(host, "/") {
				host = host[0:strings.Index(host, "/")]
			}

			if types.IsIPv4(host) == true || types.IsIPv6(host) == true {
				return host
			}

		}

	}

	return ""

}
