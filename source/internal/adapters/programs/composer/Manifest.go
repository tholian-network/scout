package composer

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
	Authors         []types.Maintainer
	Dependencies    map[string]string
	DevDependencies map[string]string
	Conflicts       map[string]string
	Replaces        map[string]string
	Provides        map[string]string
}

func NewManifest(file string) Manifest {

	var manifest Manifest

	manifest.file = strings.TrimSpace(file)
	manifest.Dependencies = make(map[string]string)
	manifest.DevDependencies = make(map[string]string)
	manifest.Conflicts = make(map[string]string)
	manifest.Replaces = make(map[string]string)
	manifest.Provides = make(map[string]string)
	manifest.Authors = make([]types.Maintainer, 0)

	manifest.Parse()

	return manifest

}

func (manifest *Manifest) Parse() {

	if strings.HasSuffix(manifest.file, "/composer.json") {

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

				if value := decodeDependencyMap(schema["require"]); len(value) > 0 {
					manifest.Dependencies = value
				}

				if value := decodeDependencyMap(schema["require-dev"]); len(value) > 0 {
					manifest.DevDependencies = value
				}

				if value := decodeDependencyMap(schema["conflict"]); len(value) > 0 {
					manifest.Conflicts = value
				}

				if value := decodeDependencyMap(schema["replace"]); len(value) > 0 {
					manifest.Replaces = value
				}

				if value := decodeDependencyMap(schema["provide"]); len(value) > 0 {
					manifest.Provides = value
				}

				manifest.Authors = decodeAuthors(schema["authors"])

			}

		}

	}

}

func (manifest *Manifest) ToPackage() structs.Package {

	var result structs.Package = structs.NewPackage("composer")

	result.SetArchitecture("any")
	result.SetVendor("composer")

	if manifest.Homepage != "" {
		result.SetURL(manifest.Homepage)
	}

	if result.URL == "" && strings.Contains(manifest.Name, "/") {
		result.SetURL("https://packagist.org/packages/" + manifest.Name)
	}

	if manifest.Name != "" {
		result.SetName(manifest.Name)
	}

	if manifest.Version != "" {
		result.SetVersion(manifest.Version)
	}

	for a := 0; a < len(manifest.Authors); a++ {
		result.AddMaintainer(manifest.Authors[a])
	}

	addDependencies(&result, manifest.Dependencies, manifest.Version)
	addDependencies(&result, manifest.DevDependencies, manifest.Version)
	addConflicts(&result, manifest.Conflicts, manifest.Version)
	addReplaces(&result, manifest.Replaces, manifest.Version)
	addProvides(&result, manifest.Provides, manifest.Version)
	setManifestDatetime(manifest.file, &result)

	return result

}

func addDependencies(result *structs.Package, dependencies map[string]string, own_version string) {

	names := sortedNames(dependencies)

	for n := 0; n < len(names); n++ {
		addDependency(result, names[n], dependencies[names[n]], own_version)
	}

}

func addConflicts(result *structs.Package, conflicts map[string]string, own_version string) {

	names := sortedNames(conflicts)

	for n := 0; n < len(names); n++ {

		if version, ok := normalizeRelation(conflicts[names[n]], own_version); ok == true {
			result.AddConflict(newMatcher(names[n], version))
		}

	}

}

func addReplaces(result *structs.Package, replaces map[string]string, own_version string) {

	names := sortedNames(replaces)

	for n := 0; n < len(names); n++ {

		if version, ok := normalizeRelation(replaces[names[n]], own_version); ok == true {
			result.AddReplace(newMatcher(names[n], version))
		}

	}

}

func addProvides(result *structs.Package, provides map[string]string, own_version string) {

	names := sortedNames(provides)

	for n := 0; n < len(names); n++ {

		if version, ok := normalizeRelation(provides[names[n]], own_version); ok == true {
			result.AddProvide(newMatcher(names[n], version))
		}

	}

}

func sortedNames(dependencies map[string]string) []string {

	names := make([]string, 0)

	for name, _ := range dependencies {

		name = strings.TrimSpace(name)

		if name != "" && isPlatformPackage(name) == false {
			names = append(names, name)
		}

	}

	sort.Strings(names)

	return names

}

func isPlatformPackage(name string) bool {

	lower := strings.ToLower(strings.TrimSpace(name))

	if lower == "php" || lower == "hhvm" || lower == "composer" {
		return true
	}

	if strings.HasPrefix(lower, "ext-") ||
		strings.HasPrefix(lower, "lib-") ||
		strings.HasPrefix(lower, "composer-plugin-api") ||
		strings.HasPrefix(lower, "composer-runtime-api") {
		return true
	}

	return false

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

func decodeAuthors(value json.RawMessage) []types.Maintainer {

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
				result = appendAuthors(result, decodeAuthor(entries[e]))
			}

			return result

		}

		return appendAuthors(result, decodeAuthor(value))

	}

}

func decodeAuthor(value json.RawMessage) []types.Maintainer {

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

func appendAuthors(result []types.Maintainer, values []types.Maintainer) []types.Maintainer {

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
