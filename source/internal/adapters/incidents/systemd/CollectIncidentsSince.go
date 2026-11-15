package systemd

import "tholian-core/structs"
import "tholian-core/types"
import "os/exec"

func CollectIncidentsSince(system string, datetime types.Datetime) []structs.Incident {

	var result []structs.Incident

	if SUPPORTED == true {

		since := datetime.String()

		cmd := exec.Command("journalctl", "--since", since)
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
