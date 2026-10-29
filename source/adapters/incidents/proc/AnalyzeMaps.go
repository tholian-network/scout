package proc

import "tholian-core/structs"
import utils_os "tholian-scout/internal/utils/os"
import utils_strings "tholian-scout/internal/utils/strings"
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
				result = append(result, toDestroyIncident(system, datetime, program, "mapped "+utils_strings.Truncate(value, 512)))
				seen[value] = true
				collected += 1
			}

		} else if strings.HasPrefix(value, "/") {

			is_deleted := strings.Contains(value, " (deleted)")
			is_trusted := utils_os.IsTrustedPath(value)
			is_exe := strings.Contains(entry.Permissions, "x")
			is_so := strings.Contains(value, ".so")

			if (is_deleted == true || is_trusted == false) && (is_exe == true || is_so == true) {

				if seen[value] == false {
					result = append(result, toDestroyIncident(system, datetime, program, "mapped "+utils_strings.Truncate(value, 512)))
					seen[value] = true
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
