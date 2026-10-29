package pkg

import "strings"
import "testing"

func TestParseInstalledPackages(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"e2e-scout-pkg\t1.0",
		"pkg\t1.21.3",
		"curl\t8.14.1_1",
	}, "\r\n"))

	installed := parseInstalledPackages(string(buffer))

	if len(installed) != 3 {
		t.Fatalf("Expected 3 installed packages, got %d", len(installed))
	}

	e2e, ok := installed["e2e-scout-pkg"]

	if ok == false {
		t.Fatal("Expected e2e-scout-pkg to be installed")
	}

	if e2e.Upstream.Major != 1 || e2e.Upstream.Minor != 0 {
		t.Fatalf("Expected version 1.0, got %s", e2e.String())
	}

	pkg, ok := installed["pkg"]

	if ok == false {
		t.Fatal("Expected pkg to be installed")
	}

	if pkg.Upstream.Major != 1 || pkg.Upstream.Minor != 21 || pkg.Upstream.Patch != 3 {
		t.Fatalf("Expected version 1.21.3, got %s", pkg.String())
	}

	curl, ok := installed["curl"]

	if ok == false {
		t.Fatal("Expected curl to be installed")
	}

	if curl.Upstream.Major != 8 || curl.Upstream.Minor != 14 || curl.Upstream.Patch != 1 {
		t.Fatalf("Expected version 8.14.1, got %s", curl.String())
	}

	if curl.Revision != "1" {
		t.Fatalf("Expected revision 1, got %s", curl.Revision)
	}

}

func TestParseInstalledPackagesInvalid(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"invalid",
		"broken\tnot-a-version",
		"curl\t8.14.1_1",
	}, "\r\n"))

	installed := parseInstalledPackages(string(buffer))

	if len(installed) != 1 {
		t.Fatalf("Expected 1 installed package, got %d", len(installed))
	}

	if _, ok := installed["curl"]; ok == false {
		t.Fatal("Expected curl to be installed")
	}

}
