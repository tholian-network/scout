package proc

import "testing"
import "tholian-core/structs"
import "time"

func programWithEnvironment(key string, value string) structs.Program {

	program := structs.NewProgram(1234, "test")
	program.SetCommand("/usr/bin/test")
	program.SetArguments([]string{"test"})
	program.SetEnvironment(key, value)

	return program

}

func programWithArguments(name string, command string, arguments ...string) structs.Program {

	program := structs.NewProgram(1234, name)
	program.SetCommand(command)
	program.SetArguments(arguments)

	return program

}

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

func TestParseStartTicks(t *testing.T) {

	buffer := []byte("123 (evil process) S 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 999")
	ticks, ok := parseStartTicks(buffer)

	if ok == false || ticks != 999 {
		t.Fatalf("Expected 999 ticks, got %d, %v", ticks, ok)
	}

}

func TestToStartDatetime(t *testing.T) {

	datetime := toStartDatetime(1700000000, 12345)
	expected := time.Unix(1700000123, 0).Format("2006-01-02 15:04:05")

	if datetime.String() != expected {
		t.Fatalf("Unexpected datetime: %s, want %s", datetime.String(), expected)
	}

}
