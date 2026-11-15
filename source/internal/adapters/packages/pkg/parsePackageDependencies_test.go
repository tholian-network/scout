package pkg

import "strings"
import "testing"

func TestParsePackageDependencies(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"curl\tca_root_nss\t3.115",
		"curl\tlibnghttp2\t1.64.0",
		"e2e-scout-pkg\tcurl\t8.14.1_1,1",
	}, "\r\n"))

	dependencies := parsePackageDependencies(string(buffer), "")

	if len(dependencies) != 3 {
		t.Fatalf("Expected 3 dependencies, got %d", len(dependencies))
	}

	if dependencies[0].Name != "ca_root_nss" || dependencies[0].Version != "3.115" {
		t.Fatalf("Unexpected dependency: %v", dependencies[0])
	}

	if dependencies[1].Name != "libnghttp2" || dependencies[1].Version != "1.64.0" {
		t.Fatalf("Unexpected dependency: %v", dependencies[1])
	}

	if dependencies[2].Name != "curl" || dependencies[2].Version != "8.14.1_1,1" {
		t.Fatalf("Unexpected dependency: %v", dependencies[2])
	}

}

func TestParsePackageDependenciesByParent(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"curl\tca_root_nss\t3.115",
		"curl\tlibnghttp2\t1.64.0",
		"e2e-scout-pkg\tcurl\t8.14.1_1,1",
	}, "\r\n"))

	dependencies := parsePackageDependencies(string(buffer), "e2e-scout-pkg")

	if len(dependencies) != 1 {
		t.Fatalf("Expected 1 dependency for e2e-scout-pkg, got %d", len(dependencies))
	}

	if dependencies[0].Name != "curl" || dependencies[0].Version != "8.14.1_1,1" {
		t.Fatalf("Expected curl dependency, got %v", dependencies[0])
	}

	if dependencies[0].Architecture != "any" {
		t.Fatalf("Expected architecture any, got %s", dependencies[0].Architecture)
	}

	if dependencies[0].Manager != "pkg" {
		t.Fatalf("Expected manager pkg, got %s", dependencies[0].Manager)
	}

}

func TestParsePackageDependenciesUnknownParent(t *testing.T) {

	buffer := []byte(strings.Join([]string{
		"curl\tca_root_nss\t3.115",
		"curl\tlibnghttp2\t1.64.0",
		"e2e-scout-pkg\tcurl\t8.14.1_1,1",
	}, "\r\n"))

	dependencies := parsePackageDependencies(string(buffer), "missing")

	if len(dependencies) != 0 {
		t.Fatalf("Expected 0 dependencies for missing parent, got %d", len(dependencies))
	}

}
