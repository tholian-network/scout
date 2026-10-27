package pacman

import "strings"

func splitShortName(shortname string) (string, string) {

	name := make([]string, 0)
	version := make([]string, 0)

	chunks := strings.Split(shortname, "-")

	if len(chunks) >= 1 {

		mode := "name"

		for c := 0; c < len(chunks); c++ {

			chunk := chunks[c]

			if mode == "name" {

				if strings.Contains(chunk, ".") {
					version = append(version, chunk)
					mode = "version"
				} else {
					name = append(name, chunk)
				}

			} else if mode == "version" {
				version = append(version, chunk)
			}

		}

	}

	return strings.Join(name, "-"), strings.Join(version, "-")

}
