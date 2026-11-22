package openwrt

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectIncidentsSince(t *testing.T) {

	console, system := harness.NewSystem(t)
	search := matchers.ToDatetime("> 2000-01-01")

	if len(actions.CollectIncidentsSince(console, system, search)) == 0 {
		t.Fatal("Expected CollectIncidentsSince() to succeed")
	}

	for _, incident := range system.Incidents {
		if incident.IsValid() == false {
			t.Fatalf("Expected incident %s to be valid", incident.Type)
		}
	}

}
