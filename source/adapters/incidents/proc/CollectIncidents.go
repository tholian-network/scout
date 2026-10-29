package proc

import "tholian-core/structs"
import utils_incidents "tholian-scout/internal/utils/incidents"
import "tholian-scout/adapters/programs/proc"
import "os"
import "sort"
import "time"

func CollectIncidents(system string) []structs.Incident {

	var result []structs.Incident

	if SUPPORTED == true {

		processes := proc.CollectProcesses()
		seen := make(map[string]bool)
		self := uint(os.Getpid())

		for p := 0; p < len(processes); p++ {

			program := processes[p]

			if program.PID == self {
				continue
			}

			datetime := time.Now().Format("2006-01-02 15:04:05")

			if started, ok := ProcessStartedAt(program.PID); ok == true {
				datetime = started.String()
			}

			incidents := AnalyzeProgram(system, datetime, program)

			for i := 0; i < len(incidents); i++ {

				incident := incidents[i]

				if incident.IsValid() == true {

					identity := incident.Hash()

					if identity == "" {
						identity = utils_incidents.ToIdentity(incident)
					}

					if seen[identity] == false {
						seen[identity] = true
						result = append(result, incident)
					}

				}

			}

		}

	}

	sort.SliceStable(result, func(a int, b int) bool {
		return result[a].Datetime.IsBefore(result[b].Datetime)
	})

	return result

}
