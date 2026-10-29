package pam

import "tholian-core/matchers"
import "tholian-core/structs"

func toPersistIncident(system string, datetime string, username string) structs.Incident {

	var incident = structs.NewIncident(system, "Persist")

	incident.SetDatetime(datetime)

	user := matchers.ToUser(username)
	incident.AddUser(user)

	return incident

}
