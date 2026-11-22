package fedora

import "testing"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectIncidents(t *testing.T) {

	console, system := harness.NewSystem(t)

	if len(actions.CollectIncidents(console, system)) == 0 {
		t.Fatal("Expected CollectIncidents() to succeed")
	}

	for _, incident := range system.Incidents {
		if incident.IsValid() == false {
			t.Fatalf("Expected incident %s to be valid", incident.Type)
		}
	}

}
