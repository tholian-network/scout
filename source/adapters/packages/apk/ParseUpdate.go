package apk

import "tholian-core/structs"
import "strings"

func ParseUpdate(buffer string, result *structs.Update) {

	blocks := strings.Split(strings.TrimSpace(buffer), "\n\n")

	if len(blocks) > 0 {

		name, version := toNameAndVersion(blocks[0][0:strings.Index(blocks[0], " ")])

		result.SetName(name)
		result.SetVersion(version)

		for b := 0; b < len(blocks); b++ {

			lines := strings.Split(strings.TrimSpace(blocks[b]), "\n")

			if len(lines) > 0 {

				first_line := lines[0]

				if strings.HasSuffix(first_line, " description:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " webpage:") {

					if len(lines) == 2 {
						result.SetURL(strings.TrimSpace(lines[1]))
					}

				} else if strings.HasSuffix(first_line, " installed size:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " depends on:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " provides:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " is required by:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " contains:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " triggers:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " has auto-install rule:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " affects auto-installation of:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " replaces:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " license:") {
					// Do Nothing
				}

			}

		}

	}

}
