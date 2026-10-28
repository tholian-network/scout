package cblmariner

import "testing"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectServices(t *testing.T) {

	config := harness.FromEnv(t)
	console, system := harness.NewSystem(t)

	if actions.CollectServices(console, system) == false {
		t.Fatal("Expected CollectServices() to succeed")
	}

	service, ok := harness.FindService(system.Services, config.Service)

	if ok == false {
		t.Fatalf("Expected services to contain %s, got %v", config.Service, system.Services)
	}

	if service.IsService() == false {
		t.Fatalf("Expected %s to be classified as a service", config.Service)
	}

	found := false

	for _, connection := range service.Connections {
		if connection.Type == "server" {
			found = true
			break
		}
	}

	if found == false {
		t.Fatalf("Expected %s to have a listening server connection", config.Service)
	}

}
