package programs

import "tholian-core/structs"
import "path"

func toProgramNames(program structs.Program) []string {

	values := make([]string, 0)

	if program.Name != "" {
		values = append(values, path.Base(program.Name))
	}

	if program.Command != "" {
		values = append(values, path.Base(program.Command))
	}

	if len(program.Arguments) > 0 && program.Arguments[0] != "" {
		values = append(values, path.Base(program.Arguments[0]))
	}

	for v := 0; v < len(values); v++ {

		if values[v] == "busybox" && len(program.Arguments) > 1 {

			values = append(values, path.Base(program.Arguments[1]))
			break

		}

	}

	return values

}
