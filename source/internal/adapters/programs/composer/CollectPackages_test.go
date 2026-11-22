package composer

import "testing"
import "tholian-core/types"

func TestCollectPackagesInstalled(t *testing.T) {

	packages := collectPackages(fixture(t, "project"))

	if len(packages) != 2 {
		t.Fatalf("Expected 2 installed packages, got %d", len(packages))
	}

	if packages[0].Name != "monolog/monolog" || packages[1].Name != "symfony/console" {
		t.Fatalf("Expected sorted packages, got %v", packages)
	}

	if packages[0].Version.Equal(types.ToVersion("3.5.0")) == false {
		t.Fatalf("Expected monolog 3.5.0, got %s", packages[0].Version.String())
	}

	if packages[1].Version.Equal(types.ToVersion("6.4.4")) == false {
		t.Fatalf("Expected symfony console 6.4.4, got %s", packages[1].Version.String())
	}

	if packages[0].URL != "https://github.com/Seldaek/monolog" {
		t.Fatalf("Expected monolog homepage, got %s", packages[0].URL)
	}

	if len(packages[0].Maintainers) != 1 || packages[0].Maintainers[0].Name != "Jordi Boggiano" {
		t.Fatalf("Expected monolog maintainer, got %v", packages[0].Maintainers)
	}

	if len(packages[1].Dependencies) != 1 || packages[1].Dependencies[0].Name != "symfony/string" {
		t.Fatalf("Expected symfony/string dependency, got %v", packages[1].Dependencies)
	}

	if packages[0].IsValid() == false {
		t.Fatal("Expected installed package to be valid")
	}

}

func TestCollectPackagesLock(t *testing.T) {

	packages := collectPackages(fixture(t, "lock"))

	if len(packages) != 2 {
		t.Fatalf("Expected 2 locked packages, got %d", len(packages))
	}

	if packages[0].Name != "acme/dev-tool" || packages[1].Name != "acme/locked" {
		t.Fatalf("Expected sorted lock packages, got %v", packages)
	}

	if len(packages[1].Dependencies) != 1 || packages[1].Dependencies[0].Name != "psr/log" {
		t.Fatalf("Expected psr/log dependency, got %v", packages[1].Dependencies)
	}

}

func TestCollectPackagesVendorScan(t *testing.T) {

	packages := collectPackages(fixture(t, "vendor"))

	if len(packages) != 1 {
		t.Fatalf("Expected 1 scanned package, got %d", len(packages))
	}

	if packages[0].Name != "acme/scanned" {
		t.Fatalf("Expected acme/scanned, got %s", packages[0].Name)
	}

	if packages[0].URL != "https://example.com/scanned" {
		t.Fatalf("Expected homepage URL, got %s", packages[0].URL)
	}

}
