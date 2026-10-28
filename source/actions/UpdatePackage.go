package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-scout/adapters/packages/apk"
import "tholian-scout/adapters/packages/apt"
import "tholian-scout/adapters/packages/dnf"
import "tholian-scout/adapters/packages/pacman"
import "tholian-scout/adapters/packages/tdnf"
import "tholian-scout/adapters/packages/zypper"
import "strconv"

func UpdatePackage(console *structs.Console, system *structs.System, search matchers.Package) bool {

	var result bool

	console.Group("actions/UpdatePackage")

	updates := make([]structs.Update, 0)

	for u := 0; u < len(system.Updates); u++ {

		update := system.Updates[u]

		if search.Matches(update.Name, update.Version.String(), update.Architecture.String(), update.Manager.String(), "any") {
			updates = append(updates, update)
		}

	}

	if len(updates) > 0 {

		var updated []structs.Update
		var err error = nil

		if pacman.SUPPORTED == true {

			updated, err = pacman.Update(updates)

		} else if apt.SUPPORTED == true {

			updated, err = apt.Update(updates)

		} else if apk.SUPPORTED == true {

			updated, err = apk.Update(updates)

		} else if dnf.SUPPORTED == true {

			updated, err = dnf.Update(updates)

		} else if tdnf.SUPPORTED == true {

			updated, err = tdnf.Update(updates)

		} else if zypper.SUPPORTED == true {

			updated, err = zypper.Update(updates)

		} else {

			// freebsd

			// TODO: pkg.Update()

		}

		if err == nil && len(updated) > 0 {

			remaining := make([]structs.Update, 0)

			for u1 := 0; u1 < len(updates); u1++ {

				update := updates[u1]
				found := false

				for u2 := 0; u2 < len(updated); u2++ {

					if updated[u2].Name == update.Name {
						found = true
						break
					}

				}

				if found == false {
					remaining = append(remaining, update)
				}

			}

			system.SetUpdates(remaining)
			console.Log("Updated " + strconv.Itoa(len(updated)) + "/" + strconv.Itoa(len(updates)) + " Packages")
			result = true

		} else if err != nil {

			console.Error("Update(): Failed")
			console.Error(err.Error())

		}

	}

	console.GroupEnd("actions/UpdatePackage")

	return result

}
