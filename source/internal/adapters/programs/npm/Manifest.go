package npm

import "tholian-core/matchers"
import "tholian-core/structs"
import "encoding/json"
import "os"
import "strings"

type Schema struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Version         string            `json:"version"`
	Private         bool              `json:"private"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type Manifest struct {
	file            string
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Version         string            `json:"version"`
	Private         bool              `json:"private"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func NewManifest(file string) Manifest {

	var manifest Manifest

	manifest.file = strings.TrimSpace(file)
	manifest.Dependencies = make(map[string]string)
	manifest.DevDependencies = make(map[string]string)

	manifest.Parse()

	return manifest

}

func (manifest *Manifest) Parse() {

	if strings.HasSuffix(manifest.file, "/package.json") {

		buffer, err1 := os.ReadFile(manifest.file)

		if err1 == nil {

			var schema Schema

			err2 := json.Unmarshal(buffer, &schema)

			if err2 == nil {

				if schema.Name != "" {
					manifest.Name = schema.Name
				}

				if schema.Description != "" {
					manifest.Description = schema.Description
				}

				if schema.Version != "" {
					manifest.Version = schema.Version
				}

				if schema.Private == true {
					manifest.Private = true
				}

				if len(schema.Dependencies) > 0 {

					for name, version := range schema.Dependencies {
						manifest.Dependencies[name] = version
					}

				}

				if len(schema.DevDependencies) > 0 {

					for name, version := range schema.DevDependencies {
						manifest.DevDependencies[name] = version
					}

				}

			}

		}

	}

}

func (manifest *Manifest) ToPackage() structs.Package {

	var result structs.Package = structs.NewPackage("npm")

	result.SetArchitecture("any")
	result.SetVendor("npm")

	if manifest.Private == false {
		result.SetURL("https://www.npmjs.com/package/" + manifest.Name)
	}

	if manifest.Name != "" {
		result.SetName(manifest.Name)
	}

	if manifest.Version != "" {
		result.SetVersion(manifest.Version)
	}

	if len(manifest.Dependencies) > 0 {

		for name, version := range manifest.Dependencies {

			dependency := matchers.NewPackage()
			dependency.SetName(name)
			dependency.SetVersion(version)
			dependency.SetArchitecture("any")
			dependency.SetManager("npm")
			dependency.SetVendor("npm")

			result.AddDependency(dependency)

		}

	}

	if len(manifest.DevDependencies) > 0 {

		for name, version := range manifest.DevDependencies {

			dependency := matchers.NewPackage()
			dependency.SetName(name)
			dependency.SetVersion(version)
			dependency.SetArchitecture("any")
			dependency.SetManager("npm")
			dependency.SetVendor("npm")

			result.AddDependency(dependency)

		}

	}

	return result

}
