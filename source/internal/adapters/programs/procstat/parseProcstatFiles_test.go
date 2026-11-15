package procstat

import "strings"
import "testing"

func TestParseProcstatFiles(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"    1 init           cwd  v d r-------   -       - -   /",
		"    1 init             0  v c rw------   7  372071 -   /dev/pts/0",
		" 1234 e2e-scout-pkg  cwd  v d r-------   -       - -   /",
		" 1234 e2e-scout-pkg    0  v c rw------   7  372071 -   /dev/pts/1",
		" 1234 e2e-scout-pkg    3  s - rw------   -       - -   tcp4 *:8080",
		" 1234 e2e-scout-pkg    4  v r r-------   -       - -   /usr/local/lib/libc.so.7",
	}, "\r\n"))

	files := parseProcstatFiles(string(buffer))

	if len(files) != 2 {
		t.Fatalf("Expected 2 files, got %d", len(files))
	}

	if files[1].Folder != "/" {
		t.Fatalf("Expected folder /, got %q", files[1].Folder)
	}

	if len(files[1].Files) != 1 || files[1].Files[0] != "/dev/pts/0" {
		t.Fatalf("Unexpected files: %v", files[1].Files)
	}

	if files[1234].Folder != "/" {
		t.Fatalf("Expected folder /, got %q", files[1234].Folder)
	}

	if len(files[1234].Files) != 2 {
		t.Fatalf("Expected 2 files, got %d", len(files[1234].Files))
	}

	if files[1234].Files[0] != "/dev/pts/1" {
		t.Fatalf("Expected /dev/pts/1, got %q", files[1234].Files[0])
	}

	if files[1234].Files[1] != "/usr/local/lib/libc.so.7" {
		t.Fatalf("Expected /usr/local/lib/libc.so.7, got %q", files[1234].Files[1])
	}

}
