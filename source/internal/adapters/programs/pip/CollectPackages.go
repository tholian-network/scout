package pip

import "tholian-core/structs"
import "tholian-core/types"
import "os"
import "path/filepath"
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

		for _, site := range sitePackagesDirs(folder) {
			collectSitePackages(&collected, site)
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

func sitePackagesDirs(folder string) []string {

	result := make([]string, 0)
	found := make(map[string]bool)

	venvs := []string{
		folder + "/.venv",
		folder + "/venv",
		folder + "/env",
		folder,
	}

	for v := 0; v < len(venvs); v++ {

		venv := venvs[v]

		if isVenv(venv) {

			appendGlob(&result, found, venv+"/lib/python*/site-packages")
			appendGlob(&result, found, venv+"/lib64/python*/site-packages")
			appendGlob(&result, found, venv+"/Lib/site-packages")

		}

	}

	appendGlob(&result, found, folder+"/__pypackages__/*/lib")
	appendGlob(&result, found, folder+"/lib/python*/site-packages")

	return result

}

func appendGlob(result *[]string, found map[string]bool, pattern string) {

	values, err := filepath.Glob(pattern)

	if err != nil {
		return
	}

	for v := 0; v < len(values); v++ {

		if isDirectory(values[v]) && found[values[v]] == false {
			found[values[v]] = true
			*result = append(*result, values[v])
		}

	}

}

func collectSitePackages(collected *[]structs.Package, site string) {

	entries, err := os.ReadDir(site)

	if err != nil {
		return
	}

	for e := 0; e < len(entries); e++ {

		name := entries[e].Name()
		path := site + "/" + name

		if strings.HasSuffix(name, ".dist-info") && isDirectory(path) {

			appendPackage(collected, parseMetadata(path+"/METADATA"))

		} else if strings.HasSuffix(name, ".egg-info") {

			if isDirectory(path) {
				appendPackage(collected, parseMetadata(path+"/PKG-INFO"))
			} else if isFile(path) {
				appendPackage(collected, parseMetadata(path))
			}

		}

	}

}

func parseMetadata(file string) structs.Package {

	var result structs.Package = structs.NewPackage("pip")

	result.SetArchitecture("any")
	result.SetVendor("pip")

	buffer, err := os.ReadFile(file)

	if err == nil {

		headers := parseHeaders(string(buffer))

		name := firstHeader(headers, "name")
		version := firstHeader(headers, "version")

		if name != "" {
			result.SetName(name)
		}

		if version != "" {
			result.SetVersion(version)
		}

		homepage := firstHeader(headers, "home-page")

		if homepage == "" {
			homepage = projectURL(headers["project-url"])
		}

		if homepage != "" {
			result.SetURL(homepage)
		}

		maintainers := metadataMaintainers(headers)

		for m := 0; m < len(maintainers); m++ {
			result.AddMaintainer(maintainers[m])
		}

		requires := headers["requires-dist"]

		for r := 0; r < len(requires); r++ {

			if isExtraRequirement(requires[r]) {
				continue
			}

			dependency_name, dependency_version := parseRequirement(requires[r])

			if dependency_name != "" {
				addDependency(&result, dependency_name, dependency_version)
			}

		}

	}

	if stat, err := os.Stat(file); err == nil {
		result.SetDatetime(stat.ModTime().Format(time.RFC3339))
	} else {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

	return result

}

func parseHeaders(buffer string) map[string][]string {

	result := make(map[string][]string)

	key := ""
	lines := strings.Split(strings.ReplaceAll(buffer, "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		line := lines[l]

		if line == "" {
			break
		}

		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {

			if key != "" {

				values := result[key]

				if len(values) > 0 {
					values[len(values)-1] += " " + strings.TrimSpace(line)
					result[key] = values
				}

			}

			continue

		}

		if index := strings.Index(line, ":"); index > 0 {

			key = strings.ToLower(strings.TrimSpace(line[0:index]))
			value := strings.TrimSpace(line[index+1:])
			result[key] = append(result[key], value)

		}

	}

	return result

}

func firstHeader(headers map[string][]string, key string) string {

	if values, ok := headers[key]; ok == true {

		if len(values) > 0 {
			return strings.TrimSpace(values[0])
		}

	}

	return ""

}

func projectURL(values []string) string {

	for v := 0; v < len(values); v++ {

		parts := strings.SplitN(values[v], ",", 2)

		if len(parts) == 2 {

			label := strings.ToLower(strings.TrimSpace(parts[0]))
			url := strings.TrimSpace(parts[1])

			if label == "homepage" || label == "home" {
				return url
			}

		}

	}

	for v := 0; v < len(values); v++ {

		if index := strings.Index(values[v], ","); index != -1 {

			url := strings.TrimSpace(values[v][index+1:])

			if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
				return url
			}

		}

	}

	return ""

}

func metadataMaintainers(headers map[string][]string) []types.Maintainer {

	result := make([]types.Maintainer, 0)

	keys := []string{"maintainer", "author"}

	for k := 0; k < len(keys); k++ {

		names := headers[keys[k]]
		emails := headers[keys[k]+"-email"]

		for n := 0; n < len(names); n++ {

			text := strings.TrimSpace(names[n])

			if n < len(emails) && strings.Contains(text, "@") == false {

				email := strings.TrimSpace(emails[n])

				if strings.Contains(email, "@") == true {
					text = text + " <" + email + ">"
				}

			}

			appendMaintainer(&result, types.ToMaintainer(text))

		}

		for e := 0; e < len(emails); e++ {
			appendMaintainer(&result, types.ToMaintainer(emails[e]))
		}

		if len(result) > 0 {
			break
		}

	}

	return result

}

func appendMaintainer(result *[]types.Maintainer, value types.Maintainer) {

	if value.IsValid() {

		found := false

		for r := 0; r < len(*result); r++ {

			if (*result)[r].Equal(value) {
				found = true
				break
			}

		}

		if found == false {
			*result = append(*result, value)
		}

	}

}

func isExtraRequirement(value string) bool {

	lower := strings.ToLower(value)

	if strings.Contains(lower, "extra ==") || strings.Contains(lower, "extra==") {
		return true
	}

	return false

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
