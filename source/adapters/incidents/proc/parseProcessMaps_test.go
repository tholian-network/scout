package proc

import "testing"

func TestParseProcessMaps(t *testing.T) {

	buffer := []byte(
		"00400000-00401000 r-xp 00000000 00:00 0 /usr/bin/good\n" +
			"7f000000-7f001000 r-xp 00000000 00:00 0 /tmp/evil.so\n" +
			"7f001000-7f002000 r-xp 00000000 00:00 0 /usr/lib/libfoo.so (deleted)\n" +
			"7f002000-7f003000 r-xp 00000000 00:00 0 /memfd:evil (deleted)\n" +
			"7f003000-7f004000 rw-p 00000000 00:00 0 [heap]\n")

	entries := parseProcessMaps(buffer)

	if len(entries) != 5 {
		t.Fatalf("Expected 5 entries, got %d", len(entries))
	}

	if entries[1].Path != "/tmp/evil.so" {
		t.Fatalf("Expected /tmp/evil.so, got %q", entries[1].Path)
	}

	if entries[1].Permissions != "r-xp" {
		t.Fatalf("Expected r-xp, got %q", entries[1].Permissions)
	}

}

func TestParseProcessMapsWithoutPath(t *testing.T) {

	buffer := []byte("7f004000-7f005000 rw-p 00000000 00:00 0\n")

	entries := parseProcessMaps(buffer)

	if len(entries) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries))
	}

	if entries[0].Start != "7f004000" || entries[0].End != "7f005000" {
		t.Fatalf("Expected 7f004000-7f005000, got %q-%q", entries[0].Start, entries[0].End)
	}

	if entries[0].Inode != "0" {
		t.Fatalf("Expected inode 0, got %q", entries[0].Inode)
	}

	if entries[0].Path != "" {
		t.Fatalf("Expected empty path, got %q", entries[0].Path)
	}

}

func TestParseStartTicks(t *testing.T) {

	buffer := []byte("123 (evil process) S 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 999")
	ticks := getProcessTicks(buffer)

	if ticks != 999 {
		t.Fatalf("Expected 999 ticks, got %d", ticks)
	}

}
