package actions

import "tholian-core/structs"
import "tholian-scout/adapters/packages/apk"
import "tholian-scout/adapters/packages/apt"
import "tholian-scout/adapters/packages/dnf"
import "tholian-scout/adapters/packages/dpkg"
import "tholian-scout/adapters/packages/opkg"
import "tholian-scout/adapters/packages/pacman"
import "tholian-scout/adapters/packages/rpm"
import "tholian-scout/adapters/packages/tdnf"
import "tholian-scout/adapters/packages/zypper"
import "strconv"

func CollectPackages(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectPackages")

	collected := make([]structs.Package, 0)

	if pacman.SUPPORTED == true {

		tmp := pacman.CollectPackages()

		if len(tmp) > 0 {

			console.Info("pacman.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("pacman.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)
		result = true

	} else if apt.SUPPORTED == true {

		tmp := apt.CollectPackages()

		if len(tmp) > 0 {

			console.Info("apt.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("apt.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)
		result = true

	} else if apk.SUPPORTED == true {

		tmp := apk.CollectPackages()

		if len(tmp) > 0 {

			console.Info("apk.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("apk.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)
		result = true

	} else if dnf.SUPPORTED == true {

		tmp := dnf.CollectPackages()

		if len(tmp) > 0 {

			console.Info("dnf.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("dnf.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)
		result = true

	} else if tdnf.SUPPORTED == true {

		tmp := tdnf.CollectPackages()

		if len(tmp) > 0 {

			console.Info("tdnf.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("tdnf.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)
		result = true

	} else if zypper.SUPPORTED == true {

		tmp := zypper.CollectPackages()

		if len(tmp) > 0 {

			console.Info("zypper.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("zypper.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)
		result = true

	} else if dpkg.SUPPORTED == true {

		// dpkg is used as a fallback when no downstream
		// Debian family package manager is available

		tmp := dpkg.CollectPackages()

		if len(tmp) > 0 {

			console.Info("dpkg.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("dpkg.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)
		result = true

	} else if rpm.SUPPORTED == true {

		// RPM can only CollectPackages() and is used as a fallback
		// when no downstream package manager is available

		tmp := rpm.CollectPackages()

		if len(tmp) > 0 {

			console.Info("rpm.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("rpm.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)
		result = true

	} else if opkg.SUPPORTED == true {

		// openwrt

		tmp := opkg.CollectPackages()

		if len(tmp) > 0 {

			console.Info("opkg.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("opkg.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)
		result = true

	} else {

		// TODO: pkg.CollectPackages()

		console.Warn("CollectPackages(): Unsupported")

	}

	console.Log("Collected " + strconv.Itoa(len(system.Packages)) + "/" + strconv.Itoa(len(collected)) + " Packages")
	console.GroupEnd("actions/CollectPackages")

	return result

}
