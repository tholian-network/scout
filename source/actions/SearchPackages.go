package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "strconv"

func SearchPackages(console *structs.Console, system *structs.System, search matchers.Package) []structs.Package {

	console.Group("actions/SearchPackages")

	var filtered []structs.Package

	for p := 0; p < len(system.Packages); p++ {

		pkg := system.Packages[p]

		if search.Matches(pkg.Name, pkg.Version.String(), pkg.Architecture.String(), pkg.Manager.String(), pkg.Vendor) {
			filtered = append(filtered, pkg)
		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Packages)) + " Packages")
	console.GroupEnd("actions/SearchPackages")

	return filtered

}
