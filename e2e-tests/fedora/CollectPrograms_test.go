package fedora

import "testing"
import "os"
import "path/filepath"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectPrograms(t *testing.T) {

	console, system := harness.NewSystem(t)

	if actions.CollectPrograms(console, system) == false {
		t.Fatal("Expected CollectPrograms() to succeed")
	}

	if len(system.Programs) == 0 {
		t.Fatal("Expected CollectPrograms() to collect at least one program")
	}

	name := filepath.Base(os.Args[0])

	if _, ok := harness.FindProgram(system.Programs, name); ok == false {
		t.Fatalf("Expected programs to contain the running test process %s", name)
	}

}
