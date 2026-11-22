package composer

import "testing"
import "tholian-core/structs"

func TestNormalizeVersion(t *testing.T) {

	tests := []struct {
		value      string
		ownVersion string
		want       string
		ok         bool
	}{
		{"1.2.3", "", "= 1.2.3", true},
		{"1.2", "", "= 1.2", true},
		{"^6.4", "", ">= 6.4", true},
		{"~1.2", "", ">= 1.2", true},
		{"1.2.*", "", ">= 1.2.0", true},
		{">=1.0 <2.0", "", ">= 1.0", true},
		{">=1.0,<2.0", "", ">= 1.0", true},
		{"<2.0", "", "< 2.0", true},
		{"^1.0@dev", "", ">= 1.0", true},
		{"self.version", "2.0.0", "= 2.0.0", true},
		{"self.version", "", "", false},
		{"dev-main", "", "", false},
		{"1.0.x-dev", "", "", false},
		{"!=1.5", "", "", false},
		{"*", "", "", false},
	}

	for _, test := range tests {

		got, ok := normalizeVersion(test.value, test.ownVersion)

		if got != test.want || ok != test.ok {
			t.Fatalf("normalizeVersion(%q, %q) = %q, %v; want %q, %v", test.value, test.ownVersion, got, ok, test.want, test.ok)
		}

	}

}

func TestAddDependencyOr(t *testing.T) {

	result := structs.NewPackage("composer")
	addDependency(&result, "acme/a", "^1.0 || ^2.0", "")

	if len(result.Unresolved) != 1 {
		t.Fatalf("Expected 1 unresolved OR, got %d", len(result.Unresolved))
	}

	candidates := result.Unresolved[0].Candidates

	if len(candidates) != 2 {
		t.Fatalf("Expected 2 candidates, got %d", len(candidates))
	}

	if candidates[0].Version != ">= 1.0" || candidates[1].Version != ">= 2.0" {
		t.Fatalf("Expected normalized OR candidates, got %v", candidates)
	}

}

func TestAddDependencySkipsBadAlternative(t *testing.T) {

	result := structs.NewPackage("composer")
	addDependency(&result, "acme/a", "dev-main || ^1.2", "")

	if len(result.Unresolved) != 1 {
		t.Fatalf("Expected 1 unresolved OR, got %d", len(result.Unresolved))
	}

	if len(result.Unresolved[0].Candidates) != 1 {
		t.Fatalf("Expected only the representable candidate, got %d", len(result.Unresolved[0].Candidates))
	}

	if result.Unresolved[0].Candidates[0].Version != ">= 1.2" {
		t.Fatalf("Expected >= 1.2, got %s", result.Unresolved[0].Candidates[0].Version)
	}

}

func TestAddDependencyFallsBackToAny(t *testing.T) {

	result := structs.NewPackage("composer")
	addDependency(&result, "acme/a", "dev-main || dev-develop", "")

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
