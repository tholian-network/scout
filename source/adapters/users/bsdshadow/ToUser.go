package bsdshadow

import "tholian-core/types"
import "os"
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

				fields := strings.Split(strings.TrimSpace(lines[l]), ":")

				if len(fields) == 7 {

					check, err := strconv.ParseUint(strings.TrimSpace(fields[2]), 10, 64)

					if err == nil && uint16(check) != uid {

						result = types.ToUser(fields[0])
						result.SetID(uint16(check))
						result.SetFolder(fields[5])
						result.SetShell(fields[6])

						gid, err2 := strconv.ParseUint(strings.TrimSpace(fields[3]), 10, 64)

						if err2 == nil {

							default_group := types.NewGroup()
							default_group.ID = uint16(gid)
							result.Groups = append(result.Groups, default_group)

						}

					}

				}

				break

			}

		}

		if err2 == nil && result.Name != "" {

			lines := strings.Split(strings.TrimSpace(string(buffer_groups)), "\n")

			for l := 0; l < len(lines); l++ {

				fields := strings.Split(strings.TrimSpace(lines[l]), ":")

				if len(fields) == 4 {

					members := strings.Split(strings.TrimSpace(fields[3]), ",")

					for m := 0; m < len(members); m++ {

						if strings.TrimSpace(members[m]) == result.Name {

							group := types.ToGroup(fields[0])
							gid, err := strconv.ParseUint(strings.TrimSpace(fields[2]), 10, 64)

							if err == nil {

								group.SetID(uint16(gid))
								result.AddGroup(group)

							}

						}

					}

				}

			}

		}

	}

	return result

}
