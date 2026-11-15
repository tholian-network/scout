package proc

import "testing"

func TestAnalyzeExecutableInWritablePath(t *testing.T) {

	program := programWithArguments("evil", "/tmp/evil", "/tmp/evil")
	incidents := AnalyzeExecutable("test", "2026-10-29 10:00:00", program)

	if len(incidents) == 0 || incidents[0].Type != "Destroy" {
		t.Fatalf("Expected Persist incident, got %#v", incidents)
	}

}
