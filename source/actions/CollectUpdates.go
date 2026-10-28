package actions

import "tholian-core/structs"
import "tholian-scout/adapters/packages/apk"
import "tholian-scout/adapters/packages/apt"
import "tholian-scout/adapters/packages/dnf"
import "tholian-scout/adapters/packages/pacman"
import "tholian-scout/adapters/packages/tdnf"
import "tholian-scout/adapters/packages/zypper"
import "strconv"

func CollectUpdates(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectUpdates")

	collected := make([]structs.Update, 0)

	if pacman.SUPPORTED == true {

		// antergos
		// archlinux
		// manjaro

		tmp := pacman.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("pacman.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("pacman.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else if apt.SUPPORTED == true {

		// debian
		// ubuntu
		// linuxmint

		tmp := apt.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("apt.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("apt.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else if apk.SUPPORTED == true {

		// alpinelinux

		tmp := apk.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("apk.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("apk.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else if dnf.SUPPORTED == true {

		// redhat
		// centos
		// oraclelinux
		// almalinux
		// rockylinux
		// fedora
		// amazonlinux

		tmp := dnf.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("dnf.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("dnf.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else if tdnf.SUPPORTED == true {

		// cblmariner
		// photonos

		tmp := tdnf.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("tdnf.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("tdnf.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else if zypper.SUPPORTED == true {

		// opensuse
		// suse-desktop
		// suse-server

		tmp := zypper.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("zypper.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("zypper.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else {

		// TODO: pkg.CollectUpdates()

		console.Warn("CollectUpdates(): Unsupported")

	}

	console.Log("Collected " + strconv.Itoa(len(system.Updates)) + "/" + strconv.Itoa(len(collected)) + " Updates")
	console.GroupEnd("actions/CollectUpdates")

	return result

}
