package opkg

import "testing"
import "tholian-core/types"

func TestParseUpgradable(t *testing.T) {

	buffer := readFixture(t, "status")
	architectures := collectArchitectures(buffer)
	updates := parseUpgradable(readFixture(t, "list-upgradable"), architectures)

	if len(updates) != 2 {
		t.Fatalf("Expected 2 updates, got %d", len(updates))
	}

	var uhttpd_update *types.Version = nil
	var uhttpd_architecture string = ""

	for u := 0; u < len(updates); u++ {

		if updates[u].Name == "uhttpd" {
			uhttpd_update = &updates[u].Version
			uhttpd_architecture = updates[u].Architecture.String()
		}

	}

	if uhttpd_update == nil {
		t.Fatal("Expected uhttpd update")
	}

	if uhttpd_update.IsAfter(types.ToVersion("2023-06-25-34a8a74d-1")) == false {
		t.Fatalf("Expected uhttpd update to be newer than the old version, got %s", uhttpd_update.String())
	}

	if uhttpd_architecture != "x86_64" {
		t.Fatalf("Expected uhttpd update architecture x86_64, got %s", uhttpd_architecture)
	}

}
