package alpine

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchDrives(t *testing.T) {

	console, system := harness.NewSystem(t)

	if actions.CollectDrives(console, system) == false {
		t.Fatal("Expected CollectDrives() to succeed")
	}

	results := actions.SearchDrives(console, system, matchers.ToDrive("/"))

	if len(results) == 0 {
		t.Fatal("Expected SearchDrives(/) to find the root filesystem")
	}

}
