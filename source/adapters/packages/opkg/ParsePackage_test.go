package opkg

import "os"
import "strings"
import "testing"
import "tholian-core/structs"
import "tholian-core/types"

func readFixture(t *testing.T, name string) string {

	t.Helper()

	buffer, err := os.ReadFile("testdata/" + name)

	if err != nil {
		t.Fatalf("Cannot read testdata/%s: %s", name, err.Error())
	}

	return string(buffer)

}

func findBlock(buffer string, name string) string {

	blocks := strings.Split(strings.ReplaceAll(buffer, "\r\n", "\n"), "\n\n")
	prefix := "Package: " + name + "\n"

	for b := 0; b < len(blocks); b++ {

		block := strings.TrimSpace(blocks[b])

		if strings.HasPrefix(block, prefix) {
			return block
		}

	}

	return ""

}

func findPackage(packages []structs.Package, name string) (structs.Package, bool) {

	for p := 0; p < len(packages); p++ {
		if packages[p].Name == name {
			return packages[p], true
		}
	}

	return structs.Package{}, false

}

func hasDependency(pkg structs.Package, name string) bool {

	for d := 0; d < len(pkg.Dependencies); d++ {
		if pkg.Dependencies[d].Name == name {
			return true
		}
	}

	return false

}

func TestParsePackage(t *testing.T) {

	buffer := readFixture(t, "status")
	block := findBlock(buffer, "uhttpd")

	if block == "" {
		t.Fatal("Expected testdata/status to contain uhttpd")
	}

	pkg := structs.NewPackage("opkg")
	ParsePackage(block, &pkg)

	if pkg.Name != "uhttpd" {
		t.Fatalf("Expected name uhttpd, got %s", pkg.Name)
	}

	if pkg.Version.Equal(types.ToVersion("2023-06-25-34a8a74d-1")) == false {
		t.Fatalf("Expected old uhttpd version, got %s", pkg.Version.String())
	}

	if pkg.Architecture.String() != "x86_64" {
		t.Fatalf("Expected architecture x86_64, got %s", pkg.Architecture.String())
	}

	if pkg.Manager.String() != "opkg" {
		t.Fatalf("Expected manager opkg, got %s", pkg.Manager.String())
	}

	if pkg.Datetime.IsValid() == false {
		t.Fatal("Expected datetime from Installed-Time")
	}

	if pkg.URL != "https://openwrt.org/" {
		t.Fatalf("Expected homepage URL, got %s", pkg.URL)
	}

	if len(pkg.Maintainers) != 1 || pkg.Maintainers[0].Name != "OpenWrt Team" {
		t.Fatalf("Expected OpenWrt Team maintainer, got %v", pkg.Maintainers)
	}

	if hasDependency(pkg, "libc") == false {
		t.Fatal("Expected libc dependency")
	}

	if hasDependency(pkg, "libgcc1") == false {
		t.Fatal("Expected libgcc1 dependency")
	}

	if hasDependency(pkg, "missingpkg") == false {
		t.Fatal("Expected missingpkg dependency")
	}

	if len(pkg.Unresolved) != 1 {
		t.Fatalf("Expected 1 unresolved alternative, got %d", len(pkg.Unresolved))
	}

	if len(pkg.Unresolved[0].Candidates) != 2 {
		t.Fatalf("Expected 2 alternative candidates, got %d", len(pkg.Unresolved[0].Candidates))
	}

	if len(pkg.Provides) != 2 {
		t.Fatalf("Expected 2 provides, got %d", len(pkg.Provides))
	}

	if len(pkg.Replaces) != 1 || pkg.Replaces[0].Name != "old-uhttpd" {
		t.Fatalf("Expected old-uhttpd replace, got %v", pkg.Replaces)
	}

	if len(pkg.Conflicts) != 1 || pkg.Conflicts[0].Name != "lighttpd" {
		t.Fatalf("Expected lighttpd conflict, got %v", pkg.Conflicts)
	}

	if pkg.IsValid() == false {
		t.Fatal("Expected parsed package to be valid")
	}

}
