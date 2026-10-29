package freebsd

import "testing"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectAntiques(t *testing.T) {

	config := harness.FromEnv(t)
	console, system := harness.NewSystem(t)

	if actions.CollectServices(console, system) == false {
		t.Fatal("Expected CollectServices() to succeed")
	}

	if actions.CollectPackages(console, system) == false {
		t.Fatal("Expected CollectPackages() to succeed")
	}

	if actions.CollectUpdates(console, system) == false {
		t.Fatal("Expected CollectUpdates() to succeed")
	}

	if actions.CollectAntiques(console, system) == false {
		t.Fatal("Expected CollectAntiques() to succeed")
	}

	antique, ok := harness.FindAntique(system.Antiques, config.Package)

	if ok == false {
		t.Fatalf("Expected an antique for %s, got %v", config.Package, system.Antiques)
	}

	if antique.IsValid() == false {
		t.Fatalf("Expected antique for %s to be valid", config.Package)
	}

	if antique.Service == "" {
		t.Fatalf("Expected antique for %s to reference a service", config.Package)
	}

}
