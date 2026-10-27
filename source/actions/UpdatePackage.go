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
		matches := false

		if search.Matches(update.Name, update.Version.String(), update.Architecture.String(), update.Manager.String(), "any") {
			matches = true
			break
		}

		if matches == true {
			updates = append(updates, update)
		}

	}

	if len(updates) > 0 {

		if pacman.SUPPORTED == true {

			updated, err := pacman.Update(updates)

			if err == nil {

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

				if len(remaining) > 0 {
					system.SetUpdates(remaining)
				} else {
					result = true
				}

				console.Log("Updated " + strconv.Itoa(len(updates)) + "/" + strconv.Itoa(len(system.Updates)) + " Packages")

			} else {

				console.Error("pacman.Update(): Failed")
				console.Error(err.Error())

			}

		} else if apt.SUPPORTED == true {

			// TODO: apt.Update()

		} else if apk.SUPPORTED == true {

			// TODO: apk.Update()

		} else if dnf.SUPPORTED == true {

			// TODO: dnf.Update()

		} else if tdnf.SUPPORTED == true {

			// TODO: tdnf.Update()

		} else if zypper.SUPPORTED == true {

			// TODO: zypper.Update()

		} else {

			// freebsd

			// TODO: pkg.Update()

		}

	}

	console.GroupEnd("actions/UpdatePackage")

	return result

}
