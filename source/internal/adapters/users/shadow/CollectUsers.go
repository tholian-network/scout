package shadow

import "tholian-core/types"
import "os"
import "sort"
import "strconv"
import "strings"

func CollectUsers() []types.User {

	var collected []types.User

	if SUPPORTED == true {

		user_index := make(map[string]*types.User, 0)
		group_index := make(map[string]*types.Group, 0)

		buffer_passwd, err1 := os.ReadFile("/etc/passwd")
		buffer_groups, err2 := os.ReadFile("/etc/group")
		buffer_shadow, err3 := os.ReadFile("/etc/shadow")
		buffer_gshadow, err4 := os.ReadFile("/etc/gshadow")

		if err1 == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer_passwd)), "\n")

			for l := 0; l < len(lines); l++ {

				chunks := strings.Split(strings.TrimSpace(lines[l]), ":")

				if len(chunks) == 7 {

					name := strings.TrimSpace(chunks[0])
					password := strings.TrimSpace(chunks[1])
					folder := strings.TrimSpace(chunks[5])
					shell := strings.TrimSpace(chunks[6])

					if password == "x" {

						if err3 == nil {
							password = unShadow(buffer_shadow, name)
						} else {
							password = ""
						}

					}

					user := types.ToUser(name)
					user.SetID(65535)
					user.SetPassword(password)
					user.SetFolder(folder)
					user.SetShell(shell)

					uid, err11 := strconv.ParseUint(chunks[2], 10, 64)

					if err11 == nil {
						user.SetID(uint16(uid))
					}

					gid, err12 := strconv.ParseUint(chunks[3], 10, 64)

					if err12 == nil {

						default_group := types.NewGroup()
						default_group.ID = uint16(gid)
						user.Groups = append(user.Groups, default_group)

					}

					if user.Name != "" {
						user_index[user.Name] = &user
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
					password := chunks[1]
					usernames := strings.Split(strings.TrimSpace(chunks[3]), ",")

					if len(usernames) == 1 && usernames[0] == "" {
						usernames = []string{}
					}

					group := types.ToGroup(name)

					if password == "x" {

						if err4 == nil {
							group.SetPassword(unShadow(buffer_gshadow, name))
						} else {
							group.SetPassword("")
						}

					} else {
						group.SetPassword(password)
					}

					gid, err11 := strconv.ParseUint(chunks[2], 10, 64)

					if err11 == nil {
						group.SetID(uint16(gid))
					}

					if group.Name != "" {
						group_index[group.Name] = &group
					}

					for u := 0; u < len(usernames); u++ {

						user, ok := user_index[usernames[u]]

						if ok == true {
							user.Groups = append(user.Groups, group)
						}

					}

				}

			}

		}

		names := make([]string, 0)

		for name, user := range user_index {

			if user.Groups[0].Name == "" {

				for _, group := range group_index {

					if group.ID == user.Groups[0].ID {

						user.Groups[0].Name = group.Name
						user.Groups[0].Password = group.Password

						break

					}

				}

			}

			names = append(names, name)

		}

		sort.Strings(names)

		for n := 0; n < len(names); n++ {

			user, ok := user_index[names[n]]

			if ok == true {
				collected = append(collected, *user)
			}

		}

	}

	return collected

}
