package alpine

import "testing"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectDevices(t *testing.T) {

	console, system := harness.NewSystem(t)

	if len(actions.CollectDevices(console, system)) == 0 {
		t.Skip("No PCI or USB devices are exposed by the container runtime")
	}

	for _, device := range system.Devices {
		if device.IsValid() == false {
			t.Fatalf("Expected device %s to be valid", device.Name)
		}
	}

}
