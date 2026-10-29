package procstat

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-scout/adapters/users/bsdshadow"
import "os"
import "os/exec"
import "path"
import "sort"
import "strings"

func CollectProcesses() []structs.Program {

	var collected []structs.Program

	if SUPPORTED == true {

		cmd1 := exec.Command("ps", "-ax", "-o", "pid,uid,comm")
		cmd1.Env = append(os.Environ(), "LC_ALL=C")
		ps_buffer, err1 := cmd1.Output()

		cmd2 := exec.Command("procstat", "-a", "-h", "-b")
		cmd2.Env = append(os.Environ(), "LC_ALL=C")
		binary_buffer, err2 := cmd2.Output()

		cmd3 := exec.Command("procstat", "-a", "-h", "-c")
		cmd3.Env = append(os.Environ(), "LC_ALL=C")
		arguments_buffer, err3 := cmd3.Output()

		cmd4 := exec.Command("procstat", "-a", "-h", "-e")
		cmd4.Env = append(os.Environ(), "LC_ALL=C")
		environment_buffer, err4 := cmd4.Output()

		cmd5 := exec.Command("procstat", "-a", "-h", "-f")
		cmd5.Env = append(os.Environ(), "LC_ALL=C")
		files_buffer, err5 := cmd5.Output()

		cmd6 := exec.Command("sockstat", "-n", "-4", "-6")
		cmd6.Env = append(os.Environ(), "LC_ALL=C")
		sockets_buffer, err6 := cmd6.Output()

		processes := make([]process_info, 0)
		binary := make(map[uint]string)
		arguments := make(map[uint]string)
		environment := make(map[uint]string)
		files := make(map[uint]file_info)
		sockets := make([]socket_info, 0)

		if err1 == nil {
			processes = parseProcessStatus(string(ps_buffer))
		}

		if err2 == nil {
			binary = parseProcstatBinary(string(binary_buffer))
		}

		if err3 == nil {
			arguments = parseProcstatValues(string(arguments_buffer))
		}

		if err4 == nil {
			environment = parseProcstatValues(string(environment_buffer))
		}

		if err5 == nil {
			files = parseProcstatFiles(string(files_buffer))
		}

		if err6 == nil {
			sockets = parseSockstat(string(sockets_buffer))
		}

		for p := 0; p < len(processes); p++ {

			process := processes[p]
			command := binary[process.PID]

			if strings.HasPrefix(command, "/") {

				name := process.Name

				if strings.Contains(command, "/") {
					name = path.Base(command)
				}

				if name == "" {
					name = command
				}

				if name != "" {

					program := structs.NewProgram(process.PID, name)
					program.SetCommand(command)
					program.SetArguments(strings.Fields(arguments[process.PID]))

					info := files[process.PID]

					if info.Folder != "" {
						program.SetFolder(info.Folder)
					} else {
						program.SetFolder("/")
					}

					for key, val := range parseEnvironment(environment[process.PID]) {
						program.SetEnvironment(key, val)
					}

					user := bsdshadow.ToUser(process.UID)

					if user.Name != "" {
						matcher := matchers.NewUser()
						matcher.SetName(user.Name)
						matcher.SetType(user.Type)
						program.SetUser(matcher)
					}

					if strings.HasPrefix(command, "/") && program.HasFilesystem(command) == false {
						program.Filesystem = append(program.Filesystem, command)
					}

					for f := 0; f < len(info.Files); f++ {
						program.AddFilesystem(info.Files[f])
					}

					for s := 0; s < len(sockets); s++ {

						socket := sockets[s]

						if socket.PID == process.PID {

							connection := matchers.NewConnection()
							connection.SetProtocol(socket.Protocol)

							if socket.ForeignPort == 0 && socket.LocalPort > 0 {

								connection.SetHost(socket.Local)
								connection.SetPort(socket.LocalPort)
								connection.SetType("server")
								program.AddConnection(connection)

							} else if socket.LocalPort > 0 && socket.ForeignPort > 0 {

								connection.SetHost(socket.Foreign)
								connection.SetPort(socket.ForeignPort)
								connection.SetType("client")
								program.AddConnection(connection)

							}

						}

					}

					if program.PID != 0 && program.Name != "" {
						collected = append(collected, program)
					}

				}

			}

		}

	}

	sort.Slice(collected, func(a int, b int) bool {
		return collected[a].PID < collected[b].PID
	})

	return collected

}

