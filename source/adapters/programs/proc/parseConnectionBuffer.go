package proc

import "tholian-core/types"
import utils_strings "tholian-scout/internal/utils/strings"
import "slices"
import "strings"

func parseConnectionBuffer(buffer []byte, descriptors []string, protocol types.Protocol) []types.Connection {

	var result []types.Connection

	lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

	if len(lines) > 1 {

		headers := utils_strings.Split(lines[0], " ")

		if headers[0] == "sl" &&
			(headers[1] == "local_address" || headers[1] == "loc_address") &&
			(headers[2] == "remote_address" || headers[2] == "rem_address") &&
			headers[3] == "st" &&
			headers[4] == "tx_queue" &&
			headers[5] == "rx_queue" &&
			headers[6] == "tr" &&
			headers[7] == "tm->when" &&
			headers[8] == "retrnsmt" &&
			headers[9] == "uid" &&
			headers[10] == "timeout" &&
			headers[11] == "inode" {

			for l := 1; l < len(lines); l++ {

				chunks := utils_strings.Split(lines[l], " ")

				if len(chunks) >= 10 {

					// inode is offset by 2 due to "tx_queue rx_queue" and "tr tm->when"
					inode := chunks[9]

					if slices.Contains(descriptors, "socket:["+inode+"]") {

						if chunks[3] == "01" {

							// 01: Established
							connection := parseConnection(chunks[1], chunks[2])
							connection.SetProtocol(protocol)
							connection.SetType("client")

							result = append(result, connection)

						} else if chunks[3] == "0A" {

							// 0A: Listening
							connection := parseConnection(chunks[1], chunks[2])
							connection.SetProtocol(protocol)
							connection.SetType("server")

							result = append(result, connection)

						} else if chunks[3] == "07" {

							// 07: Close (used for UDP)
							connection := parseConnection(chunks[1], chunks[2])
							connection.SetProtocol(protocol)
							connection.SetType("peer")

							result = append(result, connection)

						}

					}

				}

			}

		}

	}

	return result

}
