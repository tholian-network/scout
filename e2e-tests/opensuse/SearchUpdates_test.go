package opensuse

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchUpdates(t *testing.T) {

	config := harness.FromEnv(t)
	console, system := harness.NewSystem(t)

	if actions.CollectUpdates(console, system) == false {
		t.Fatal("Expected CollectUpdates() to succeed")
	}

	results := actions.SearchUpdates(console, system, matchers.ToUpdate(config.Package))

	if len(results) == 0 {
		t.Fatalf("Expected SearchUpdates(%s) to find an update", config.Package)
	}

}
