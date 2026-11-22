package composer

import "path/filepath"
import "testing"
import "tholian-core/types"

func fixture(t *testing.T, name string) string {

	t.Helper()

	path, err := filepath.Abs("testdata/" + name)

	if err != nil {
		t.Fatal(err)
	}

	return path

}

func TestManifestStandard(t *testing.T) {

	manifest := NewManifest(fixture(t, "project") + "/composer.json")

	if manifest.Name != "acme/example-service" {
		t.Fatalf("Expected name acme/example-service, got %s", manifest.Name)
	}

	if manifest.Version != "1.4.0" {
		t.Fatalf("Expected version 1.4.0, got %s", manifest.Version)
	}

	if manifest.Homepage != "https://example.com/php-service" {
		t.Fatalf("Expected homepage, got %s", manifest.Homepage)
	}

	if len(manifest.Authors) != 2 {
		t.Fatalf("Expected 2 authors, got %d", len(manifest.Authors))
	}

	pkg := manifest.ToPackage()

	if pkg.Name != "acme/example-service" {
		t.Fatalf("Expected package name, got %s", pkg.Name)
	}

	if pkg.Version.Equal(types.ToVersion("1.4.0")) == false {
		t.Fatalf("Expected version 1.4.0, got %s", pkg.Version.String())
	}

	if pkg.URL != "https://example.com/php-service" {
		t.Fatalf("Expected homepage URL, got %s", pkg.URL)
	}

	if len(pkg.Dependencies) != 3 {
		t.Fatalf("Expected 3 dependencies after platform filtering, got %d", len(pkg.Dependencies))
	}

	if pkg.Dependencies[0].Name != "monolog/monolog" ||
		pkg.Dependencies[1].Name != "symfony/console" ||
		pkg.Dependencies[2].Name != "phpunit/phpunit" {
		t.Fatalf("Expected deterministic dependency order, got %v", pkg.Dependencies)
	}

	if len(pkg.Conflicts) != 1 || pkg.Conflicts[0].Name != "acme/old-service" {
		t.Fatalf("Expected conflict, got %v", pkg.Conflicts)
	}

	if len(pkg.Replaces) != 1 || pkg.Replaces[0].Name != "acme/legacy" {
		t.Fatalf("Expected replace, got %v", pkg.Replaces)
	}

	if len(pkg.Provides) != 1 || pkg.Provides[0].Name != "psr/log-implementation" {
		t.Fatalf("Expected provide, got %v", pkg.Provides)
	}

	if pkg.IsValid() == false {
		t.Fatal("Expected parsed package to be valid")
	}

}

func TestManifestHostileTypes(t *testing.T) {

	manifest := NewManifest(fixture(t, "hostile") + "/composer.json")

	if manifest.Name != "acme/hostile" {
		t.Fatalf("Expected name acme/hostile, got %s", manifest.Name)
	}

	if len(manifest.Authors) != 1 {
		t.Fatalf("Expected string author to be parsed, got %d", len(manifest.Authors))
	}

	pkg := manifest.ToPackage()

	if len(pkg.Dependencies) != 5 {
		t.Fatalf("Expected 5 dependencies after filtering, got %d", len(pkg.Dependencies))
	}

	found := make(map[string]string)

	for d := 0; d < len(pkg.Dependencies); d++ {
		found[pkg.Dependencies[d].Name] = pkg.Dependencies[d].Version
	}

	if found["string-value"] != "^1.0" {
		t.Fatalf("Expected string dependency, got %s", found["string-value"])
	}

	if found["number-value"] != "2" {
		t.Fatalf("Expected numeric dependency, got %s", found["number-value"])
	}

	if found["null-value"] != "any" {
		t.Fatalf("Expected null dependency version any, got %s", found["null-value"])
	}

	if found["object-value"] != "3.0" {
		t.Fatalf("Expected object dependency version, got %s", found["object-value"])
	}

	if _, ok := found["php"]; ok == true {
		t.Fatal("Expected php platform dependency to be filtered")
	}

	if _, ok := found["ext-json"]; ok == true {
		t.Fatal("Expected ext-json platform dependency to be filtered")
	}

	if pkg.IsValid() == false {
		t.Fatal("Expected hostile package to remain valid")
	}

}
