package actions

import "tholian-core/structs"
import "tholian-scout/internal/adapters/packages/apk"
import "tholian-scout/internal/adapters/packages/apt"
import "tholian-scout/internal/adapters/packages/dnf"
import "tholian-scout/internal/adapters/packages/dpkg"
import "tholian-scout/internal/adapters/packages/opkg"
import "tholian-scout/internal/adapters/packages/pkg"
import "tholian-scout/internal/adapters/packages/pacman"
import "tholian-scout/internal/adapters/packages/rpm"
import "tholian-scout/internal/adapters/packages/tdnf"
import "tholian-scout/internal/adapters/packages/zypper"
import "strconv"

func CollectPackages(console *structs.Console, system *structs.System) []structs.Package {

	if system == nil {
		system = initSystem(console)
	}

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

	} else if dpkg.SUPPORTED == true {

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

	} else if rpm.SUPPORTED == true {

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

	} else if pkg.SUPPORTED == true {

		tmp := pkg.CollectPackages()

		if len(tmp) > 0 {

			console.Info("pkg.CollectPackages(): Found " + strconv.Itoa(len(tmp)) + " Packages")

			for _, pkg := range tmp {

				if pkg.IsValid() {
					collected = append(collected, pkg)
				}

			}

		} else {

			console.Warn("pkg.CollectPackages(): Found 0 Packages")

		}

		system.SetPackages(collected)

	} else if opkg.SUPPORTED == true {

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

	} else {

		console.Warn("CollectPackages(): Unsupported")

	}

	console.Log("Collected " + strconv.Itoa(len(system.Packages)) + "/" + strconv.Itoa(len(collected)) + " Packages")
	console.GroupEnd("actions/CollectPackages")

	return collected

}
