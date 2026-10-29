package proc

import "tholian-core/structs"
import utils_os "tholian-scout/internal/utils/os"
import "strings"

func AnalyzeEnvironment(system string, datetime string, program structs.Program) []structs.Incident {

	var result []structs.Incident

	keys := []string{
		"LD_AUDIT",
		"LD_LIBRARY_PATH",
		"LD_PRELOAD",
	}

	for k := 0; k < len(keys); k++ {

		key := keys[k]
		value, ok := program.Environment[key]

		if ok == false || strings.TrimSpace(value) == "" {
			continue
		}

		entries := strings.FieldsFunc(value, func(character rune) bool {
			return character == ':' || character == ' ' || character == '\t'
		})

		suspicious := make([]string, 0)

		for e := 0; e < len(entries); e++ {

			entry := strings.TrimSpace(entries[e])

			if entry != "" && utils_os.IsTrustedPath(entry) == false {
				suspicious = append(suspicious, entry)
			}

		}

		if len(suspicious) > 0 {
			evidence := key + "=" + strings.Join(suspicious, ":")
			result = append(result, toDestroyIncident(system, datetime, program, evidence))
		}

	}

	return result

}
