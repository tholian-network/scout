package composer

import "tholian-core/structs"
import "encoding/json"
import "os"
import "sort"
import "strings"
import "time"

func CollectPackages(folder string) []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true && folder != "" && strings.HasPrefix(folder, "/") {
		collected = collectPackages(folder)
	}

	return collected

}

func collectPackages(folder string) []structs.Package {

	collected := make([]structs.Package, 0)

	if folder != "" && strings.HasPrefix(folder, "/") {

		installed := folder + "/vendor/composer/installed.json"
		lock := folder + "/composer.lock"

		if isFile(installed) {
			collected = collectInstalledFile(installed, folder)
		}

		if len(collected) == 0 && isFile(lock) {
			collected = collectLockFile(lock, folder)
		}

		if len(collected) == 0 {
			collected = collectVendorPackages(folder)
		}

	}

	for c := 0; c < len(collected); c++ {
		collected[c].ResolveDependencies(collected)
	}

	sort.Slice(collected, func(a int, b int) bool {

		if collected[a].Name == collected[b].Name {
			return collected[a].Version.String() < collected[b].Version.String()
		}

		return collected[a].Name < collected[b].Name

	})

	return collected

}

func collectInstalledFile(file string, folder string) []structs.Package {

	collected := make([]structs.Package, 0)

	buffer, err := os.ReadFile(file)

	if err == nil {

		trimmed := strings.TrimSpace(string(buffer))

		if strings.HasPrefix(trimmed, "[") {

			var entries []map[string]json.RawMessage

			if json.Unmarshal(buffer, &entries) == nil {

				for e := 0; e < len(entries); e++ {
					pkg := packageFromSchema(entries[e], folder, file)
					appendPackage(&collected, pkg)
				}

			}

		} else {

			var schema struct {
				Packages []map[string]json.RawMessage `json:"packages"`
			}

			if json.Unmarshal(buffer, &schema) == nil {

				for p := 0; p < len(schema.Packages); p++ {
					pkg := packageFromSchema(schema.Packages[p], folder, file)
					appendPackage(&collected, pkg)
				}

			}

		}

	}

	return collected

}

func collectLockFile(file string, folder string) []structs.Package {

	collected := make([]structs.Package, 0)

	buffer, err := os.ReadFile(file)

	if err == nil {

		var schema struct {
			Packages    []map[string]json.RawMessage `json:"packages"`
			PackagesDev []map[string]json.RawMessage `json:"packages-dev"`
		}

		if json.Unmarshal(buffer, &schema) == nil {

			for p := 0; p < len(schema.Packages); p++ {
				pkg := packageFromSchema(schema.Packages[p], folder, file)
				appendPackage(&collected, pkg)
			}

			for p := 0; p < len(schema.PackagesDev); p++ {
				pkg := packageFromSchema(schema.PackagesDev[p], folder, file)
				appendPackage(&collected, pkg)
			}

		}

	}

	return collected

}

func collectVendorPackages(folder string) []structs.Package {

	collected := make([]structs.Package, 0)

	vendor := folder + "/vendor"

	if isDirectory(vendor) {

		vendors, err1 := os.ReadDir(vendor)

		if err1 == nil {

			for v := 0; v < len(vendors); v++ {

				vendorName := vendors[v].Name()

				if strings.HasPrefix(vendorName, ".") {
					continue
				}

				vendorFolder := vendor + "/" + vendorName

				if isDirectory(vendorFolder) == false {
					continue
				}

				packages, err2 := os.ReadDir(vendorFolder)

				if err2 == nil {

					for p := 0; p < len(packages); p++ {

						packageName := packages[p].Name()

						if strings.HasPrefix(packageName, ".") {
							continue
						}

						packageFolder := vendorFolder + "/" + packageName

						if isDirectory(packageFolder) {

							manifest := NewManifest(packageFolder + "/composer.json")
							appendPackage(&collected, manifest.ToPackage())

						}

					}

				}

			}

		}

	}

	return collected

}

func packageFromSchema(schema map[string]json.RawMessage, folder string, source string) structs.Package {

	var result structs.Package = structs.NewPackage("composer")

	result.SetArchitecture("any")
	result.SetVendor("composer")

	name := decodeString(schema["name"])
	version := decodeString(schema["version"])

	if version == "" {
		version = decodeString(schema["version_normalized"])
	}

	if homepage := decodeString(schema["homepage"]); homepage != "" {
		result.SetURL(homepage)
	}

	if result.URL == "" && strings.Contains(name, "/") {
		result.SetURL("https://packagist.org/packages/" + name)
	}

	if name != "" {
		result.SetName(name)
	}

	if version != "" {
		result.SetVersion(version)
	}

	authors := decodeAuthors(schema["authors"])

	for a := 0; a < len(authors); a++ {
		result.AddMaintainer(authors[a])
	}

	addDependencies(&result, decodeDependencyMap(schema["require"]), version)
	addConflicts(&result, decodeDependencyMap(schema["conflict"]), version)
	addReplaces(&result, decodeDependencyMap(schema["replace"]), version)
	addProvides(&result, decodeDependencyMap(schema["provide"]), version)
	setPackageDatetime(folder, name, source, &result)

	return result

}

func setPackageDatetime(folder string, name string, source string, result *structs.Package) {

	if name != "" {

		manifest := folder + "/vendor/" + name + "/composer.json"

		if stat, err := os.Stat(manifest); err == nil {
			result.SetDatetime(stat.ModTime().Format(time.RFC3339))
			return
		}

	}

	if stat, err := os.Stat(source); err == nil {
		result.SetDatetime(stat.ModTime().Format(time.RFC3339))
	} else {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

}

func appendPackage(collected *[]structs.Package, pkg structs.Package) {

	if pkg.Name != "" && pkg.Version.IsValid() {

		found := false

		for c := 0; c < len(*collected); c++ {

			if (*collected)[c].Equal(pkg) {
				found = true
				break
			}

		}

		if found == false {
			*collected = append(*collected, pkg)
		}

	}

}
