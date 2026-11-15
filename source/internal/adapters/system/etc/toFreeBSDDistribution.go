package etc

import "tholian-core/structs"
import "os/exec"
import "strings"

func toFreeBSDDistribution() structs.Distribution {

	distribution := structs.NewDistribution()
	distribution.SetName("freebsd")
	distribution.SetVendor("freebsd")
	distribution.SetManager("pkg")

	version := ""

	cmd1 := exec.Command("freebsd-version", "-u")
	buffer1, err1 := cmd1.Output()

	if err1 == nil {
		version = strings.TrimSpace(string(buffer1))
	}

	if version == "" {

		cmd2 := exec.Command("uname", "-r")
		buffer2, err2 := cmd2.Output()

		if err2 == nil {
			version = strings.TrimSpace(string(buffer2))
		}

	}

	if version != "" {
		distribution.SetVersion(version)
	}

	return distribution

}
