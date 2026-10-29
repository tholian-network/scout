package proc

import "testing"

func TestAnalyzeEnvironment(t *testing.T) {

	program := programWithEnvironment("LD_PRELOAD", "/usr/lib/libfoo.so /tmp/evil.so")
	incidents := AnalyzeEnvironment("test", "2026-10-29 10:00:00", program)

	if len(incidents) != 1 {
		t.Fatalf("Expected 1 incident, got %d", len(incidents))
	}

	if incidents[0].Type != "Destroy" {
		t.Fatalf("Expected Destroy incident, got %s", incidents[0].Type)
	}

}

func TestAnalyzeEnvironmentTrustedOnly(t *testing.T) {

	program := programWithEnvironment("LD_PRELOAD", "/usr/lib/libfoo.so")
	incidents := AnalyzeEnvironment("test", "2026-10-29 10:00:00", program)

	if len(incidents) != 0 {
		t.Fatalf("Expected 0 incidents, got %d", len(incidents))
	}

}
