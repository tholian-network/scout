package proc

import "tholian-core/structs"

func AnalyzeProgram(system string, datetime string, program structs.Program) []structs.Incident {

	var result []structs.Incident

	result = append(result, AnalyzeEnvironment(system, datetime, program)...)
	result = append(result, AnalyzeMaps(system, datetime, program)...)
	result = append(result, AnalyzeExecutable(system, datetime, program)...)
	result = append(result, AnalyzeCommand(system, datetime, program)...)
	result = append(result, AnalyzeConnections(system, datetime, program)...)

	if len(result) > 32 {
		result = result[0:32]
	}

	return result

}
