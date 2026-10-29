package procstat

import "testing"

func TestParseEnvironment(t *testing.T) {

	environment := parseEnvironment("PATH=/sbin:/bin")

	if len(environment) != 1 {
		t.Fatalf("Expected 1 variable, got %d", len(environment))
	}

	if environment["PATH"] != "/sbin:/bin" {
		t.Fatalf("Expected PATH=/sbin:/bin, got %q", environment["PATH"])
	}

}

func TestParseEnvironmentMultiple(t *testing.T) {

	environment := parseEnvironment("PATH=/usr/local/bin:/usr/bin:/bin LD_PRELOAD=/tmp/evil.so")

	if len(environment) != 2 {
		t.Fatalf("Expected 2 variables, got %d", len(environment))
	}

	if environment["PATH"] != "/usr/local/bin:/usr/bin:/bin" {
		t.Fatalf("Unexpected PATH: %q", environment["PATH"])
	}

	if environment["LD_PRELOAD"] != "/tmp/evil.so" {
		t.Fatalf("Unexpected LD_PRELOAD: %q", environment["LD_PRELOAD"])
	}

}

func TestParseEnvironmentInvalid(t *testing.T) {

	environment := parseEnvironment("INVALID PATH=/sbin:/bin")

	if len(environment) != 1 {
		t.Fatalf("Expected 1 variable, got %d", len(environment))
	}

	if environment["PATH"] != "/sbin:/bin" {
		t.Fatalf("Expected PATH=/sbin:/bin, got %q", environment["PATH"])
	}

}
