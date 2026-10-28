package fedora

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchIncidents(t *testing.T) {

	console, system := harness.NewSystem(t)

	if actions.CollectIncidents(console, system) == false {
		t.Fatal("Expected CollectIncidents() to succeed")
	}

	results := actions.SearchIncidents(console, system, matchers.ToIncident(system.Name))

	if len(results) != len(system.Incidents) {
		t.Fatalf("Expected SearchIncidents(%s) to return %d incidents, got %d", system.Name, len(system.Incidents), len(results))
	}

}
