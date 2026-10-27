package proc

import "tholian-core/structs"
import "os"
import "sort"
import "strconv"
import "strings"

func refresh() {

	pidcur := StatPID()

	if pidcur < MaxPID {

		folders, err0 := os.ReadDir("/proc")

		var kthreadd uint64 = 0

		pids := make([]uint64, 0)
		pthreads := make(map[uint64]uint64)
		kthreads := make(map[uint64]uint64)

		if err0 == nil {

			var pid uint64

			for pid = 0; pid <= 10; pid++ {

				comm_buf, err2 := os.ReadFile("/proc/" + strconv.FormatUint(pid, 10) + "/comm")

				if err2 == nil {

					comm := strings.TrimSpace(string(comm_buf))

					if comm == "kthreadd" {
						kthreadd = pid
						break
					}

				}

			}

			for f := 0; f < len(folders); f++ {

				name := folders[f].Name()
				check := string(name[0])

				if check >= "0" && check <= "9" {

					pid, err1 := strconv.ParseUint(name, 10, 64)

					if err1 == nil {

						if kthreadd != 0 && pid != kthreadd {

							stat_buf, err2 := os.ReadFile("/proc/" + name + "/stat")

							if err2 == nil {

								stat := strings.TrimSpace(string(stat_buf))

								// <pid> (process name with whitespace) S ppid ...
								if strings.Contains(stat, " (") && strings.Contains(stat, ") ") {

									tmp := strings.TrimSpace(stat[strings.Index(stat, ") ")+4:])
									ppid, err3 := strconv.ParseUint(tmp[0:strings.Index(tmp, " ")], 10, 64)

									if err3 == nil {

										if ppid == kthreadd {

											kthreads[pid] = ppid

										} else {

											_, err4 := os.Readlink("/proc/" + name + "/exe")

											if err4 == nil {
												pids = append(pids, pid)
											}

										}

									}

								}

							} else {
								// Unknown process
							}

						} else {

							_, err2 := os.Readlink("/proc/" + name + "/exe")

							if err2 == nil {
								pids = append(pids, pid)
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
			pid_str := strconv.FormatUint(pid, 10)
			ppid, is_pthread := pthreads[pid]

			if is_pthread == true && ppid != 0 {

				program := Programs.Get(uint(ppid))

				if program != nil {

					AssembleProgramConnections(program, uint(pid))
					AssembleProgramFilesystem(program, uint(pid))

				}

			} else if is_pthread == false {

				_, err1 := os.Stat("/proc/" + pid_str)

				if err1 == nil {

					// Ignore threads/tasks/forks of this process (temporarily, in this loop)
					tasks, err2 := os.ReadDir("/proc/" + pid_str + "/task")

					if err2 == nil {

						for t := 0; t < len(tasks); t++ {

							tid, err := strconv.ParseUint(tasks[t].Name(), 10, 64)

							if err == nil {

								if tid > pid {
									pthreads[tid] = pid
								}

							}

						}

					}

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

		}

	}

}
