package proc

import "tholian-core/types"
import "strconv"

func parseConnection(hex_source string, hex_target string) types.Connection {

	var connection = types.NewConnection()

	if len(hex_source) == 13 && string(hex_source[8]) == ":" {

		ipv4_a, err_a := strconv.ParseUint(string(hex_source[6])+string(hex_source[7]), 16, 8)
		ipv4_b, err_b := strconv.ParseUint(string(hex_source[4])+string(hex_source[5]), 16, 8)
		ipv4_c, err_c := strconv.ParseUint(string(hex_source[2])+string(hex_source[3]), 16, 8)
		ipv4_d, err_d := strconv.ParseUint(string(hex_source[0])+string(hex_source[1]), 16, 8)
		port, err_p := strconv.ParseUint(string(hex_source[9])+string(hex_source[10])+string(hex_source[11])+string(hex_source[12]), 16, 16)

		if err_a == nil && err_b == nil && err_c == nil && err_d == nil && err_p == nil {
			socket := types.NewSocket(strconv.FormatUint(uint64(ipv4_a), 10)+"."+strconv.FormatUint(uint64(ipv4_b), 10)+"."+strconv.FormatUint(uint64(ipv4_c), 10)+"."+strconv.FormatUint(uint64(ipv4_d), 10), uint16(port))
			connection.SetSource(socket)
		}

	} else if len(hex_source) == 37 && string(hex_source[32]) == ":" {

		ipv6_1 := string(hex_source[6]) + string(hex_source[7]) + string(hex_source[4]) + string(hex_source[5])
		ipv6_2 := string(hex_source[2]) + string(hex_source[3]) + string(hex_source[0]) + string(hex_source[1])
		ipv6_3 := string(hex_source[14]) + string(hex_source[15]) + string(hex_source[12]) + string(hex_source[13])
		ipv6_4 := string(hex_source[10]) + string(hex_source[11]) + string(hex_source[8]) + string(hex_source[9])
		ipv6_5 := string(hex_source[22]) + string(hex_source[23]) + string(hex_source[20]) + string(hex_source[21])
		ipv6_6 := string(hex_source[18]) + string(hex_source[19]) + string(hex_source[16]) + string(hex_source[17])
		ipv6_7 := string(hex_source[30]) + string(hex_source[31]) + string(hex_source[28]) + string(hex_source[29])
		ipv6_8 := string(hex_source[26]) + string(hex_source[27]) + string(hex_source[24]) + string(hex_source[25])

		port, err_p := strconv.ParseUint(string(hex_source[33])+string(hex_source[34])+string(hex_source[35])+string(hex_source[36]), 16, 16)

		if err_p == nil {
			socket := types.NewSocket("["+ipv6_1+":"+ipv6_2+":"+ipv6_3+":"+ipv6_4+":"+ipv6_5+":"+ipv6_6+":"+ipv6_7+":"+ipv6_8+"]", uint16(port))
			connection.SetSource(socket)
		}

	}

	if len(hex_target) == 13 && string(hex_target[8]) == ":" {

		ipv4_a, err_a := strconv.ParseUint(string(hex_target[6])+string(hex_target[7]), 16, 8)
		ipv4_b, err_b := strconv.ParseUint(string(hex_target[4])+string(hex_target[5]), 16, 8)
		ipv4_c, err_c := strconv.ParseUint(string(hex_target[2])+string(hex_target[3]), 16, 8)
		ipv4_d, err_d := strconv.ParseUint(string(hex_target[0])+string(hex_target[1]), 16, 8)
		port, err_p := strconv.ParseUint(string(hex_target[9])+string(hex_target[10])+string(hex_target[11])+string(hex_target[12]), 16, 16)

		if err_a == nil && err_b == nil && err_c == nil && err_d == nil && err_p == nil {
			socket := types.NewSocket(strconv.FormatUint(uint64(ipv4_a), 10)+"."+strconv.FormatUint(uint64(ipv4_b), 10)+"."+strconv.FormatUint(uint64(ipv4_c), 10)+"."+strconv.FormatUint(uint64(ipv4_d), 10), uint16(port))
			connection.SetTarget(socket)
		}

	} else if len(hex_target) == 37 && string(hex_target[32]) == ":" {

		ipv6_1 := string(hex_target[6]) + string(hex_target[7]) + string(hex_target[4]) + string(hex_target[5])
		ipv6_2 := string(hex_target[2]) + string(hex_target[3]) + string(hex_target[0]) + string(hex_target[1])
		ipv6_3 := string(hex_target[14]) + string(hex_target[15]) + string(hex_target[12]) + string(hex_target[13])
		ipv6_4 := string(hex_target[10]) + string(hex_target[11]) + string(hex_target[8]) + string(hex_target[9])
		ipv6_5 := string(hex_target[22]) + string(hex_target[23]) + string(hex_target[20]) + string(hex_target[21])
		ipv6_6 := string(hex_target[18]) + string(hex_target[19]) + string(hex_target[16]) + string(hex_target[17])
		ipv6_7 := string(hex_target[30]) + string(hex_target[31]) + string(hex_target[28]) + string(hex_target[29])
		ipv6_8 := string(hex_target[26]) + string(hex_target[27]) + string(hex_target[24]) + string(hex_target[25])

		port, err_p := strconv.ParseUint(string(hex_target[33])+string(hex_target[34])+string(hex_target[35])+string(hex_target[36]), 16, 16)

		if err_p == nil {
			socket := types.NewSocket("["+ipv6_1+":"+ipv6_2+":"+ipv6_3+":"+ipv6_4+":"+ipv6_5+":"+ipv6_6+":"+ipv6_7+":"+ipv6_8+"]", uint16(port))
			connection.SetTarget(socket)
		}

	}

	return connection

}
