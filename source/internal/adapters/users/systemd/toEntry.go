package systemd

import "strings"

func toEntry(line string) []string {

	var result []string

	if strings.HasPrefix(line, "g") {

		result = append(result, "g")
		line = strings.TrimSpace(line[1:])

		if strings.Contains(line, " ") {

			name := line[0:strings.Index(line, " ")]
			result = append(result, name)
			line = strings.TrimSpace(line[len(name):])

		}

		if strings.Contains(line, " ") {

			gid_or_path := line[0:strings.Index(line, " ")]
			result = append(result, gid_or_path)
			line = strings.TrimSpace(line[len(gid_or_path):])

		}

	} else if strings.HasPrefix(line, "m") {

		result = append(result, "m")
		line = strings.TrimSpace(line[1:])

		if strings.Contains(line, " ") {

			username := line[0:strings.Index(line, " ")]
			result = append(result, username)
			line = strings.TrimSpace(line[len(username):])

		}

		if strings.Contains(line, " ") {

			groupname := line[0:strings.Index(line, " ")]
			result = append(result, groupname)
			line = strings.TrimSpace(line[len(groupname):])

		} else {

			groupname := strings.TrimSpace(line)
			result = append(result, groupname)

		}

	} else if strings.HasPrefix(line, "u") {

		result = append(result, "u")
		line = strings.TrimSpace(line[1:])

		if strings.Contains(line, " ") {

			username := line[0:strings.Index(line, " ")]
			result = append(result, username)
			line = strings.TrimSpace(line[len(username):])

		}

		if strings.Contains(line, " ") {

			uid_or_path := line[0:strings.Index(line, " ")]
			result = append(result, uid_or_path)
			line = strings.TrimSpace(line[len(uid_or_path):])

			if strings.Contains(line, " ") {

				if strings.HasPrefix(line, "\"") {

					line = line[1:]
					description := line[0:strings.Index(line, "\"")]
					result = append(result, description)
					line = strings.TrimSpace(line[len(description):])

					if strings.HasPrefix(line, "\"") {
						line = strings.TrimSpace(line[1:])
					}

				} else {

					description := line[0:strings.Index(line, " ")]
					result = append(result, description)
					line = strings.TrimSpace(line[len(description):])

				}

				if strings.Contains(line, " ") {

					home := line[0:strings.Index(line, " ")]
					result = append(result, home)
					line = strings.TrimSpace(line[len(home):])

					if strings.Contains(line, " ") {

						shell := line[0:strings.Index(line, " ")]
						result = append(result, shell)
						line = strings.TrimSpace(line[len(shell):])

					} else {

						shell := strings.TrimSpace(line)
						result = append(result, shell)

					}

				}

			}

		} else {

			uid_or_path := strings.TrimSpace(line)
			result = append(result, uid_or_path)

		}

	} else if strings.HasPrefix(line, "r") {

		result = append(result, "u")
		line = strings.TrimSpace(line[1:])

		if strings.Contains(line, " ") {

			uid_gid_range := line[0:strings.Index(line, " ")]
			result = append(result, uid_gid_range)

		} else {

			uid_gid_range := strings.TrimSpace(line)
			result = append(result, uid_gid_range)

		}

	}

	return result

}
