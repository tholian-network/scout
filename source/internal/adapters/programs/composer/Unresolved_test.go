package composer

import "testing"
import "tholian-core/structs"

func findCollected(packages []structs.Package, name string) (structs.Package, bool) {

	for p := 0; p < len(packages); p++ {
		if packages[p].Name == name {
			return packages[p], true
		}
	}

	return structs.Package{}, false

}

func TestCollectPackagesResolvesOr(t *testing.T) {

	packages := collectPackages(fixture(t, "or"))

	if len(packages) != 3 {
		t.Fatalf("Expected 3 packages, got %d", len(packages))
	}

	parent, ok := findCollected(packages, "acme/parent")

	if ok == false {
		t.Fatal("Expected acme/parent package")
	}

	if len(parent.Dependencies) != 1 {
		t.Fatalf("Expected 1 resolved dependency, got %d", len(parent.Dependencies))
	}

	if parent.Dependencies[0].Name != "acme/a" || parent.Dependencies[0].Version != "0:1.5.0" {
		t.Fatalf("Expected resolved acme/a 0:1.5.0, got %v", parent.Dependencies[0])
	}

	if len(parent.Unresolved) != 0 {
		t.Fatalf("Expected parent OR to resolve, got %d unresolved", len(parent.Unresolved))
	}

	missing, ok := findCollected(packages, "acme/missing-parent")

	if ok == false {
		t.Fatal("Expected acme/missing-parent package")
	}

	if len(missing.Unresolved) != 1 {
		t.Fatalf("Expected 1 unresolved entry, got %d", len(missing.Unresolved))
	}

	if len(missing.Unresolved[0].Candidates) != 2 {
		t.Fatalf("Expected 2 unresolved candidates, got %d", len(missing.Unresolved[0].Candidates))
	}

}

func TestAssembleProgramOr(t *testing.T) {

	SUPPORTED = true

	program := structs.NewProgram(1337, "php")
	program.SetCommand("/usr/bin/php")
	program.SetFolder(fixture(t, "or"))

	AssembleProgramPackages(&program)

	if len(program.Packages) != 3 {
		t.Fatalf("Expected 3 program packages, got %d", len(program.Packages))
	}

	AssembleProgramDependencies(&program)

	if len(program.Dependencies) != 1 {
		t.Fatalf("Expected 1 resolved program dependency, got %d", len(program.Dependencies))
	}

	if program.Dependencies[0].Name != "acme/a" || program.Dependencies[0].Version != "0:1.5.0" {
		t.Fatalf("Expected resolved acme/a, got %v", program.Dependencies[0])
	}

	if len(program.Unresolved) != 1 {
		t.Fatalf("Expected 1 unresolved program dependency, got %d", len(program.Unresolved))
	}

	if len(program.Unresolved[0].Candidates) != 2 {
		t.Fatalf("Expected 2 unresolved candidates, got %d", len(program.Unresolved[0].Candidates))
	}

}
