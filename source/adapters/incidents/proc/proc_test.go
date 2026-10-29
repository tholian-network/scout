package proc

import "tholian-core/structs"

func programWithEnvironment(key string, value string) structs.Program {

	program := structs.NewProgram(1234, "test")
	program.SetCommand("/usr/bin/test")
	program.SetArguments([]string{"test"})
	program.SetEnvironment(key, value)

	return program

}

func programWithArguments(name string, command string, arguments ...string) structs.Program {

	program := structs.NewProgram(1234, name)
	program.SetCommand(command)
	program.SetArguments(arguments)

	return program

}

