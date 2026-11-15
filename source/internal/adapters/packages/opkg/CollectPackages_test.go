package opkg

import "testing"

func TestParseStatus(t *testing.T) {

	buffer := readFixture(t, "status")
	packages := parseStatus(buffer)

	if len(packages) != 4 {
		t.Fatalf("Expected 4 installed packages, got %d", len(packages))
	}

	if _, ok := findPackage(packages, "luci-i18n-base-de"); ok == true {
		t.Fatal("Expected config-files package to be skipped")
	}

	if _, ok := findPackage(packages, "half-installed"); ok == true {
		t.Fatal("Expected unpacked package to be skipped")
	}

	uhttpd, ok := findPackage(packages, "uhttpd")

	if ok == false {
		t.Fatal("Expected uhttpd package")
	}

	if hasDependency(uhttpd, "libc") == false || hasDependency(uhttpd, "libgcc1") == false {
		t.Fatal("Expected installed dependencies on uhttpd")
	}

	ca_bundle, ok := findPackage(packages, "ca-bundle")

	if ok == false {
		t.Fatal("Expected ca-bundle package")
	}

	if ca_bundle.Architecture.String() != "any" {
		t.Fatalf("Expected all architecture to normalize to any, got %s", ca_bundle.Architecture.String())
	}

}

func TestCollectArchitectures(t *testing.T) {

	buffer := readFixture(t, "status")
	architectures := collectArchitectures(buffer)

	if architectures["uhttpd"] != "x86_64" {
		t.Fatalf("Expected uhttpd architecture x86_64, got %s", architectures["uhttpd"])
	}

	if architectures["ca-bundle"] != "any" {
		t.Fatalf("Expected ca-bundle architecture any, got %s", architectures["ca-bundle"])
	}

	if _, ok := architectures["luci-i18n-base-de"]; ok == true {
		t.Fatal("Expected config-files package architecture to be skipped")
	}

}
