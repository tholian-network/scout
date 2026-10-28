package almalinux

import "testing"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectDrives(t *testing.T) {

	console, system := harness.NewSystem(t)

	if actions.CollectDrives(console, system) == false {
		t.Fatal("Expected CollectDrives() to succeed")
	}

	found := false

	for _, drive := range system.Drives {
		if drive.Mountpoint == "/" {
			found = true
			break
		}
	}

	if found == false {
		t.Fatalf("Expected drives to contain the root filesystem, got %v", system.Drives)
	}

}
