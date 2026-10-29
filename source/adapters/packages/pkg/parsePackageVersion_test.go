package pkg

import "testing"

func TestParsePackageVersion(t *testing.T) {

	version := parsePackageVersion("1.2.3_4,5")

	if version == nil {
		t.Fatal("Expected version")
	}

	if version.Epoche != 5 {
		t.Fatalf("Expected epoch 5, got %d", version.Epoche)
	}

	if version.Revision != "4" {
		t.Fatalf("Expected revision 4, got %s", version.Revision)
	}

	if version.Upstream.Major != 1 || version.Upstream.Minor != 2 || version.Upstream.Patch != 3 {
		t.Fatalf("Expected version 1.2.3, got %s", version.String())
	}

}

func TestParsePackageVersionRevision(t *testing.T) {

	version := parsePackageVersion("8.14.1_1")

	if version == nil {
		t.Fatal("Expected version")
	}

	if version.Epoche != 0 {
		t.Fatalf("Expected epoch 0, got %d", version.Epoche)
	}

	if version.Revision != "1" {
		t.Fatalf("Expected revision 1, got %s", version.Revision)
	}

	if version.Upstream.Major != 8 || version.Upstream.Minor != 14 || version.Upstream.Patch != 1 {
		t.Fatalf("Expected version 8.14.1, got %s", version.String())
	}

}

func TestParsePackageVersionInvalid(t *testing.T) {

	if version := parsePackageVersion(""); version != nil {
		t.Fatalf("Expected nil version for empty value, got %s", version.String())
	}

	if version := parsePackageVersion("invalid"); version != nil {
		t.Fatalf("Expected nil version for invalid value, got %s", version.String())
	}

}
