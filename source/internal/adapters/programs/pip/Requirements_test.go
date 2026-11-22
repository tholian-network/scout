package pip

import "os"
import "path/filepath"
import "testing"

func fixture(t *testing.T, name string) string {

	t.Helper()

	path, err := filepath.Abs("testdata/" + name)

	if err != nil {
		t.Fatal(err)
	}

	return path

}

func TestManifestRequirements(t *testing.T) {

	manifest := NewManifest(fixture(t, "project") + "/requirements.txt")

	expected := map[string]string{
		"requests":     "= 2.31.0",
		"Flask":        ">= 3.0, < 4.0",
		"uvicorn":      "~= 0.29.0",
		"django":       "= 5.0.2",
		"click":        "= 8.1.7",
		"editable-pkg": "*",
		"local-pkg":    "*",
	}

	if len(manifest.Dependencies) != len(expected) {
		t.Fatalf("Expected %d dependencies, got %d: %v", len(expected), len(manifest.Dependencies), manifest.Dependencies)
	}

	for name, version := range expected {

		if manifest.Dependencies[name] != version {
			t.Fatalf("Expected %s %s, got %s", name, version, manifest.Dependencies[name])
		}

	}

	if _, ok := manifest.Dependencies["urllib3"]; ok == true {
		t.Fatal("Expected constraint urllib3 to be skipped")
	}

	if _, ok := manifest.Dependencies["unnamed"]; ok == true {
		t.Fatal("Expected unnamed URL requirement to be skipped")
	}

	pkg := manifest.ToPackage()

	if pkg.Name != "project" {
		t.Fatalf("Expected project name, got %s", pkg.Name)
	}

	if len(pkg.Dependencies) != len(expected) {
		t.Fatalf("Expected %d package dependencies, got %d", len(expected), len(pkg.Dependencies))
	}

	if pkg.Dependencies[0].Name != "Flask" ||
		pkg.Dependencies[1].Name != "click" ||
		pkg.Dependencies[2].Name != "django" ||
		pkg.Dependencies[3].Name != "editable-pkg" ||
		pkg.Dependencies[4].Name != "local-pkg" ||
		pkg.Dependencies[5].Name != "requests" ||
		pkg.Dependencies[6].Name != "uvicorn" {
		t.Fatalf("Expected sorted package dependencies, got %v", pkg.Dependencies)
	}

	if pkg.Dependencies[0].Version != ">= 3.0" ||
		pkg.Dependencies[1].Version != "= 8.1.7" ||
		pkg.Dependencies[2].Version != "= 5.0.2" ||
		pkg.Dependencies[3].Version != "any" ||
		pkg.Dependencies[4].Version != "any" ||
		pkg.Dependencies[5].Version != "= 2.31.0" ||
		pkg.Dependencies[6].Version != ">= 0.29.0" {
		t.Fatalf("Expected normalized package dependency versions, got %v", pkg.Dependencies)
	}

}

func TestManifestHostile(t *testing.T) {

	manifest := NewManifest(fixture(t, "hostile") + "/requirements.txt")

	if manifest.Dependencies["Requests"] != "= 2.31.0" {
		t.Fatalf("Expected Requests dependency, got %s", manifest.Dependencies["Requests"])
	}

	if manifest.Dependencies["weird-name"] != "*" {
		t.Fatalf("Expected weird-name dependency, got %s", manifest.Dependencies["weird-name"])
	}

	if manifest.Dependencies["local-path"] != "*" {
		t.Fatalf("Expected local-path dependency, got %s", manifest.Dependencies["local-path"])
	}

	if manifest.Dependencies["foo"] != "*" {
		t.Fatalf("Expected marker dependency foo, got %s", manifest.Dependencies["foo"])
	}

	if manifest.Dependencies["baz"] != "= 1.0" {
		t.Fatalf("Expected trailing comment dependency baz, got %s", manifest.Dependencies["baz"])
	}

	if _, ok := manifest.Dependencies["bar"]; ok == true {
		t.Fatal("Expected extra marker dependency to be skipped")
	}

	if _, ok := manifest.Dependencies["."]; ok == true {
		t.Fatal("Expected editable local path to be skipped")
	}

}

func TestManifestIncludesCycle(t *testing.T) {

	root := t.TempDir()

	writeRequirement(t, root+"/requirements.txt", "-r requirements-base.txt\nroot-pkg==1.0\n")
	writeRequirement(t, root+"/requirements-base.txt", "-r requirements.txt\nbase-pkg==2.0\n")

	manifest := NewManifest(root + "/requirements.txt")

	if manifest.Dependencies["root-pkg"] != "= 1.0" {
		t.Fatalf("Expected root-pkg, got %v", manifest.Dependencies)
	}

	if manifest.Dependencies["base-pkg"] != "= 2.0" {
		t.Fatalf("Expected base-pkg, got %v", manifest.Dependencies)
	}

}

func writeRequirement(t *testing.T, file string, content string) {

	t.Helper()

	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

}
