package pip

import "testing"
import "tholian-core/structs"

func TestNormalizeVersion(t *testing.T) {

	tests := []struct {
		value string
		want  string
		ok    bool
	}{
		{"= 2.31.0", "= 2.31.0", true},
		{"==2.31.0", "= 2.31.0", true},
		{"===2.31.0", "= 2.31.0", true},
		{"~=1.4.2", ">= 1.4.2", true},
		{">=1.0,<2.0", ">= 1.0", true},
		{"<2.0", "< 2.0", true},
		{"==1.2.*", ">= 1.2.0", true},
		{"!=1.5", "", false},
		{"1!2.0", "= 1:2.0", true},
		{"*", "", false},
		{"latest", "", false},
	}

	for _, test := range tests {

		got, ok := normalizeVersion(test.value)

		if got != test.want || ok != test.ok {
			t.Fatalf("normalizeVersion(%q) = %q, %v; want %q, %v", test.value, got, ok, test.want, test.ok)
		}

	}

}

func TestAddDependencyOr(t *testing.T) {

	result := structs.NewPackage("pip")
	addDependency(&result, "requests", "==2.0.0 || ==2.1.0")

	if len(result.Unresolved) != 1 {
		t.Fatalf("Expected 1 unresolved OR, got %d", len(result.Unresolved))
	}

	candidates := result.Unresolved[0].Candidates

	if len(candidates) != 2 {
		t.Fatalf("Expected 2 candidates, got %d", len(candidates))
	}

	if candidates[0].Version != "= 2.0.0" || candidates[1].Version != "= 2.1.0" {
		t.Fatalf("Expected normalized OR candidates, got %v", candidates)
	}

}

func TestAddDependencyFallsBackToAny(t *testing.T) {

	result := structs.NewPackage("pip")
	addDependency(&result, "requests", "!=1.0")

	if len(result.Unresolved) != 0 {
		t.Fatalf("Expected no unresolved entries, got %d", len(result.Unresolved))
	}

	if len(result.Dependencies) != 1 {
		t.Fatalf("Expected one fallback dependency, got %d", len(result.Dependencies))
	}

	if result.Dependencies[0].Version != "any" {
		t.Fatalf("Expected any version fallback, got %s", result.Dependencies[0].Version)
	}

}
