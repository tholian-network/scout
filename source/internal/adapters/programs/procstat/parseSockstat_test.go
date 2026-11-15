package procstat

import "strings"
import "testing"

func TestParseSockstat(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"USER     COMMAND         PID FD PROTO LOCAL ADDRESS FOREIGN ADDRESS",
		"root     init              1  4 tcp4  *:22          *:*",
		"root     e2e-scout-pkg  1234  3 tcp46 *:8080        *:*",
	}, "\r\n"))

	sockets := parseSockstat(string(buffer))

	if len(sockets) != 2 {
		t.Fatalf("Expected 2 sockets, got %d", len(sockets))
	}

	if sockets[0].PID != 1 || sockets[0].Protocol != "tcp" {
		t.Fatalf("Unexpected socket: %v", sockets[0])
	}

	if sockets[0].Local != "any" || sockets[0].LocalPort != 22 {
		t.Fatalf("Unexpected local address: %s:%d", sockets[0].Local, sockets[0].LocalPort)
	}

	if sockets[0].Foreign != "any" || sockets[0].ForeignPort != 0 {
		t.Fatalf("Unexpected foreign address: %s:%d", sockets[0].Foreign, sockets[0].ForeignPort)
	}

	if sockets[1].PID != 1234 || sockets[1].Protocol != "tcp" {
		t.Fatalf("Unexpected socket: %v", sockets[1])
	}

	if sockets[1].Local != "any" || sockets[1].LocalPort != 8080 {
		t.Fatalf("Unexpected local address: %s:%d", sockets[1].Local, sockets[1].LocalPort)
	}

}
