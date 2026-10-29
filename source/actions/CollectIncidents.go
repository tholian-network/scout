package actions

import "tholian-core/structs"
import "tholian-scout/adapters/incidents/proc"
import "tholian-scout/adapters/incidents/syslog"
import "tholian-scout/adapters/incidents/systemd"
import utils_incidents "tholian-scout/internal/utils/incidents"
import "strconv"

func CollectIncidents(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectIncidents")

	collected := make([]structs.Incident, 0)
	seen := make(map[string]bool)

	if proc.SUPPORTED == true {

		tmp := proc.CollectIncidents(system.Name)

		if len(tmp) > 0 {

			console.Info("proc.CollectIncidents(): Found " + strconv.Itoa(len(tmp)) + " Incidents")

			for _, incident := range tmp {

				if incident.IsValid() {

					identity := utils_incidents.ToIdentity(incident)

					if seen[identity] == false {
						collected = append(collected, incident)
						seen[identity] = true
					}

				}

			}

		} else {

			console.Warn("proc.CollectIncidents(): Found 0 Incidents")

		}

		result = true

	} else {

		console.Warn("proc.CollectIncidents(): Unsupported")

	}

	if systemd.SUPPORTED == true {

		tmp := systemd.CollectIncidents(system.Name)

		if len(tmp) > 0 {

			console.Info("systemd.CollectIncidents(): Found " + strconv.Itoa(len(tmp)) + " Incidents")

			for _, incident := range tmp {

				if incident.IsValid() {

					identity := utils_incidents.ToIdentity(incident)

					if seen[identity] == false {
						collected = append(collected, incident)
						seen[identity] = true
					}

				}

			}

		} else {

			console.Warn("systemd.CollectIncidents(): Found 0 Incidents")

		}

		result = true

	} else {

		console.Warn("systemd.CollectIncidents(): Unsupported")

	}

	if syslog.SUPPORTED == true {

		tmp := syslog.CollectIncidents(system.Name)

		if len(tmp) > 0 {

			console.Info("syslog.CollectIncidents(): Found " + strconv.Itoa(len(tmp)) + " Incidents")

			for _, incident := range tmp {

				if incident.IsValid() {

					identity := utils_incidents.ToIdentity(incident)

					if seen[identity] == false {
						collected = append(collected, incident)
						seen[identity] = true
					}

				}

			}

		} else {

			console.Warn("syslog.CollectIncidents(): Found 0 Incidents")

		}

		result = true

	} else {

		console.Warn("syslog.CollectIncidents(): Unsupported")

	}

	system.SetIncidents(collected)

	console.Log("Collected " + strconv.Itoa(len(system.Incidents)) + "/" + strconv.Itoa(len(collected)) + " Incidents")
	console.GroupEnd("actions/CollectIncidents")

	return result

}
