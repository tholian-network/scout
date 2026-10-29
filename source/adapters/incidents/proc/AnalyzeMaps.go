package proc

import "tholian-core/structs"
import "os"
import "strconv"
import "strings"

func AnalyzeMaps(system string, datetime string, program structs.Program) []structs.Incident {

	var result []structs.Incident

	pid := strconv.FormatUint(uint64(program.PID), 10)
	buffer, err := os.ReadFile("/proc/" + pid + "/maps")

	if err != nil {
		return result
	}

	entries := parseProcessMaps(buffer)
	seen := make(map[string]bool)
	collected := 0

	for e := 0; e < len(entries); e++ {

		entry := entries[e]
		value := entry.Path

		if value == "" {
			continue
		}

		if strings.HasPrefix(value, "memfd:") {

			if seen[value] == false {
				seen[value] = true
				result = append(result, toDestroyIncident(system, datetime, program, "mapped "+truncate(value, 512)))
				collected += 1
			}

		} else if strings.HasPrefix(value, "/") {

			deleted := strings.Contains(value, " (deleted)")
			untrusted := IsUntrustedPath(value)
			executable := strings.Contains(entry.Permissions, "x")
			shared := strings.Contains(value, ".so")

			if (deleted == true || untrusted == true) && (executable == true || shared == true) {

				if seen[value] == false {
					seen[value] = true
					result = append(result, toDestroyIncident(system, datetime, program, "mapped "+truncate(value, 512)))
					collected += 1
				}

			}

		}

		if collected >= 8 {
			break
		}

	}

	return result

}
