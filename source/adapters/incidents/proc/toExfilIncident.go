package proc

import "tholian-core/matchers"
import "tholian-core/structs"

func toExfilIncident(system string, datetime string, program structs.Program, evidence string) structs.Incident {

	var incident = structs.NewIncident(system, "Exfil")

	incident.SetDatetime(datetime)

	matcher := matchers.NewProgram()

	if program.Name != "" {
		matcher.SetName(program.Name)
	}

	if program.Command != "" {
		matcher.SetCommand(program.Command)
	}

	if len(program.Arguments) > 0 {
		matcher.SetArguments(capArguments(program.Arguments))
	}

	incident.AddProgram(matcher)

	if program.User.IsValid() {
		incident.AddUser(program.User)
	}

	if evidence != "" {

		value := structs.NewEvidence(structs.EvidenceTypeText, []byte(truncate(evidence, 1024)))
		value.SetDatetime(datetime)
		incident.AddEvidence(value)

	}

	return incident

}

