package pkg

import "strings"
import "testing"

func TestParseRemotePackages(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"e2e-scout-pkg\t2.0\tFreeBSD:15:amd64",
		"pkg\t1.21.3\tFreeBSD:15:amd64",
		"curl\t8.15.0\tFreeBSD:15:amd64",
	}, "\r\n"))

	remote := parseRemotePackages(string(buffer))

	if len(remote) != 3 {
		t.Fatalf("Expected 3 remote packages, got %d", len(remote))
	}

	e2e, ok := remote["e2e-scout-pkg"]

	if ok == false {
		t.Fatal("Expected e2e-scout-pkg to be available")
	}

	if e2e.Version.Upstream.Major != 2 || e2e.Version.Upstream.Minor != 0 {
		t.Fatalf("Expected version 2.0, got %s", e2e.Version.String())
	}

	if e2e.Architecture != "FreeBSD:15:amd64" {
		t.Fatalf("Expected FreeBSD:15:amd64, got %s", e2e.Architecture)
	}

	pkg, ok := remote["pkg"]

	if ok == false {
		t.Fatal("Expected pkg to be available")
	}

	if pkg.Version.Upstream.Major != 1 || pkg.Version.Upstream.Minor != 21 || pkg.Version.Upstream.Patch != 3 {
		t.Fatalf("Expected version 1.21.3, got %s", pkg.Version.String())
	}

	curl, ok := remote["curl"]

	if ok == false {
		t.Fatal("Expected curl to be available")
	}

	if curl.Version.Upstream.Major != 8 || curl.Version.Upstream.Minor != 15 || curl.Version.Upstream.Patch != 0 {
		t.Fatalf("Expected version 8.15.0, got %s", curl.Version.String())
	}

}

func TestParseRemotePackagesDefaultArchitecture(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"curl\t8.15.0",
		"pkg\t1.21.3\t",
	}, "\r\n"))

	remote := parseRemotePackages(string(buffer))

	if len(remote) != 2 {
		t.Fatalf("Expected 2 remote packages, got %d", len(remote))
	}

	if remote["curl"].Architecture != "any" {
		t.Fatalf("Expected default architecture any, got %s", remote["curl"].Architecture)
	}

	if remote["pkg"].Architecture != "any" {
		t.Fatalf("Expected default architecture any, got %s", remote["pkg"].Architecture)
	}

}
