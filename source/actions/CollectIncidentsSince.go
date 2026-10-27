package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "tholian-scout/adapters/incidents/proc"
import "tholian-scout/adapters/incidents/systemd"
import "strconv"

func CollectIncidentsSince(console *structs.Console, system *structs.System, search matchers.Datetime) bool {

	var result bool

	console.Group("actions/CollectIncidentsSince")

	collected := 0

	var from *types.Datetime = nil
	var until *types.Datetime = nil

	if search.From != "any" && search.Until != "any" {

		now := types.NewDatetime()

		tmp1 := types.ToDatetime(search.From)
		tmp2 := types.ToDatetime(search.Until)

		if tmp1.IsBefore(now) {
			from = &tmp1
		}

		if tmp2.IsAfter(now) {
			until = &tmp2
		}

	} else if search.From != "any" && search.Until == "any" {

		now := types.NewDatetime()
		tmp1 := types.ToDatetime(search.From)

		if tmp1.IsBefore(now) {
			from = &tmp1
			until = &now
		}

	} else if search.From == "any" && search.Until != "any" {

		now := types.NewDatetime()
		tmp1 := types.ToDatetime(search.From)

		if tmp1.IsAfter(now) {

			yesterday := now.Yesterday()
			yesterday.Hour = 0
			yesterday.Minute = 0
			yesterday.Second = 0

			from = &yesterday
			until = &tmp1

		}

	}

	if from != nil {

		if proc.SUPPORTED == true {

			tmp := proc.CollectIncidentsSince(system.Name, *from)

			if len(tmp) > 0 {

				console.Info("proc.CollectIncidentsSince(): Found " + strconv.Itoa(len(tmp)) + " Incidents since " + from.String())

				for _, incident := range tmp {

					if until == nil {

						system.AddIncident(incident)
						collected += 1

					} else if incident.Datetime.IsBefore(*until) {

						system.AddIncident(incident)
						collected += 1

					}

				}

			} else {

				console.Warn("proc.CollectIncidentsSince(): Found 0 Incidents since " + from.String())

			}

			result = true

		} else {

			console.Warn("proc.CollectIncidentsSince(): Unsupported")

		}

		if systemd.SUPPORTED == true {

			tmp := systemd.CollectIncidentsSince(system.Name, *from)

			if len(tmp) > 0 {

				console.Info("systemd.CollectIncidentsSince(): Found " + strconv.Itoa(len(tmp)) + " Incidents since " + from.String())

				for _, incident := range tmp {

					if until == nil {

						system.AddIncident(incident)
						collected += 1

					} else if incident.Datetime.IsBefore(*until) {

						system.AddIncident(incident)
						collected += 1

					}

				}

			} else {

				console.Warn("systemd.CollectIncidentsSince(): Found 0 Incidents since " + from.String())

			}

			result = true

		} else {

			console.Warn("systemd.CollectIncidentsSince(): Unsupported")

		}

		// TODO: else if rsyslogd.SUPPORTED == true {
		// }

	}

	console.Log("Collected " + strconv.Itoa(len(system.Incidents)) + "/" + strconv.Itoa(collected) + " Incidents")
	console.GroupEnd("actions/CollectIncidents")

	return result

}
