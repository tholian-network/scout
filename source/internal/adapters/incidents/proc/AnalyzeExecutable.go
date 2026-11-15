package proc

import "tholian-core/structs"
import utils_os "tholian-scout/internal/utils/os"
import utils_strings "tholian-scout/internal/utils/strings"
import "strings"

func AnalyzeExecutable(system string, datetime string, program structs.Program) []structs.Incident {

	var result []structs.Incident

	command := strings.TrimSpace(program.Command)

	if command == "" {
		return result
	}

	if strings.Contains(command, " (deleted)") {
		result = append(result, toDestroyIncident(system, datetime, program, "executable "+utils_strings.Truncate(command, 512)))
	} else if utils_os.IsTrustedPath(command) == false {
		result = append(result, toDestroyIncident(system, datetime, program, "executable "+utils_strings.Truncate(command, 512)))
	}

	return result

}
