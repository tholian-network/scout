package procstat

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-scout/adapters/incidents/proc"
import programs "tholian-scout/adapters/programs/procstat"
import utils_os "tholian-scout/internal/utils/os"
import utils_strings "tholian-scout/internal/utils/strings"
import "strings"

func AnalyzeProgram(system string, datetime string, program structs.Program, mappings []programs.Mapping) []structs.Incident {

	result := proc.AnalyzeProgram(system, datetime, program)

	seen := make(map[string]bool)

	for m := 0; m < len(mappings); m++ {

		path := strings.TrimSpace(mappings[m].Path)

		if strings.HasPrefix(path, "/") {

			is_executable := strings.Contains(mappings[m].Permissions, "x")
			is_shared_library := strings.Contains(path, ".so")

			if (is_executable || is_shared_library) && utils_os.IsTrustedPath(path) == false {

				if seen[path] == false {

					incident := structs.NewIncident(system, "Destroy")
					incident.SetDatetime(datetime)

					matcher := matchers.NewProgram()

					if program.Name != "" {
						matcher.SetName(program.Name)
					}

					if program.Command != "" {
						matcher.SetCommand(program.Command)
					}

					if len(program.Arguments) > 0 {
						matcher.SetArguments(program.Arguments)
					}

					incident.AddProgram(matcher)

					if program.User.IsValid() {
						incident.AddUser(program.User)
					}

					evidence := structs.NewEvidence(structs.EvidenceTypeText, []byte("mapped "+utils_strings.Truncate(path, 512)))
					evidence.SetDatetime(datetime)
					incident.AddEvidence(evidence)

					result = append(result, incident)
					seen[path] = true

				}

			}

		}

	}

	if len(result) > 32 {
		result = result[0:32]
	}

	return result

}
