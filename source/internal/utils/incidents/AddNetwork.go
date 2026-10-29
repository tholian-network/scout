package incidents

import "tholian-core/matchers"
import "tholian-core/structs"
import "strings"

func AddNetwork(incident *structs.Incident, value string) bool {

	value = strings.TrimSpace(value)

	if value != "" && value != "any" {

		if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
			value = value[1 : len(value)-1]
		}

		subnet := structs.ToSubnet(value)
		network := matchers.NewNetwork()
		network.SetSubnet(subnet.String())

		if network.IsValid() {

			incident.AddNetwork(network)

			return true

		}

	}

	return false

}
