package archlinux

import "testing"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectNetworks(t *testing.T) {

	console, system := harness.NewSystem(t)

	if actions.CollectNetworks(console, system) == false {
		t.Fatal("Expected CollectNetworks() to succeed")
	}

	found := false

	for _, network := range system.Networks {
		if network.Name == "lo" {
			found = true
			break
		}
	}

	if found == false {
		t.Fatalf("Expected networks to contain loopback, got %v", system.Networks)
	}

}
