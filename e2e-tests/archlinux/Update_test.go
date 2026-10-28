package archlinux

import "testing"
import "tholian-core/types"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestUpdate(t *testing.T) {

	config := harness.FromEnv(t)
	harness.Reset(t, config)

	console, system := harness.NewSystem(t)

	if actions.CollectUpdates(console, system) == false {
		t.Fatal("Expected CollectUpdates() to succeed")
	}

	if _, ok := harness.FindUpdate(system.Updates, config.Package); ok == false {
		t.Fatalf("Expected updates to contain %s before Update()", config.Package)
	}

	if actions.Update(console, system) == false {
		t.Fatal("Expected Update() to succeed")
	}

	if _, ok := harness.FindUpdate(system.Updates, config.Package); ok == true {
		t.Fatalf("Expected %s to be updated", config.Package)
	}

	if actions.CollectPackages(console, system) == false {
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
