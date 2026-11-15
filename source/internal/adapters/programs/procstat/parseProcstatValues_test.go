package procstat

import "strings"
import "testing"

func TestParseProcstatValues(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"    1 init            /sbin/init --",
		" 1234 e2e-scout-pkg   /usr/local/bin/e2e-scout-pkg --listen :8080",
	}, "\r\n"))

	values := parseProcstatValues(string(buffer))

	if len(values) != 2 {
		t.Fatalf("Expected 2 values, got %d", len(values))
	}

	if values[1] != "/sbin/init --" {
		t.Fatalf("Expected /sbin/init --, got %q", values[1])
	}

	if values[1234] != "/usr/local/bin/e2e-scout-pkg --listen :8080" {
		t.Fatalf("Expected /usr/local/bin/e2e-scout-pkg --listen :8080, got %q", values[1234])
	}

}

func TestParseProcstatValuesEnvironment(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"    1 init            PATH=/sbin:/bin",
		" 1234 e2e-scout-pkg   PATH=/usr/local/bin:/usr/bin:/bin LD_PRELOAD=/tmp/evil.so",
	}, "\r\n"))

	values := parseProcstatValues(string(buffer))

	if len(values) != 2 {
		t.Fatalf("Expected 2 values, got %d", len(values))
	}

	if values[1] != "PATH=/sbin:/bin" {
		t.Fatalf("Expected PATH=/sbin:/bin, got %q", values[1])
	}

	if values[1234] != "PATH=/usr/local/bin:/usr/bin:/bin LD_PRELOAD=/tmp/evil.so" {
		t.Fatalf("Unexpected environment: %q", values[1234])
	}

}
