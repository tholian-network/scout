package actions

import "tholian-insights/countries"
import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-scout/adapters/system/coreutils"
import "tholian-scout/adapters/system/etc"
import "tholian-scout/adapters/system/proc"
import "tholian-scout/adapters/system/sys"

func Init(console *structs.Console) *structs.System {

	console.Group("actions/Init")

	var system structs.System = structs.NewSystem()

	if etc.SUPPORTED == true {
		system = etc.CollectSystem()
		console.Log("System Name: \"" + system.Name + "\"")
	}

	if coreutils.SUPPORTED == true {

		kernel := coreutils.CollectKernel()
		kernel_architecture := coreutils.CollectKernelArchitecture()
		kernel_modules := proc.CollectKernelModules()
		kernel_version := coreutils.CollectKernelVersion()

		if kernel != "" {
			system.Distribution.SetKernel(kernel)
		}

		if kernel_architecture != "" {
			system.Distribution.SetKernelArchitecture(kernel_architecture)
		}

		if len(kernel_modules) > 0 {
			system.Distribution.SetKernelModules(kernel_modules)
		}

		if kernel_version != "" {
			system.Distribution.SetKernelVersion(kernel_version)
		}

	}

	if sys.SUPPORTED == true {
		system.SetBIOS(sys.CollectBIOS())
		system.SetBoard(sys.CollectBoard())
	}

	if system.Fingerprint.Timezone != "" {

		candidates := countries.Countries.Query(matchers.Country{
			Name: "any",
			Continent: "any",
			Allegiance: "any",
			Subnet: "any",
			Timezone: system.Fingerprint.Timezone,
		})

		if len(candidates) > 0 {
			system.SetCountry(candidates[0].ISO)
		} else {
			system.SetCountry("??")
		}

	}

	console.GroupEnd("actions/Init")

	return &system

}
