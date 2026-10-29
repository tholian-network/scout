package pkg

import "strings"
import "testing"

func TestParsePackages(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"e2e-scout-pkg\t1.0\tFreeBSD:15:amd64\tsysutils/e2e-scout-pkg\thttps://example.com/e2e-scout-pkg\te2e@example.com\t/usr/local\t1767225600\tFreeBSD",
		"pkg\t1.21.3\tFreeBSD:15:amd64\tports-mgmt/pkg\thttps://github.com/freebsd/pkg\tpkg@FreeBSD.org\t/usr/local\t1767225600\tFreeBSD",
		"curl\t8.14.1_1,1\tFreeBSD:15:amd64\tftp/curl\thttps://curl.se/\tcurl@example.com\t/usr/local\t1767225600\tFreeBSD",
	}, "\r\n"))

	packages := parsePackages(string(buffer))

	if len(packages) != 3 {
		t.Fatalf("Expected 3 packages, got %d", len(packages))
	}

	if packages[0].Name != "e2e-scout-pkg" {
		t.Fatalf("Expected e2e-scout-pkg first, got %s", packages[0].Name)
	}

	if packages[1].Name != "pkg" {
		t.Fatalf("Expected pkg second, got %s", packages[1].Name)
	}

	if packages[2].Name != "curl" {
		t.Fatalf("Expected curl third, got %s", packages[2].Name)
	}

	pkg := packages[2]

	if pkg.Version.Epoche != 1 {
		t.Fatalf("Expected port epoch 1, got %d", pkg.Version.Epoche)
	}

	if pkg.Version.Revision != "1" {
		t.Fatalf("Expected port revision 1, got %s", pkg.Version.Revision)
	}

	if pkg.Version.Upstream.Major != 8 || pkg.Version.Upstream.Minor != 14 || pkg.Version.Upstream.Patch != 1 {
		t.Fatalf("Expected version 8.14.1, got %s", pkg.Version.String())
	}

	if pkg.Architecture.String() != "x86_64" {
		t.Fatalf("Expected x86_64, got %s", pkg.Architecture.String())
	}

	if pkg.Manager.String() != "pkg" {
		t.Fatalf("Expected manager pkg, got %s", pkg.Manager.String())
	}

	if pkg.Vendor != "ftp/curl" {
		t.Fatalf("Expected origin vendor ftp/curl, got %s", pkg.Vendor)
	}

	if pkg.URL != "https://curl.se/" {
		t.Fatalf("Expected homepage, got %s", pkg.URL)
	}

	if len(pkg.Maintainers) != 1 || pkg.Maintainers[0].Name != "curl" || pkg.Maintainers[0].Email != "curl@example.com" {
		t.Fatalf("Expected curl maintainer, got %v", pkg.Maintainers)
	}

	if pkg.Datetime.IsValid() == false {
		t.Fatal("Expected install timestamp")
	}

	if pkg.IsValid() == false {
		t.Fatal("Expected package to be valid")
	}

}
