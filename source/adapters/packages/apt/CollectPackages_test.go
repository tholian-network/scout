package apt

import "tholian-core/structs"
import "slices"
import "testing"

func TestCollectPackages(t *testing.T) {

	t.Run("CollectPackages()", func(t *testing.T) {

		packages := CollectPackages()

		if slices.ContainsFunc[[]structs.Package, structs.Package](packages, func(pkg structs.Package) bool {
			return pkg.Name == "libc6" && pkg.Version.IsValid()
		}) == false {
			t.Errorf("Expected %v to contain %s", packages, "Name=libc6")
		}

	})

}
