package proc

import "tholian-core/structs"
import utils_incidents "tholian-scout/internal/utils/incidents"
import utils_programs "tholian-scout/internal/utils/programs"
import "strconv"

func AnalyzeConnections(system string, datetime string, program structs.Program) []structs.Incident {

	var result []structs.Incident

	is_backdoor := utils_programs.IsBackdoor(program)
	is_shell := utils_programs.IsShell(program)

	if is_backdoor == false && is_shell == false {
		return result
	}

	for c := 0; c < len(program.Connections); c++ {

		connection := program.Connections[c]
		port := strconv.FormatUint(uint64(connection.Port), 10)

		if connection.HasHost() == false {

			if connection.Type == "server" {
				result = append(result, toDestroyIncident(system, datetime, program, "listener any:"+port))
			}

			continue

		}

		if connection.Type == "client" {

			incident := toExfilIncident(system, datetime, program, "connection "+connection.Host+":"+port)
			utils_incidents.AddNetwork(&incident, connection.Host)
			result = append(result, incident)

		} else if connection.Type == "server" {

			if is_backdoor == true {

				incident := toReconIncident(system, datetime, program, "listener "+connection.Host+":"+port)
				utils_incidents.AddNetwork(&incident, connection.Host)
				result = append(result, incident)

			} else if is_shell == true {

				incident := toExfilIncident(system, datetime, program, "listener "+connection.Host+":"+port)
				utils_incidents.AddNetwork(&incident, connection.Host)
				result = append(result, incident)

			}

		}

	}

	return result

}
