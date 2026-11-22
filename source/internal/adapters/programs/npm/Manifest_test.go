package npm

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

	manifest := NewManifest(fixture(t, "project") + "/package.json")

	if manifest.Name != "example-service" {
		t.Fatalf("Expected name example-service, got %s", manifest.Name)
	}

	if manifest.Version != "1.2.3" {
		t.Fatalf("Expected version 1.2.3, got %s", manifest.Version)
	}

	if manifest.Private != false {
		t.Fatal("Expected public package")
	}

	if manifest.Homepage != "https://example.com/service" {
		t.Fatalf("Expected homepage, got %s", manifest.Homepage)
	}

	if len(manifest.Dependencies) != 2 {
		t.Fatalf("Expected 2 dependencies, got %d", len(manifest.Dependencies))
	}

	if len(manifest.DevDependencies) != 1 {
		t.Fatalf("Expected 1 dev dependency, got %d", len(manifest.DevDependencies))
	}

	if len(manifest.Maintainers) != 4 {
		t.Fatalf("Expected 4 maintainers, got %d", len(manifest.Maintainers))
	}

	pkg := manifest.ToPackage()

	if pkg.Name != "example-service" {
		t.Fatalf("Expected package name example-service, got %s", pkg.Name)
	}

	if pkg.Version.Equal(types.ToVersion("1.2.3")) == false {
		t.Fatalf("Expected package version 1.2.3, got %s", pkg.Version.String())
	}

	if pkg.URL != "https://example.com/service" {
		t.Fatalf("Expected package URL from homepage, got %s", pkg.URL)
	}

	if len(pkg.Dependencies) != 3 {
		t.Fatalf("Expected 3 package dependencies, got %d", len(pkg.Dependencies))
	}

	if pkg.Dependencies[0].Name != "@scope/pkg" ||
		pkg.Dependencies[1].Name != "left-pad" ||
		pkg.Dependencies[2].Name != "typescript" {
		t.Fatalf("Expected sorted dependencies, got %v", pkg.Dependencies)
	}

	if pkg.IsValid() == false {
		t.Fatal("Expected parsed package to be valid")
	}

}

func TestManifestHostileTypes(t *testing.T) {

	manifest := NewManifest(fixture(t, "hostile") + "/package.json")

	if manifest.Name != "123" {
		t.Fatalf("Expected numeric name to be converted, got %s", manifest.Name)
	}

	if manifest.Version != "1.2" {
		t.Fatalf("Expected numeric version to be converted, got %s", manifest.Version)
	}

	if manifest.Private == false {
		t.Fatal("Expected string private flag to be true")
	}

	if manifest.Homepage != "https://example.com/hostile" {
		t.Fatalf("Expected homepage, got %s", manifest.Homepage)
	}

	if len(manifest.Maintainers) != 4 {
		t.Fatalf("Expected 4 valid maintainers, got %d", len(manifest.Maintainers))
	}

	if len(manifest.Dependencies) != 5 {
		t.Fatalf("Expected 5 valid dependencies, got %d", len(manifest.Dependencies))
	}

	if manifest.Dependencies["string-value"] != "^1.0.0" {
		t.Fatalf("Expected string dependency, got %s", manifest.Dependencies["string-value"])
	}

	if manifest.Dependencies["number-value"] != "2" {
		t.Fatalf("Expected numeric dependency, got %s", manifest.Dependencies["number-value"])
	}

	if manifest.Dependencies["null-value"] != "*" {
		t.Fatalf("Expected null dependency version *, got %s", manifest.Dependencies["null-value"])
	}

	if manifest.Dependencies["object-value"] != "3.0.0" {
		t.Fatalf("Expected object dependency version, got %s", manifest.Dependencies["object-value"])
	}

	if manifest.Dependencies["empty-value"] != "*" {
		t.Fatalf("Expected empty dependency version *, got %s", manifest.Dependencies["empty-value"])
	}

	if _, ok := manifest.Dependencies[""]; ok == true {
		t.Fatal("Expected empty dependency name to be skipped")
	}

	pkg := manifest.ToPackage()

	if pkg.IsValid() == false {
		t.Fatal("Expected hostile manifest package to remain valid")
	}

}

func TestManifestInvalidJSON(t *testing.T) {

	manifest := NewManifest(fixture(t, "invalid") + "/package.json")

	if manifest.Name != "" {
		t.Fatalf("Expected empty name, got %s", manifest.Name)
	}

	if len(manifest.Dependencies) != 0 {
		t.Fatalf("Expected no dependencies, got %d", len(manifest.Dependencies))
	}

}
