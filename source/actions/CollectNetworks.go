package actions

import "tholian-core/structs"
import "net"
import "strconv"

func CollectNetworks(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectNetworks")

	collected := make([]structs.Network, 0)

	ifaces, err1 := net.Interfaces()

	if err1 == nil {

		for _, iface := range ifaces {

			addresses, err2 := iface.Addrs()
			network := structs.NewNetwork(iface.Name)

			if err2 == nil {

				for _, raw := range addresses {

					switch value := raw.(type) {

					case *net.IPAddr:

						// Do Nothing

					case *net.IPNet:

						address_v4 := value.IP.To4()
						address_v6 := value.IP.To16()

						if address_v6 != nil {

							prefix, _ := value.Mask.Size()
							subnet := structs.NewSubnet(value.IP.String(), uint8(prefix))
							network.AddSubnet(subnet)

						} else if address_v4 != nil {

							prefix, _ := value.Mask.Size()
							subnet := structs.NewSubnet(value.IP.String(), uint8(prefix))
							network.AddSubnet(subnet)

						}

					}

				}

			}

			if network.IsValid() {
				collected = append(collected, network)
			}

		}

		system.SetNetworks(collected)
		result = true

	}

	console.Log("Collected " + strconv.Itoa(len(system.Networks)) + "/" + strconv.Itoa(len(collected)) + " Networks")
	console.GroupEnd("actions/CollectNetworks")

	return result

}
