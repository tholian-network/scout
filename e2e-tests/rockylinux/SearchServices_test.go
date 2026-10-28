package rockylinux

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchServices(t *testing.T) {

	config := harness.FromEnv(t)
	console, system := harness.NewSystem(t)

	if actions.CollectServices(console, system) == false {
		t.Fatal("Expected CollectServices() to succeed")
	}

	results := actions.SearchServices(console, system, matchers.ToProgram(config.Service))

	if len(results) == 0 {
		t.Fatalf("Expected SearchServices(%s) to find a service", config.Service)
	}

}
