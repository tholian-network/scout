package systemd

import "tholian-core/types"
import "os"
import "sort"
import "strings"

func CollectUsers() []types.User {

	var collected []types.User

	if SUPPORTED == true {

		user_index := make(map[string]*types.User)
		group_index := make(map[string]*types.Group)

		root := types.ToUser("root")
		root.SetID(0)
		user_index["root"] = &root

		stat, err1 := os.Stat("/usr/lib/sysusers.d")

		if err1 == nil && stat != nil && stat.IsDir() {

			files, err2 := os.ReadDir("/usr/lib/sysusers.d")

			if err2 == nil {

				for _, file := range files {

					filename := file.Name()

					if strings.HasSuffix(filename, ".conf") {

						buffer, err3 := os.ReadFile("/usr/lib/sysusers.d/" + filename)

						if err3 == nil {

							lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

							for l := 0; l < len(lines); l++ {

								line := strings.TrimSpace(lines[l])

								if strings.Contains(line, "#") {
									line = strings.TrimSpace(line[0:strings.Index(line, "#")])
								}

								if len(line) > 0 {

									entry := toEntry(line)

									if len(entry) > 0 {

										if entry[0] == "u" {

											if len(entry) >= 2 {

												user := types.ToUser(entry[1])
												user.SetID(65535)
												user.SetType("program")

												if len(entry) >= 3 && entry[2] != "-" && strings.HasPrefix(entry[2], "/") == false {

													uid := toUID(entry[2])

													if uid == 0 && user.Name == "root" {
														user.SetID(0)
													} else if uid != 0 {
														user.SetID(uid)
													}

												}

												if len(entry) >= 5 && strings.HasPrefix(entry[4], "/") {
													user.SetFolder(entry[4])
												}

												if len(entry) >= 6 && strings.HasPrefix(entry[5], "/") {
													user.SetShell(entry[5])
												}

												user_index[user.Name] = &user

											}

										} else if entry[0] == "g" {

											if len(entry) >= 2 {

												group := types.ToGroup(entry[1])
												group.SetType("program")

												if len(entry) >= 3 && entry[2] != "-" && strings.HasPrefix(entry[2], "/") == false {

													gid := toGID(entry[2])

													if gid == 0 && group.Name == "root" {
														group.SetID(0)
													} else if gid != 0 {
														group.SetID(gid)
													}

												}

												group_index[group.Name] = &group

											}

										} else if entry[0] == "m" {

											if len(entry) == 3 && entry[1] != "-" && entry[2] != "-" {

												name_user := entry[1]
												name_group := entry[2]

												user, _ := user_index[name_user]
												group, ok2 := group_index[name_group]

												if ok2 == false {

													tmp := types.ToGroup(name_group)
													tmp.SetType("program")

													group = &tmp
													group_index[group.Name] = &tmp

												}

												if user != nil && group != nil {

													if user.Name != "" && group.Name == "" && user.ID == group.ID {
														group.SetName(user.Name)
													}

													// Can't use AddGroup() for incomplete group data
													user.Groups = append(user.Groups, *group)

												}

											}

										} else if entry[0] == "r" {

											// Do Nothing

										}

									}

								}

							}

						}

					}

				}

			}

		}

		names := make([]string, 0)

		for name := range user_index {
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
