package pam

import "tholian-core/matchers"
import "tholian-core/structs"

func toReconIncident(system string, datetime string, ip string, username string) structs.Incident {

	var incident = structs.NewIncident(system, "Recon")

	incident.SetDatetime(datetime)

	subnet := structs.ToSubnet(ip)
	network := matchers.NewNetwork()
	network.SetSubnet(subnet.String())
	incident.AddNetwork(network)

	user := matchers.ToUser(username)
	incident.AddUser(user)

	return incident

}
