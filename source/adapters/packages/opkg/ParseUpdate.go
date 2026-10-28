package opkg

import "tholian-core/structs"
import "strings"

func ParseUpdate(buffer string, result *structs.Update) {

	line := strings.TrimSpace(buffer)
	chunks := strings.Split(line, " - ")

	if len(chunks) >= 3 {

		result.SetName(strings.TrimSpace(chunks[0]))
		result.SetVersion(strings.TrimSpace(chunks[len(chunks)-1]))

	}

}
