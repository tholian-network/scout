package syslog

import "testing"
import "time"

func TestParseIncidentsBusyBoxSSHD(t *testing.T) {

	now := time.Now()
	line := now.Format("Jan 02 15:04:05") + " host sshd[123]: Failed password for root from 203.0.113.7 port 22 ssh2\n"
	incidents := ParseIncidents([]byte(line), "test")

	if len(incidents) != 1 {
		t.Fatalf("Expected 1 incident, got %d", len(incidents))
	}

	if incidents[0].Type != "Recon" {
		t.Fatalf("Expected Recon incident, got %s", incidents[0].Type)
	}

	if len(incidents[0].Networks) == 0 {
		t.Fatal("Expected a network")
	}

	if len(incidents[0].Users) == 0 || incidents[0].Users[0].Name != "root" {
		t.Fatalf("Expected root user, got %#v", incidents[0].Users)
	}

}

func TestParseIncidentsOpenWrtLogreadDropbear(t *testing.T) {

	now := time.Now()
	line := now.Format("Mon Jan 02 15:04:05 2006") + " user.info dropbear[123]: Bad password attempt for 'root' from 203.0.113.8:2222\n"
	incidents := ParseIncidents([]byte(line), "test")

	if len(incidents) != 1 {
		t.Fatalf("Expected 1 incident, got %d", len(incidents))
	}

	if incidents[0].Type != "Recon" {
		t.Fatalf("Expected Recon incident, got %s", incidents[0].Type)
	}

	if len(incidents[0].Users) == 0 || incidents[0].Users[0].Name != "root" {
		t.Fatalf("Expected root user, got %#v", incidents[0].Users)
	}

}

func TestParseIncidentsISOAccess(t *testing.T) {

	line := time.Now().UTC().Format("2006-01-02T15:04:05Z") + " host sshd[1]: Accepted password for root from 203.0.113.9 port 22 ssh2\n"
	incidents := ParseIncidents([]byte(line), "test")

	if len(incidents) != 1 {
		t.Fatalf("Expected 1 incident, got %d", len(incidents))
	}

	if incidents[0].Type != "Exfil" {
		t.Fatalf("Expected Exfil incident, got %s", incidents[0].Type)
	}

}

func TestParseIncidentsIgnoresUnknownLines(t *testing.T) {

	lines := "not a syslog line\n" +
		"Oct 29 09:21:00 host unknown[1]: nothing to classify\n"

	incidents := ParseIncidents([]byte(lines), "test")

	if len(incidents) != 0 {
		t.Fatalf("Expected 0 incidents, got %d", len(incidents))
	}

}
