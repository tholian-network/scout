package npm

import "tholian-core/structs"
import "tholian-core/types"
import "encoding/json"
import "os"
import "sort"
import "strings"
import "time"

type Manifest struct {
	file            string
	Name            string
	Description     string
	Version         string
	Homepage        string
	Private         bool
	Dependencies    map[string]string
	DevDependencies map[string]string
	Maintainers     []types.Maintainer
}

func NewManifest(file string) Manifest {

	var manifest Manifest

	manifest.file = strings.TrimSpace(file)
	manifest.Dependencies = make(map[string]string)
	manifest.DevDependencies = make(map[string]string)
	manifest.Maintainers = make([]types.Maintainer, 0)

	manifest.Parse()

	return manifest

}

func (manifest *Manifest) Parse() {

	if strings.HasSuffix(manifest.file, "/package.json") {

		buffer, err1 := os.ReadFile(manifest.file)

		if err1 == nil {

			var schema map[string]json.RawMessage

			err2 := json.Unmarshal(buffer, &schema)

			if err2 == nil {

				if value := decodeString(schema["name"]); value != "" {
					manifest.Name = value
				}

				if value := decodeString(schema["description"]); value != "" {
					manifest.Description = value
				}

				if value := decodeString(schema["version"]); value != "" {
					manifest.Version = value
				}

				if value := decodeString(schema["homepage"]); value != "" {
					manifest.Homepage = value
				}

				manifest.Private = decodeBool(schema["private"])

				if value := decodeDependencyMap(schema["dependencies"]); len(value) > 0 {
					manifest.Dependencies = value
				}

				if value := decodeDependencyMap(schema["devDependencies"]); len(value) > 0 {
					manifest.DevDependencies = value
				}

				manifest.Maintainers = decodeMaintainers(
					schema["author"],
					schema["maintainers"],
					schema["contributors"],
				)

			}

		}

	}

}

func (manifest *Manifest) ToPackage() structs.Package {

	var result structs.Package = structs.NewPackage("npm")

	result.SetArchitecture("any")
	result.SetVendor("npm")

	if manifest.Homepage != "" {
		result.SetURL(manifest.Homepage)
	}

	if result.URL == "" && manifest.Private == false && manifest.Name != "" {
		result.SetURL("https://www.npmjs.com/package/" + manifest.Name)
	}

	if manifest.Name != "" {
		result.SetName(manifest.Name)
	}

	if manifest.Version != "" {
		result.SetVersion(manifest.Version)
	}

	for m := 0; m < len(manifest.Maintainers); m++ {
		result.AddMaintainer(manifest.Maintainers[m])
	}

	addPackageDependencies(&result, manifest.Dependencies)
	addPackageDependencies(&result, manifest.DevDependencies)
	setManifestDatetime(manifest.file, &result)

	return result

}

func addPackageDependencies(result *structs.Package, dependencies map[string]string) {

	names := make([]string, 0)

	for name, _ := range dependencies {
		if strings.TrimSpace(name) != "" {
			names = append(names, name)
		}
	}

	sort.Strings(names)

	for n := 0; n < len(names); n++ {
		addDependency(result, names[n], dependencies[names[n]])
	}

}

func decodeString(value json.RawMessage) string {

	if len(value) == 0 {

		return ""

	} else {

		var result string

		if err := json.Unmarshal(value, &result); err == nil {
			return strings.TrimSpace(result)
		}

		var number json.Number

		if err := json.Unmarshal(value, &number); err == nil {
			return strings.TrimSpace(number.String())
		}

		var boolean bool

		if err := json.Unmarshal(value, &boolean); err == nil {

			if boolean == true {
				return "true"
			} else {
				return "false"
			}

		}

	}

	return ""

}

func decodeBool(value json.RawMessage) bool {

	if len(value) == 0 {

		return false

	} else {

		var result bool

		if err := json.Unmarshal(value, &result); err == nil {
			return result
		}

		var number json.Number

		if err := json.Unmarshal(value, &number); err == nil {

			num, err := number.Int64()

			if err == nil && num != 0 {
				return true
			}

		}

		text := strings.ToLower(decodeString(value))

		if text == "true" || text == "yes" || text == "1" {
			return true
		}

	}

	return false

}

func decodeDependencyMap(value json.RawMessage) map[string]string {

	result := make(map[string]string)

	if len(value) > 0 {

		var schema map[string]json.RawMessage

		if err := json.Unmarshal(value, &schema); err == nil {

			for name, version := range schema {

				name = strings.TrimSpace(name)

				if name != "" {
					result[name] = decodeDependencyVersion(version)
				}

			}

		}

	}

	return result

}

func decodeDependencyVersion(value json.RawMessage) string {

	if len(value) == 0 {

		return "*"

	} else {

		var result string

		if err := json.Unmarshal(value, &result); err == nil {

			result = strings.TrimSpace(result)

			if result != "" {
				return result
			}

			return "*"

		}

		var number json.Number

		if err := json.Unmarshal(value, &number); err == nil {
			return strings.TrimSpace(number.String())
		}

		var schema map[string]json.RawMessage

		if err := json.Unmarshal(value, &schema); err == nil {

			if version := decodeString(schema["version"]); version != "" {
				return version
			}

		}

	}

	return "*"

}

func decodeMaintainers(values ...json.RawMessage) []types.Maintainer {

	result := make([]types.Maintainer, 0)

	for v := 0; v < len(values); v++ {
		result = appendMaintainers(result, decodeMaintainerValue(values[v]))
	}

	return result

}

func decodeMaintainerValue(value json.RawMessage) []types.Maintainer {

	result := make([]types.Maintainer, 0)

	if len(value) == 0 {

		return result

	} else {

		var text string

		if err := json.Unmarshal(value, &text); err == nil {

			maintainer := types.ToMaintainer(text)

			if maintainer.IsValid() {
				result = append(result, maintainer)
			}

			return result

		}

		var entries []json.RawMessage

		if err := json.Unmarshal(value, &entries); err == nil {

			for e := 0; e < len(entries); e++ {
				result = appendMaintainers(result, decodeMaintainerValue(entries[e]))
			}

			return result

		}

		var schema map[string]json.RawMessage

		if err := json.Unmarshal(value, &schema); err == nil {

			maintainer := types.NewMaintainer()
			maintainer.Name = decodeString(schema["name"])
			maintainer.Email = decodeString(schema["email"])

			if maintainer.IsValid() {
				result = append(result, maintainer)
			}

		}

	}

	return result

}

func appendMaintainers(result []types.Maintainer, values []types.Maintainer) []types.Maintainer {

	for v := 0; v < len(values); v++ {

		found := false

		for r := 0; r < len(result); r++ {

			if result[r].Equal(values[v]) {
				found = true
				break
			}

		}

		if found == false {
			result = append(result, values[v])
		}

	}

	return result

}

func setManifestDatetime(file string, result *structs.Package) {

	if stat, err := os.Stat(file); err == nil {
		result.SetDatetime(stat.ModTime().Format(time.RFC3339))
	} else {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

}
