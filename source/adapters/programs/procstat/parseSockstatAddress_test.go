package procstat

import "testing"

func TestParseSockstatAddress(t *testing.T) {

	host, port := parseSockstatAddress("*:*")

	if host != "any" || port != 0 {
		t.Fatalf("Expected any:0, got %s:%d", host, port)
	}

	host, port = parseSockstatAddress("*:8080")

	if host != "any" || port != 8080 {
		t.Fatalf("Expected any:8080, got %s:%d", host, port)
	}

	host, port = parseSockstatAddress("127.0.0.1:22")

	if host != "127.0.0.1" || port != 22 {
		t.Fatalf("Expected 127.0.0.1:22, got %s:%d", host, port)
	}

	host, port = parseSockstatAddress("[::1]:443")

	if host != "::1" || port != 443 {
		t.Fatalf("Expected ::1:443, got %s:%d", host, port)
	}

	host, port = parseSockstatAddress("")

	if host != "" || port != 0 {
		t.Fatalf("Expected empty host and port, got %s:%d", host, port)
	}

}
