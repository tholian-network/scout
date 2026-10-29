package syslog

import "tholian-core/structs"
import utils_incidents "tholian-scout/internal/utils/incidents"
import "os/exec"
import "sort"

func CollectIncidents(system string) []structs.Incident {

	var result []structs.Incident

	if SUPPORTED == true {

		seen := make(map[string]bool)

		for f := 0; f < len(Files); f++ {

			buffer, err := readSyslogFile(Files[f])

			if err == nil && len(buffer) > 0 {

				incidents := ParseIncidents(buffer, system)

				if len(incidents) > 0 {

					for i := 0; i < len(incidents); i++ {

						incident := incidents[i]

						if incident.IsValid() == true {

							identity := incident.Hash()

							if identity == "" {
								identity = utils_incidents.ToIdentity(incident)
							}

							if seen[identity] == false {
								result = append(result, incident)
								seen[identity] = true
							}

						}

					}

				}

			}

		}

		if Logread != "" {

			cmd := exec.Command(Logread)
			buffer, err := cmd.Output()

			if err == nil && len(buffer) > 0 {

				incidents := ParseIncidents(buffer, system)

				if len(incidents) > 0 {

					for i := 0; i < len(incidents); i++ {

						incident := incidents[i]

						if incident.IsValid() == true {

							identity := incident.Hash()

							if identity == "" {
								identity = utils_incidents.ToIdentity(incident)
							}

							if seen[identity] == false {
								result = append(result, incident)
								seen[identity] = true
							}

						}

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
