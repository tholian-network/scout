package proc

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-scout/adapters/users/shadow"
import "os"
import "strconv"
import "strings"
import "syscall"

func AssembleProgram(result *structs.Program) {

	if result.PID != 0 {

		pid := strconv.FormatUint(uint64(result.PID), 10)

		stat, err0 := os.Stat("/proc/" + pid)
		exe, err1 := os.Readlink("/proc/" + pid + "/exe")
		cwd, err2 := os.Readlink("/proc/" + pid + "/cwd")
		buf_cmdline, err3 := os.ReadFile("/proc/" + pid + "/cmdline")
		buf_environ, err4 := os.ReadFile("/proc/" + pid + "/environ")

		if err0 == nil {

			passwd_user := shadow.ToUser(uint16(stat.Sys().(*syscall.Stat_t).Uid))

			user := matchers.NewUser()
			user.SetName(passwd_user.Name)
			user.SetType(passwd_user.Type)

			result.SetUser(user)

		}

		if err1 == nil {

			// exe is always empty for kernel processes
			command := strings.TrimSpace(string(exe))

			if command != "" {
				result.SetCommand(command)
			}

		}

		// Only interested in non-kernel processes
		if result.Command != "" {

			if err2 == nil {

				folder := strings.TrimSpace(string(cwd))

				if strings.HasSuffix(folder, " (deleted)") {
					folder = strings.TrimSpace(folder[0 : len(folder)-10])
				}

				if strings.HasPrefix(folder, "/") {
					result.SetFolder(folder)
				}

			}

			if err3 == nil {

				var arguments []string

				tmp := strings.Split(strings.TrimSpace(string(buf_cmdline)), "\x00")

				for t := 0; t < len(tmp); t++ {

					arg := strings.TrimSpace(tmp[t])

					if arg != "" {
						arguments = append(arguments, arg)
					}

				}

				result.SetArguments(arguments)

			}

			if err4 == nil {

				tmp := strings.Split(strings.TrimSpace(string(buf_environ)), "\x00")

				for t := 0; t < len(tmp); t++ {

					chunk := strings.TrimSpace(tmp[t])

					if strings.Contains(chunk, "=") {

						key := strings.TrimSpace(chunk[0:strings.Index(chunk, "=")])
						val := strings.TrimSpace(chunk[strings.Index(chunk, "=")+1:])

						result.SetEnvironment(key, val)

					}

				}

			}

			if strings.Contains(result.Command, "/") {

				name := strings.TrimSpace(result.Command[strings.LastIndex(result.Command, "/")+1:])

				if name != "" {
					result.SetName(name)
				}

			}

		}

	}

}
