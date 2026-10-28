package alpine

import "testing"
import "tholian-core/types"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectUpdates(t *testing.T) {

	config := harness.FromEnv(t)
	console, system := harness.NewSystem(t)

	if actions.CollectUpdates(console, system) == false {
		t.Fatal("Expected CollectUpdates() to succeed")
	}

	update, ok := harness.FindUpdate(system.Updates, config.Package)

	if ok == false {
		t.Fatalf("Expected updates to contain %s, got %v", config.Package, system.Updates)
	}

	if update.Version.IsAfter(types.ToVersion(config.OldVersion)) == false {
		t.Fatalf("Expected %s update to be newer than %s, got %s", config.Package, config.OldVersion, update.Version.String())
	}

	if update.Manager.String() != config.PackageManager {
		t.Fatalf("Expected %s update manager %s, got %s", config.Package, config.PackageManager, update.Manager.String())
	}

}
