package shadow

import "strings"

func unShadow(buffer []byte, name string) string {

	var result string = ""

	lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

	for l := 0; l < len(lines); l++ {

		chunks := strings.Split(strings.TrimSpace(lines[l]), ":")

		if len(chunks) == 5 {

			groupname := strings.TrimSpace(chunks[0])
			password := strings.TrimSpace(chunks[1])
			// groupadmin := strings.TrimSpace(chunks[2])
			members := strings.Split(strings.TrimSpace(chunks[3]), ",")

			if name == groupname {

				result = unHash(password)
				break

			} else {

				for m := 0; m < len(members); m++ {

					if name == members[m] {
						result = unHash(password)
						break
					}

				}

				break

			}

		} else if len(chunks) == 9 {

			username := strings.TrimSpace(chunks[0])
			password := strings.TrimSpace(chunks[1])
			// last_change, err1 := strconv.ParseUint(strings.TrimSpace(chunks[2]), 10, 64);
			// min_change,  err2 := strconv.ParseUint(strings.TrimSpace(chunks[3]), 10, 64);
			// max_change,  err3 := strconv.ParseUint(strings.TrimSpace(chunks[4]), 10, 64);
			// warn_change, err4 := strconv.ParseUint(strings.TrimSpace(chunks[5]), 10, 64);
			// lock_change, err5 := strconv.ParseUint(strings.TrimSpace(chunks[6]), 10, 64);
			// expires,     err6 := strconv.ParseUint(strings.TrimSpace(chunks[7]), 10, 64);
			// reserved := strings.TrimSpace(tmp[8]);

			if name == username {
				result = password
				break
			}

		}

	}

	return result

}
