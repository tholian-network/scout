package actions

import "tholian-core/structs"
import "tholian-scout/adapters/incidents/proc"
import "tholian-scout/adapters/incidents/systemd"
import "strconv"

func CollectIncidents(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectIncidents")

	collected := make([]structs.Incident, 0)

	if proc.SUPPORTED == true {

		tmp := proc.CollectIncidents(system.Name)

		if len(tmp) > 0 {

			console.Info("proc.CollectIncidents(): Found " + strconv.Itoa(len(tmp)) + " Incidents")

			for _, incident := range tmp {

				if incident.IsValid() {
					collected = append(collected, incident)
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
					collected = append(collected, incident)
				}

			}

		} else {

			console.Warn("systemd.CollectIncidents(): Found 0 Incidents")

		}

		result = true

	} else {

		console.Warn("systemd.CollectIncidents(): Unsupported")

	}

	// TODO: else if rsyslogd.SUPPORTED == true {
	// }

	system.SetIncidents(collected)

	console.Log("Collected " + strconv.Itoa(len(system.Incidents)) + "/" + strconv.Itoa(len(collected)) + " Incidents")
	console.GroupEnd("actions/CollectIncidents")

	return result

}
