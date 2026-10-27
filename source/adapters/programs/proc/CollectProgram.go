package proc

import "tholian-core/structs"
import "os"
import "sort"
import "strconv"
import "strings"

func CollectProgram(name string) []structs.Program {

	var collected []structs.Program

	if SUPPORTED == true {

		folders, err0 := os.ReadDir("/proc")

		pids := make([]uint64, 0)
		threads := make(map[uint64]uint64)

		if err0 == nil {

			for f := 0; f < len(folders); f++ {

				pid_str := folders[f].Name()
				check := string(pid_str[0])

				if check >= "0" && check <= "9" {

					pid, err1 := strconv.ParseUint(pid_str, 10, 64)

					if err1 == nil {

						exe, err2 := os.Readlink("/proc/" + pid_str + "/exe")

						if err2 == nil {

							command := strings.TrimSpace(string(exe))

							if command != "" {

								if (strings.HasPrefix(name, "/") && command == name) || strings.HasSuffix(command, "/"+name) {

									tasks, err3 := os.ReadDir("/proc/" + pid_str + "/task")

									if err3 == nil {

										for t := 0; t < len(tasks); t++ {

											tid, err := strconv.ParseUint(tasks[t].Name(), 10, 64)

											if err == nil {

												if tid > pid {
													threads[tid] = pid
												}

											}

										}

									}

									pids = append(pids, pid)

								}

							}

						}

					}

				}

			}

		}

		sort.Slice(pids, func(a int, b int) bool {
			return pids[a] < pids[b]
		})

		for p := 0; p < len(pids); p++ {

			pid := pids[p]
			owner_pid, is_thread := threads[pid]

			if is_thread == true && owner_pid != 0 {

				program := Programs.Get(uint(owner_pid))

				if program != nil {

					AssembleProgramConnections(program, uint(pid))
					AssembleProgramFilesystem(program, uint(pid))

				}

			} else if is_thread == false {

				program := Programs.Get(uint(pid))

				if program != nil {

					AssembleProgramConnections(program, uint(pid))
					AssembleProgramFilesystem(program, uint(pid))

				} else {

					program := structs.NewProgram(uint(pid), "")

					AssembleProgram(&program)
					AssembleProgramConnections(&program, uint(pid))
					AssembleProgramFilesystem(&program, uint(pid))

					if program.PID != 0 && program.Name != "" {
						Programs.Add(program)
					}

				}

			}

		}

		for p := 0; p < len(pids); p++ {

			pid := pids[p]
			_, is_thread := threads[pid]

			if is_thread == false {

				program := Programs.Get(uint(pid))

				if program != nil {
					collected = append(collected, *program)
				}

			}

		}

	}

	return collected

}
