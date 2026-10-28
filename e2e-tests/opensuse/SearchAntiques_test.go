package opensuse

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchAntiques(t *testing.T) {

	config := harness.FromEnv(t)
	console, system := harness.NewSystem(t)

	actions.CollectServices(console, system)
	actions.CollectPackages(console, system)
	actions.CollectUpdates(console, system)

	if actions.CollectAntiques(console, system) == false {
		t.Fatal("Expected CollectAntiques() to succeed")
	}

	results := actions.SearchAntiques(console, system, matchers.ToAntique(config.Package))

	if len(results) == 0 {
		t.Fatalf("Expected SearchAntiques(%s) to find an antique", config.Package)
	}

}
