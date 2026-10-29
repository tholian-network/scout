package gogs

import "tholian-core/matchers"
import "tholian-core/structs"

func toReconIncident(system string, datetime string, ip string) structs.Incident {

	var incident = structs.NewIncident(system, "Recon")

	incident.SetDatetime(datetime)

	subnet := structs.ToSubnet(ip)
	network := matchers.NewNetwork()
	network.SetSubnet(subnet.String())
	incident.AddNetwork(network)

	program := matchers.NewProgram()
	program.SetName("gogs")
	incident.AddProgram(program)

	return incident

}
