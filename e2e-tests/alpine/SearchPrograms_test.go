package alpine

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchPrograms(t *testing.T) {

	console, system := harness.NewSystem(t)

	if actions.CollectPrograms(console, system) == false {
		t.Fatal("Expected CollectPrograms() to succeed")
	}

	if len(system.Programs) == 0 {
		t.Fatal("Expected CollectPrograms() to collect at least one program")
	}

	search := matchers.ToProgram(system.Programs[0].Name)
	results := actions.SearchPrograms(console, system, search)

	if len(results) == 0 {
		t.Fatalf("Expected SearchPrograms(%s) to find a program", system.Programs[0].Name)
	}

}
