package dropbear

import "testing"

func TestClassifyBadPassword(t *testing.T) {

	incident := Classify("test", "2026-10-29 10:00:00", "host", "dropbear", "Bad password attempt for 'root' from [2001:db8::1]:2222")

	if incident.Type != "Recon" {
		t.Fatalf("Expected Recon incident, got %s", incident.Type)
	}

	if len(incident.Users) == 0 || incident.Users[0].Name != "root" {
		t.Fatalf("Expected root user, got %#v", incident.Users)
	}

}

func TestClassifyPasswordSuccess(t *testing.T) {

	incident := Classify("test", "2026-10-29 10:00:00", "host", "dropbear", "Password auth succeeded for 'root' from 203.0.113.7:2222")

	if incident.Type != "Exfil" {
		t.Fatalf("Expected Exfil incident, got %s", incident.Type)
	}

}

func TestParseHost(t *testing.T) {

	if value := parseHost("203.0.113.7:2222"); value != "203.0.113.7" {
		t.Fatalf("Unexpected IPv4 host: %s", value)
	}

	if value := parseHost("[2001:db8::1]:2222"); value != "2001:db8::1" {
		t.Fatalf("Unexpected IPv6 host: %s", value)
	}

}
