package proc

import "tholian-core/matchers"
import "tholian-core/structs"
import utils_strings "tholian-scout/internal/utils/strings"

func toDestroyIncident(system string, datetime string, program structs.Program, evidence string) structs.Incident {

	var incident = structs.NewIncident(system, "Destroy")

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

	if evidence != "" {

		value := structs.NewEvidence(structs.EvidenceTypeText, []byte(utils_strings.Truncate(evidence, 1024)))
		value.SetDatetime(datetime)
		incident.AddEvidence(value)

	}

	return incident

}
