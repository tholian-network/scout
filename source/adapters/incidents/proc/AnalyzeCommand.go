package proc

import "tholian-core/structs"
import utils_incidents "tholian-scout/internal/utils/incidents"
import utils_programs "tholian-scout/internal/utils/programs"
import "strings"

func AnalyzeCommand(system string, datetime string, program structs.Program) []structs.Incident {

	var result []structs.Incident

	joined := truncate(strings.Join(program.Arguments, " "), 1024)
	evidence := "command " + truncate(strings.TrimSpace(program.Command+" "+joined), 1024)
	network := networkFromProgram(program)

	if strings.Contains(joined, "/dev/tcp/") || strings.Contains(joined, "/dev/udp/") {

		if network != "" {

			incident := toExfilIncident(system, datetime, program, evidence)
			utils_incidents.AddNetwork(&incident, network)
			result = append(result, incident)

		} else {
			result = append(result, toDestroyIncident(system, datetime, program, evidence))
		}

	}

	if utils_programs.IsShell(program) == true {

		if strings.Contains(joined, "curl") || strings.Contains(joined, "wget") {

			if strings.Contains(joined, "|") && (strings.Contains(joined, "sh") || strings.Contains(joined, "ash")) {
				result = append(result, toDestroyIncident(system, datetime, program, evidence))
			}

		}

	}

	if utils_programs.IsBackdoor(program) == true {

		if strings.Contains(joined, "EXEC:") || strings.Contains(joined, "SYSTEM:") {
			result = append(result, toDestroyIncident(system, datetime, program, evidence))
		} else {

			for a := 0; a < len(program.Arguments); a++ {

				arg := strings.TrimSpace(program.Arguments[a])

				if arg == "-e" || arg == "--exec" || arg == "-c" {

					if network != "" {

						incident := toExfilIncident(system, datetime, program, evidence)
						utils_incidents.AddNetwork(&incident, network)
						result = append(result, incident)

					} else {
						result = append(result, toDestroyIncident(system, datetime, program, evidence))
					}

					break

				}

			}

		}

	}

	if utils_programs.IsScanner(program) == true {

		if network != "" {

			incident := toReconIncident(system, datetime, program, evidence)
			utils_incidents.AddNetwork(&incident, network)
			result = append(result, incident)

		} else {
			result = append(result, toDestroyIncident(system, datetime, program, evidence))
		}

	}

	if utils_programs.IsMiner(program) == true {

		if network != "" {

			incident := toExfilIncident(system, datetime, program, evidence)
			utils_incidents.AddNetwork(&incident, network)
			result = append(result, incident)

		} else {
			result = append(result, toDestroyIncident(system, datetime, program, evidence))
		}

	}

	if utils_programs.IsInterpreter(program) == true {

		for a := 1; a < len(program.Arguments); a++ {

			arg := strings.TrimSpace(program.Arguments[a])

			if arg == "" || strings.HasPrefix(arg, "-") || strings.Contains(arg, "=") {
				continue
			}

			if (strings.HasPrefix(arg, "/") || strings.HasPrefix(arg, "./") || strings.HasPrefix(arg, "../")) && IsUntrustedPath(arg) == true {
				result = append(result, toDestroyIncident(system, datetime, program, "script "+truncate(arg, 512)))
				break
			}

		}

	}

	return result

}
