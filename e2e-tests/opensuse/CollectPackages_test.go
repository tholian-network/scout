package opensuse

import "testing"
import "tholian-core/types"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectPackages(t *testing.T) {

	config := harness.FromEnv(t)
	console, system := harness.NewSystem(t)

	if actions.CollectPackages(console, system) == false {
		t.Fatal("Expected CollectPackages() to succeed")
	}

	pkg, ok := harness.FindPackage(system.Packages, config.Package)

	if ok == false {
		t.Fatalf("Expected packages to contain %s", config.Package)
	}

	if pkg.Version.Equal(types.ToVersion(config.OldVersion)) == false {
		t.Fatalf("Expected %s version %s, got %s", config.Package, config.OldVersion, pkg.Version.String())
	}

	if pkg.Manager.String() != config.PackageManager {
		t.Fatalf("Expected %s manager %s, got %s", config.Package, config.PackageManager, pkg.Manager.String())
	}

}
