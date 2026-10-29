package freebsd

import "testing"
import "strings"
import "tholian-scout/actions"
import "tholian-scout-e2e/internal/harness"

func TestCollectIncidentsProc(t *testing.T) {

	console, system := harness.NewSystem(t)

	actions.CollectIncidents(console, system)

	found := false

	for _, incident := range system.Incidents {

		if incident.Type != "Destroy" {
			continue
		}

		for _, program := range incident.Programs {
			if program.Command == "/tmp/e2e-sleep" {
				found = true
			}
		}

	}

	if found == false {
		t.Fatal("Expected a Destroy incident for /tmp/e2e-sleep")
	}

}

func TestCollectIncidentsSyslog(t *testing.T) {

	console, system := harness.NewSystem(t)

	actions.CollectIncidents(console, system)

	found := false

	for _, incident := range system.Incidents {

		if incident.Type != "Recon" {
			continue
		}

		for _, network := range incident.Networks {
			if strings.HasPrefix(network.Subnet, "203.0.113.7") {
				found = true
			}
		}

	}

	if found == false {
		t.Fatal("Expected a Recon incident for 203.0.113.7")
	}

}
