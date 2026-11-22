package debian

import "testing"
import "tholian-core/matchers"
import "tholian-core/types"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestUpdatePackages(t *testing.T) {

	config := harness.FromEnv(t)
	harness.Reset(t, config)

	console, system := harness.NewSystem(t)

	if len(actions.CollectUpdates(console, system)) == 0 {
		t.Fatal("Expected CollectUpdates() to succeed")
	}

	searches := []matchers.Package{
		matchers.ToPackage(config.Package),
	}

	if len(actions.UpdatePackages(console, system, searches)) == 0 {
		t.Fatalf("Expected UpdatePackages(%s) to succeed", config.Package)
	}

	if _, ok := harness.FindUpdate(system.Updates, config.Package); ok == true {
		t.Fatalf("Expected %s to be updated", config.Package)
	}

	if len(actions.CollectPackages(console, system)) == 0 {
		t.Fatal("Expected CollectPackages() to succeed")
	}

	pkg, ok := harness.FindPackage(system.Packages, config.Package)

	if ok == false {
		t.Fatalf("Expected packages to contain %s", config.Package)
	}

	if pkg.Version.IsAfter(types.ToVersion(config.OldVersion)) == false {
		t.Fatalf("Expected %s to be newer than %s, got %s", config.Package, config.OldVersion, pkg.Version.String())
	}

}
