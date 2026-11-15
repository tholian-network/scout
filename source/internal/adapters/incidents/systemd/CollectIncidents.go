package systemd

import "tholian-core/structs"
import "os/exec"

func CollectIncidents(system string) []structs.Incident {

	var result []structs.Incident

	if SUPPORTED == true {

		cmd := exec.Command("journalctl")
		buffer, err := cmd.Output()

		if err == nil {

			incidents := ParseIncidents(buffer, system)

			for i := 0; i < len(incidents); i++ {
				result = append(result, incidents[i])
			}

		}

	}

	return result

}
