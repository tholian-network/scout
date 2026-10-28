package tdnf

import "tholian-core/structs"
import "os"
import "os/exec"

func CollectUpdate(name string) structs.Update {

	var result structs.Update = structs.NewUpdate("tdnf")

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("tdnf", "--noplugins", "list", "updates", name)
		buffer, err := cmd.Output()

		if err == nil {

			updates := parseUpdates(string(buffer), name)

			if len(updates) > 0 {
				result = updates[0]
			}

		}

	}

	return result

}
