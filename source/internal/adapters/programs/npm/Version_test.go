package npm

import "testing"
import "tholian-core/structs"

func TestNormalizeVersion(t *testing.T) {

	tests := []struct {
		value string
		want  string
		ok    bool
	}{
		{"1.2.3", "= 1.2.3", true},
		{"^1.2.3", ">= 1.2.3", true},
		{"~1.2.3", ">= 1.2.3", true},
		{"1.2", ">= 1.2", true},
		{"1.2.x", ">= 1.2.0", true},
		{"1.2.3 - 2.3.4", ">= 1.2.3", true},
		{">=1.2.3 <2.0.0", ">= 1.2.3", true},
		{"<2.0.0", "< 2.0.0", true},
		{"npm:foo@^1.2.3", ">= 1.2.3", true},
		{"workspace:^1.2.3", ">= 1.2.3", true},
		{"latest", "", false},
		{"workspace:*", "", false},
		{"file:../foo", "", false},
		{"git+https://example.com/foo.git", "", false},
		{"*", "", false},
		{"!=1.2.3", "", false},
	}

	for _, test := range tests {

		got, ok := normalizeVersion(test.value)

		if got != test.want || ok != test.ok {
			t.Fatalf("normalizeVersion(%q) = %q, %v; want %q, %v", test.value, got, ok, test.want, test.ok)
		}

	}

}

func TestAddDependencyOr(t *testing.T) {

	result := structs.NewPackage("npm")
	addDependency(&result, "left-pad", "^1.0.0 || ^2.0.0")

	if len(result.Unresolved) != 1 {
		t.Fatalf("Expected 1 unresolved OR, got %d", len(result.Unresolved))
	}

	candidates := result.Unresolved[0].Candidates

	if len(candidates) != 2 {
		t.Fatalf("Expected 2 candidates, got %d", len(candidates))
	}

	if candidates[0].Name != "left-pad" || candidates[0].Version != ">= 1.0.0" {
		t.Fatalf("Expected left-pad >= 1.0.0, got %v", candidates[0])
	}

	if candidates[1].Name != "left-pad" || candidates[1].Version != ">= 2.0.0" {
		t.Fatalf("Expected left-pad >= 2.0.0, got %v", candidates[1])
	}

}

func TestAddDependencySkipsBadAlternative(t *testing.T) {

	result := structs.NewPackage("npm")
	addDependency(&result, "left-pad", "latest || ^1.2.3")

	if len(result.Unresolved) != 1 {
		t.Fatalf("Expected 1 unresolved OR, got %d", len(result.Unresolved))
	}

	if len(result.Unresolved[0].Candidates) != 1 {
		t.Fatalf("Expected only the representable candidate, got %d", len(result.Unresolved[0].Candidates))
	}

	if result.Unresolved[0].Candidates[0].Version != ">= 1.2.3" {
		t.Fatalf("Expected >= 1.2.3, got %s", result.Unresolved[0].Candidates[0].Version)
	}

}

func TestAddDependencyFallsBackToAny(t *testing.T) {

	result := structs.NewPackage("npm")
	addDependency(&result, "left-pad", "latest || file:../left-pad")

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
