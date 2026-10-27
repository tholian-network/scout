package shadow

import "tholian-core/types"
import "os"
import "slices"
import "strconv"
import "strings"

func ToUser(uid uint16) types.User {

	var result types.User

	if SUPPORTED == true {

		buffer_passwd, err1 := os.ReadFile("/etc/passwd")
		buffer_groups, err2 := os.ReadFile("/etc/group")

		if err1 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer_passwd)), "\n")

			for l := 0; l < len(lines); l++ {

				chunks := strings.Split(strings.TrimSpace(lines[l]), ":")

				if len(chunks) == 7 {

					check, err11 := strconv.ParseUint(chunks[2], 10, 64)

					if err11 == nil && uint16(check) == uid {

						name := strings.TrimSpace(chunks[0])
						folder := strings.TrimSpace(chunks[5])
						shell := strings.TrimSpace(chunks[6])

						result = types.ToUser(name)
						result.SetID(65535)
						result.SetFolder(folder)
						result.SetShell(shell)

						uid, err11 := strconv.ParseUint(chunks[2], 10, 64)

						if err11 == nil {
							result.SetID(uint16(uid))
						}

						gid, err12 := strconv.ParseUint(chunks[3], 10, 64)

						if err12 == nil {

							default_group := types.NewGroup()
							default_group.ID = uint16(gid)
							result.Groups = append(result.Groups, default_group)

						}

						break

					}

				}

			}

		}

		if err2 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer_groups)), "\n")

			for l := 0; l < len(lines); l++ {

				chunks := strings.Split(strings.TrimSpace(lines[l]), ":")

				if len(chunks) == 4 {

					name := chunks[0]
					usernames := strings.Split(strings.TrimSpace(chunks[3]), ",")

					if slices.Contains(usernames, result.Name) {

						group := types.ToGroup(name)

						gid, err11 := strconv.ParseUint(chunks[2], 10, 64)

						if err11 == nil {
							group.SetID(uint16(gid))
						}

						result.Groups = append(result.Groups, group)

					}

				}

			}

		}

	}

	return result

}
