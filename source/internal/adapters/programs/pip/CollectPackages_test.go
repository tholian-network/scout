package pip

import "os"
import "testing"
import "tholian-core/types"

func TestCollectPackages(t *testing.T) {

	packages := collectPackages(fixture(t, "project"))

	if len(packages) != 3 {
		t.Fatalf("Expected 3 installed packages, got %d: %v", len(packages), packages)
	}

	if packages[0].Name != "Flask" || packages[1].Name != "legacy" || packages[2].Name != "requests" {
		t.Fatalf("Expected sorted packages, got %v", packages)
	}

	flask := packages[0]
	legacy := packages[1]
	requests := packages[2]

	if flask.Version.Equal(types.ToVersion("3.0.2")) == false {
		t.Fatalf("Expected Flask 3.0.2, got %s", flask.Version.String())
	}

	if len(flask.Dependencies) != 6 {
		t.Fatalf("Expected 6 Flask dependencies, got %d", len(flask.Dependencies))
	}

	flask_deps := make(map[string]string)

	for d := 0; d < len(flask.Dependencies); d++ {
		flask_deps[flask.Dependencies[d].Name] = flask.Dependencies[d].Version
	}

	if flask_deps["Werkzeug"] != ">= 3.0.0" ||
		flask_deps["Jinja2"] != ">= 3.1.2" ||
		flask_deps["itsdangerous"] != ">= 2.1.2" ||
		flask_deps["click"] != ">= 8.1.3" ||
		flask_deps["blinker"] != ">= 1.6.2" ||
		flask_deps["importlib-metadata"] != ">= 3.6.0" {
		t.Fatalf("Expected normalized Flask dependencies, got %v", flask.Dependencies)
	}

	if len(flask.Maintainers) != 1 || flask.Maintainers[0].Name != "Pallets Team" {
		t.Fatalf("Expected Pallets Team maintainer, got %v", flask.Maintainers)
	}

	if requests.URL != "https://requests.readthedocs.io" {
		t.Fatalf("Expected requests homepage, got %s", requests.URL)
	}

	if len(requests.Dependencies) != 4 {
		t.Fatalf("Expected 4 requests dependencies after skipping extras, got %d", len(requests.Dependencies))
	}

	requests_deps := make(map[string]string)

	for d := 0; d < len(requests.Dependencies); d++ {
		requests_deps[requests.Dependencies[d].Name] = requests.Dependencies[d].Version
	}

	if requests_deps["charset-normalizer"] != ">= 2" ||
		requests_deps["idna"] != ">= 2.5" ||
		requests_deps["urllib3"] != ">= 1.21.1" ||
		requests_deps["certifi"] != ">= 2017.4.17" {
		t.Fatalf("Expected normalized requests dependencies, got %v", requests.Dependencies)
	}

	if len(requests.Maintainers) != 1 || requests.Maintainers[0].Name != "Kenneth Reitz" {
		t.Fatalf("Expected Kenneth Reitz maintainer, got %v", requests.Maintainers)
	}

	if legacy.URL != "" {
		t.Fatalf("Expected no legacy homepage, got %s", legacy.URL)
	}

	if flask.IsValid() == false || legacy.IsValid() == false || requests.IsValid() == false {
		t.Fatal("Expected installed packages to be valid")
	}

}

func TestCollectPackagesPEP582(t *testing.T) {

	root := t.TempDir()
	metadata := root + "/__pypackages__/3.12/lib/demo-1.0.0.dist-info"

	if err := os.MkdirAll(metadata, 0755); err != nil {
		t.Fatal(err)
	}

	buffer := "Metadata-Version: 2.1\nName: demo\nVersion: 1.0.0\n\nDescription.\n"

	if err := os.WriteFile(metadata+"/METADATA", []byte(buffer), 0644); err != nil {
		t.Fatal(err)
	}

	packages := collectPackages(root)

	if len(packages) != 1 || packages[0].Name != "demo" {
		t.Fatalf("Expected PEP 582 demo package, got %v", packages)
	}

}
