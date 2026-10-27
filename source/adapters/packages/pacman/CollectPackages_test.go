package pacman

import "tholian-core/structs"
import "slices"
import "testing"

func TestCollectPackages(t *testing.T) {

	t.Run("CollectPackages(OPTIMIZED=true)", func(t *testing.T) {

		OPTIMIZED = true
		packages := CollectPackages()

		if slices.ContainsFunc[[]structs.Package, structs.Package](packages, func(pkg structs.Package) bool {
			return pkg.Name == "glibc" && pkg.Version.IsValid()
		}) == false {
			t.Errorf("Expected %v to contain %s", packages, "Name=glibc")
		}

	})

	t.Run("CollectPackages(OPTIMIZED=false)", func(t *testing.T) {

		OPTIMIZED = false
		packages := CollectPackages()

		if slices.ContainsFunc[[]structs.Package, structs.Package](packages, func(pkg structs.Package) bool {
			return pkg.Name == "glibc" && pkg.Version.IsValid()
		}) == false {
			t.Errorf("Expected %v to contain %s", packages, "Name=glibc")
		}

	})

}
