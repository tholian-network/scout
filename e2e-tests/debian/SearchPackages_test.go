package debian

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchPackages(t *testing.T) {

	config := harness.FromEnv(t)
	console, system := harness.NewSystem(t)

	if actions.CollectPackages(console, system) == false {
		t.Fatal("Expected CollectPackages() to succeed")
	}

	results := actions.SearchPackages(console, system, matchers.ToPackage(config.Package))

	if len(results) == 0 {
		t.Fatalf("Expected SearchPackages(%s) to find a package", config.Package)
	}

}
