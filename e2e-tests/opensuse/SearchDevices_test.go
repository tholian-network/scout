package opensuse

import "testing"
import "tholian-core/matchers"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestSearchDevices(t *testing.T) {

	console, system := harness.NewSystem(t)

	if actions.CollectDevices(console, system) == false {
		t.Skip("No PCI or USB devices are exposed by the container runtime")
	}

	search := matchers.ToDevice(system.Devices[0].Name)
	results := actions.SearchDevices(console, system, search)

	if len(results) == 0 {
		t.Fatalf("Expected SearchDevices(%s) to find a device", system.Devices[0].Name)
	}

}
