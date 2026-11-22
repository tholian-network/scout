package fedora

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchUsers(t *testing.T) {

	console, system := harness.NewSystem(t)

	if len(actions.CollectUsers(console, system)) == 0 {
		t.Fatal("Expected CollectUsers() to succeed")
	}

	results := actions.SearchUsers(console, system, matchers.ToUser("root"))

	if len(results) == 0 {
		t.Fatal("Expected SearchUsers(root) to find the root user")
	}

}
