package cblmariner

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchNetworks(t *testing.T) {

	console, system := harness.NewSystem(t)

	if actions.CollectNetworks(console, system) == false {
		t.Fatal("Expected CollectNetworks() to succeed")
	}

	if len(system.Networks) == 0 || len(system.Networks[0].Subnets) == 0 {
		t.Fatal("Expected at least one network with a subnet")
	}

	search := matchers.ToNetwork(system.Networks[0].Subnets[0].Address)
	results := actions.SearchNetworks(console, system, search)

	if len(results) == 0 {
		t.Fatalf("Expected SearchNetworks(%s) to find a network", system.Networks[0].Subnets[0].Address)
	}

}
