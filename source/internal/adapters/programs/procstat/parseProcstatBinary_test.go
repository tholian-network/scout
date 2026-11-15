package procstat

import "strings"
import "testing"

func TestParseProcstatBinary(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"    1 init            1401000 /sbin/init",
		" 1234 e2e-scout-pkg   1501000 /usr/local/bin/e2e-scout-pkg",
	}, "\r\n"))

	binary := parseProcstatBinary(string(buffer))

	if len(binary) != 2 {
		t.Fatalf("Expected 2 binaries, got %d", len(binary))
	}

	if binary[1] != "/sbin/init" {
		t.Fatalf("Expected /sbin/init, got %q", binary[1])
	}

	if binary[1234] != "/usr/local/bin/e2e-scout-pkg" {
		t.Fatalf("Expected /usr/local/bin/e2e-scout-pkg, got %q", binary[1234])
	}

}
