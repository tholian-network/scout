package alpine

import "testing"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectUsers(t *testing.T) {

	console, system := harness.NewSystem(t)

	if actions.CollectUsers(console, system) == false {
		t.Fatal("Expected CollectUsers() to succeed")
	}

	found := false

	for _, user := range system.Users {
		if user.Name == "root" {
			found = true
			break
		}
	}

	if found == false {
		t.Fatalf("Expected users to contain root, got %v", system.Users)
	}

}
