package procstat

import "strings"
import "testing"

func TestParseProcessStatus(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"    1     0 init",
		" 1234     0 e2e-scout-pkg",
	}, "\r\n"))

	processes := parseProcessStatus(string(buffer))

	if len(processes) != 2 {
		t.Fatalf("Expected 2 processes, got %d", len(processes))
	}

	if processes[0].PID != 1 || processes[0].UID != 0 || processes[0].Name != "init" {
		t.Fatalf("Unexpected process: %v", processes[0])
	}

	if processes[1].PID != 1234 || processes[1].UID != 0 || processes[1].Name != "e2e-scout-pkg" {
		t.Fatalf("Unexpected process: %v", processes[1])
	}

}
