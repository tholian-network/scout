package actions

import "tholian-core/types"
import "tholian-core/structs"
import "tholian-scout/adapters/users/shadow"
import "tholian-scout/adapters/users/systemd"
import "strconv"

func CollectUsers(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectUsers")

	collected := make([]types.User, 0)

	if shadow.SUPPORTED == true {

		tmp := shadow.CollectUsers()

		if len(tmp) > 0 {

			console.Info("shadow.CollectUsers(): Found " + strconv.Itoa(len(tmp)) + " Users")

			for t := 0; t < len(tmp); t++ {
				collected = append(collected, tmp[t])
			}

		} else {

			console.Warn("shadow.CollectUsers(): Found 0 Users")

		}

	} else {

		console.Warn("shadow.CollectUsers(): Unsupported")

	}

	if systemd.SUPPORTED == true {

		tmp := systemd.CollectUsers()

		if len(tmp) > 0 {

			console.Info("systemd.CollectUsers(): Found " + strconv.Itoa(len(tmp)) + " Users")

			for t := 0; t < len(tmp); t++ {

				var sysuser = tmp[t]
				var user *types.User

				for c := 0; c < len(collected); c++ {

					if collected[c].Name == sysuser.Name {
						user = &collected[c]
						break
					}

				}

				if user != nil {

					if user.Type == "user" {
						user.SetType("program")
					}

					for g1 := 0; g1 < len(sysuser.Groups); g1++ {

						var has_group bool = false

						for g2 := 0; g2 < len(user.Groups); g2++ {

							if sysuser.Groups[g1].Name == user.Groups[g2].Name {
								has_group = true
								break
							}

						}

						if has_group == false {
							user.AddGroup(sysuser.Groups[g1])
						}

					}

				} else {

					console.Warn("Cannot find System User \"" + sysuser.Name + "\"")

				}

			}

		} else {

			console.Warn("systemd.CollectUsers(): Found 0 Users")

		}

	} else {

		console.Warn("systemd.CollectUsers(): Unsupported")

	}

	if len(collected) > 0 {

		system.SetUsers(collected)
		result = true

	}

	console.Log("Collected " + strconv.Itoa(len(system.Users)) + "/" + strconv.Itoa(len(collected)) + " Users")
	console.GroupEnd("actions/CollectUsers")

	return result

}
