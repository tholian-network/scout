package bsdshadow

import "tholian-core/types"
import "os"
import "sort"
import "strconv"
import "strings"

func CollectUsers() []types.User {

	var collected []types.User

	if SUPPORTED == true {

		buffer_master, _ := os.ReadFile("/etc/master.passwd")
		buffer_passwd, _ := os.ReadFile("/etc/passwd")
		buffer_groups, _ := os.ReadFile("/etc/group")

		user_index := make(map[string]*types.User)
		group_index := make(map[string]types.Group)
		group_id_index := make(map[uint16]types.Group)
		names := make([]string, 0)

		buffer := buffer_master

		if len(strings.TrimSpace(string(buffer))) == 0 {
			buffer = buffer_passwd
		}

		lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(buffer)), "\r\n", "\n"), "\n")

		for l := 0; l < len(lines); l++ {

			fields := strings.Split(strings.TrimSpace(lines[l]), ":")

			if len(fields) == 7 || len(fields) == 10 {

				name := strings.TrimSpace(fields[0])
				password := strings.TrimSpace(fields[1])

				if name != "" {

					uid, err1 := strconv.ParseUint(strings.TrimSpace(fields[2]), 10, 64)
					gid, err2 := strconv.ParseUint(strings.TrimSpace(fields[3]), 10, 64)

					if err1 == nil && err2 == nil {

						folder := ""
						shell := ""

						if len(fields) == 10 {
							folder = strings.TrimSpace(fields[8])
							shell = strings.TrimSpace(fields[9])
						} else {
							folder = strings.TrimSpace(fields[5])
							shell = strings.TrimSpace(fields[6])
						}

						user := types.ToUser(name)
						user.SetID(65535)
						user.SetPassword(password)
						user.SetFolder(folder)
						user.SetShell(shell)
						user.SetID(uint16(uid))

						default_group := types.NewGroup()
						default_group.ID = uint16(gid)
						user.Groups = append(user.Groups, default_group)

						user_index[user.Name] = &user
						names = append(names, user.Name)

					}

				}

			}

		}

		group_lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(buffer_groups)), "\r\n", "\n"), "\n")

		for l := 0; l < len(group_lines); l++ {

			fields := strings.Split(strings.TrimSpace(group_lines[l]), ":")

			if len(fields) == 4 {

				name := strings.TrimSpace(fields[0])
				password := strings.TrimSpace(fields[1])

				if name != "" {

					gid, err := strconv.ParseUint(strings.TrimSpace(fields[2]), 10, 64)

					if err == nil {

						group := types.ToGroup(name)
						group.SetID(uint16(gid))
						group.SetPassword(password)

						group_index[name] = group
						group_id_index[uint16(gid)] = group

					}

				}

			}

		}

		for _, user := range user_index {

			if len(user.Groups) > 0 {

				default_id := user.Groups[0].ID

				if group, ok := group_id_index[default_id]; ok == true {
					user.Groups[0].Name = group.Name
					user.Groups[0].Password = group.Password
					user.Groups[0].Type = group.Type
				}

			}

		}

		for l := 0; l < len(group_lines); l++ {

			fields := strings.Split(strings.TrimSpace(group_lines[l]), ":")

			if len(fields) == 4 {

				group, ok := group_index[strings.TrimSpace(fields[0])]

				if ok == true {

					members := strings.Split(strings.TrimSpace(fields[3]), ",")

					for m := 0; m < len(members); m++ {

						name := strings.TrimSpace(members[m])

						if name != "" {

							if user, ok := user_index[name]; ok == true {
								user.AddGroup(group)
							}

						}

					}

				}

			}

		}

		sort.Strings(names)

		for n := 0; n < len(names); n++ {

			if user, ok := user_index[names[n]]; ok == true {
				collected = append(collected, *user)
			}

		}

	}

	return collected

}
