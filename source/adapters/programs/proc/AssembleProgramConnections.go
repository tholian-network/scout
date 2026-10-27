package proc

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "os"
import "strconv"
import "strings"

func AssembleProgramConnections(result *structs.Program, process_id uint) {

	// Optional argument
	if process_id == 0 {
		process_id = result.PID
	}

	if process_id != 0 {

		pid := strconv.FormatUint(uint64(process_id), 10)
		connections := make([]types.Connection, 0)
		descriptors := make([]string, 0)

		fds, err1 := os.ReadDir("/proc/" + pid + "/fd")

		if err1 == nil {

			for f := 0; f < len(fds); f++ {

				fd, err2 := os.Readlink("/proc/" + pid + "/fd/" + fds[f].Name())

				if err2 == nil {

					descriptor := strings.TrimSpace(string(fd))

					if strings.HasPrefix(descriptor, "socket:") {
						descriptors = append(descriptors, descriptor)
					}

				}

			}

		}

		buffer_tcp4, err_tcp4 := os.ReadFile("/proc/" + pid + "/net/tcp")
		buffer_tcp6, err_tcp6 := os.ReadFile("/proc/" + pid + "/net/tcp6")
		buffer_udp4, err_udp4 := os.ReadFile("/proc/" + pid + "/net/udp")
		buffer_udp6, err_udp6 := os.ReadFile("/proc/" + pid + "/net/udp6")

		if err_tcp4 == nil {
			connections = append(connections, parseConnectionBuffer(buffer_tcp4, descriptors, types.ProtocolTCP)...)
		}

		if err_tcp6 == nil {
			connections = append(connections, parseConnectionBuffer(buffer_tcp6, descriptors, types.ProtocolTCP)...)
		}

		if err_udp4 == nil {
			connections = append(connections, parseConnectionBuffer(buffer_udp4, descriptors, types.ProtocolUDP)...)
		}

		if err_udp6 == nil {
			connections = append(connections, parseConnectionBuffer(buffer_udp6, descriptors, types.ProtocolUDP)...)
		}

		if len(connections) > 0 {

			for c := 0; c < len(connections); c++ {

				connection := connections[c]

				if connection.Type == "client" {

					if connection.Target.Port != 0 {

						matcher := matchers.NewConnection()
						matcher.SetHost(connection.Target.Host)
						matcher.SetPort(connection.Target.Port)
						matcher.SetProtocol(connection.Protocol.String())
						matcher.SetType("client")

						result.AddConnection(matcher)

					}

				} else if connection.Type == "server" {

					if connection.Source.Port != 0 {

						matcher := matchers.NewConnection()

						if connection.Source.Type == "ipv4" && connection.Source.Host == "0.0.0.0" {
							matcher.SetHost("any")
						} else if connection.Source.Type == "ipv6" && connection.Source.Host == "[0000:0000:0000:0000:0000:0000:0000:0000]" {
							matcher.SetHost("any")
						} else {
							matcher.SetHost(connection.Source.Host)
						}

						matcher.SetPort(connection.Source.Port)
						matcher.SetProtocol(connection.Protocol.String())
						matcher.SetType("server")

						result.AddConnection(matcher)

					}

				} else if connection.Type == "peer" {

					if connection.Source.Port != 0 {

						matcher := matchers.NewConnection()

						if connection.Source.Type == "ipv4" && connection.Source.Host == "0.0.0.0" {
							matcher.SetHost("any")
						} else if connection.Source.Type == "ipv6" && connection.Source.Host == "[0000:0000:0000:0000:0000:0000:0000:0000]" {
							matcher.SetHost("any")
						} else {
							matcher.SetHost(connection.Source.Host)
						}

						matcher.SetPort(connection.Source.Port)
						matcher.SetProtocol(connection.Protocol.String())
						matcher.SetType("peer")

						result.AddConnection(matcher)

					}

				}

			}

		}

	}

}
