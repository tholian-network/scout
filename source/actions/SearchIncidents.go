package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchIncidents(console *structs.Console, system *structs.System, search matchers.Incident) []structs.Incident {

	console.Group("actions/SearchIncidents")

	var filtered []structs.Incident

	if search.MatchesSystem(system.Name) {

		for i := 0; i < len(system.Incidents); i++ {

			incident := system.Incidents[i]

			if search.MatchesType(incident.Type) && search.MatchesDatetime(incident.Datetime.String()) {
				filtered = append(filtered, incident)
			}

		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Incidents)) + " Incidents")
	console.GroupEnd("actions/SearchIncident")

	return filtered

}
