package npm

import "os"
import "testing"
import "tholian-core/types"

func TestCollectPackages(t *testing.T) {

	packages := collectPackages(fixture(t, "project"))

	if len(packages) != 2 {
		t.Fatalf("Expected 2 packages, got %d", len(packages))
	}

	if packages[0].Name != "@scope/pkg" || packages[1].Name != "left-pad" {
		t.Fatalf("Expected sorted scoped and normal packages, got %v", packages)
	}

	if packages[0].Version.Equal(types.ToVersion("2.0.1")) == false {
		t.Fatalf("Expected scoped package version 2.0.1, got %s", packages[0].Version.String())
	}

	if packages[1].Version.Equal(types.ToVersion("1.3.0")) == false {
		t.Fatalf("Expected left-pad version 1.3.0, got %s", packages[1].Version.String())
	}

}

func TestCollectPackagesFollowsSymlinks(t *testing.T) {

	root := t.TempDir()
	modules := root + "/node_modules"

	if err := os.MkdirAll(modules+"/real-pkg", 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll(root+"/linked-pkg", 0755); err != nil {
		t.Fatal(err)
	}

	writePackage(t, modules+"/real-pkg/package.json", "real-pkg", "1.0.0")
	writePackage(t, root+"/linked-pkg/package.json", "linked-pkg", "2.0.0")

	if err := os.Symlink(root+"/linked-pkg", modules+"/linked-pkg"); err != nil {
		t.Fatal(err)
	}

	packages := collectPackages(root)

	if len(packages) != 2 {
		t.Fatalf("Expected 2 packages including symlink, got %d", len(packages))
	}

	if packages[0].Name != "linked-pkg" || packages[1].Name != "real-pkg" {
		t.Fatalf("Expected linked and real package, got %v", packages)
	}

}

func writePackage(t *testing.T, file string, name string, version string) {

	t.Helper()

	buffer := "{\"name\":\"" + name + "\",\"version\":\"" + version + "\"}\n"

	if err := os.WriteFile(file, []byte(buffer), 0644); err != nil {
		t.Fatal(err)
	}

}
