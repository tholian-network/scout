package pip

import "tholian-core/matchers"
import "tholian-core/structs"
import "net/url"
import "os"
import "path/filepath"
import "sort"
import "strings"
import "time"
import "unicode"

type Manifest struct {
	file         string
	Dependencies map[string]string
}

func NewManifest(file string) Manifest {

	var manifest Manifest

	manifest.file = strings.TrimSpace(file)
	manifest.Dependencies = make(map[string]string)

	manifest.Parse()

	return manifest

}

func (manifest *Manifest) Parse() {

	if manifest.file != "" && filepath.Base(manifest.file) == "requirements.txt" {

		visited := make(map[string]bool)
		manifest.Dependencies = parseRequirementsFile(manifest.file, visited)

	}

}

func (manifest *Manifest) ToPackage() structs.Package {

	var result structs.Package = structs.NewPackage("pip")

	result.SetArchitecture("any")
	result.SetVendor("pip")

	project := filepath.Base(filepath.Dir(manifest.file))

	if project != "" && project != "." && project != "/" {
		result.SetName(project)
	}

	addDependencies(&result, manifest.Dependencies)
	setManifestDatetime(manifest.file, &result)

	return result

}

func addDependencies(result *structs.Package, dependencies map[string]string) {

	names := make([]string, 0)

	for name, _ := range dependencies {
		if strings.TrimSpace(name) != "" {
			names = append(names, name)
		}
	}

	sort.Strings(names)

	for n := 0; n < len(names); n++ {

		name := names[n]
		version := strings.TrimSpace(dependencies[name])

		if version == "" {
			version = "*"
		}

		dependency := matchers.NewPackage()
		dependency.SetName(name)
		dependency.SetVersion(version)
		dependency.SetArchitecture("any")
		dependency.SetManager("pip")
		dependency.SetVendor("pip")

		result.AddDependency(dependency)

	}

}

func parseRequirementsFile(file string, visited map[string]bool) map[string]string {

	result := make(map[string]string)

	absolute, err := filepath.Abs(file)

	if err != nil {
		absolute = file
	}

	if visited[absolute] == true {
		return result
	}

	visited[absolute] = true

	buffer, err := os.ReadFile(absolute)

	if err != nil {
		return result
	}

	lines := logicalLines(string(buffer))

	for l := 0; l < len(lines); l++ {

		line := os.ExpandEnv(lines[l])
		line = stripComment(line)
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "-") {

			if value, ok := optionValue(line, "-r", "--requirement"); ok == true {

				include := resolveInclude(absolute, value)

				for name, version := range parseRequirementsFile(include, visited) {
					result[name] = version
				}

			} else if _, ok := optionValue(line, "-c", "--constraint"); ok == true {
				// Constraints restrict versions but do not install packages.
			} else if value, ok := optionValue(line, "-e", "--editable"); ok == true {

				name, version := parseRequirement(value)

				if name != "" {
					result[name] = version
				}

			}

			continue

		}

		if strings.HasPrefix(line, "--hash") {
			continue
		}

		if isExtraRequirement(line) {
			continue
		}

		line = stripHashOptions(line)

		name, version := parseRequirement(line)

		if name != "" {
			result[name] = version
		}

	}

	return result

}

func logicalLines(buffer string) []string {

	result := make([]string, 0)
	current := ""

	lines := strings.Split(strings.ReplaceAll(buffer, "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		line := strings.TrimRight(lines[l], " \t")

		if strings.HasSuffix(line, "\\") {

			current += strings.TrimSuffix(line, "\\")
			current += " "

		} else {

			current += line
			result = append(result, current)
			current = ""

		}

	}

	if current != "" {
		result = append(result, current)
	}

	return result

}

func stripComment(line string) string {

	single := false
	double := false
	escaped := false

	for index, character := range line {

		if escaped == true {
			escaped = false
			continue
		}

		if character == '\\' {
			escaped = true
			continue
		}

		if character == '\'' && double == false {
			single = !single
			continue
		}

		if character == '"' && single == false {
			double = !double
			continue
		}

		if character == '#' && single == false && double == false {

			if index == 0 {
				return ""
			}

			previous := rune(line[index-1])

			if unicode.IsSpace(previous) {
				return line[0:index]
			}

		}

	}

	return line

}

func optionValue(line string, short string, long string) (string, bool) {

	if short != "" {

		if strings.HasPrefix(line, short+" ") {
			return strings.TrimSpace(line[len(short)+1:]), true
		}

		if strings.HasPrefix(line, short+"=") {
			return strings.TrimSpace(line[len(short)+1:]), true
		}

	}

	if long != "" {

		if strings.HasPrefix(line, long+" ") {
			return strings.TrimSpace(line[len(long)+1:]), true
		}

		if strings.HasPrefix(line, long+"=") {
			return strings.TrimSpace(line[len(long)+1:]), true
		}

	}

	return "", false

}

func resolveInclude(requirementsFile string, include string) string {

	include = strings.TrimSpace(include)
	include = strings.Trim(include, "\"'")

	if strings.HasPrefix(include, "/") {
		return include
	}

	return filepath.Join(filepath.Dir(requirementsFile), include)

}

func stripHashOptions(line string) string {

	fields := strings.Fields(line)
	kept := make([]string, 0)
	skip := false

	for f := 0; f < len(fields); f++ {

		field := fields[f]

		if skip == true {
			skip = false
			continue
		}

		if field == "--hash" {
			skip = true
			continue
		}

		if strings.HasPrefix(field, "--hash=") {
			continue
		}

		kept = append(kept, field)

	}

	return strings.Join(kept, " ")

}

func parseRequirement(line string) (string, string) {

	line = strings.TrimSpace(line)

	if line == "" {
		return "", ""
	}

	if index := indexOutsideQuotes(line, ';'); index != -1 {
		line = strings.TrimSpace(line[0:index])
	}

	if line == "" {
		return "", ""
	}

	if index := strings.Index(line, " @ "); index > 0 {
		return stripExtras(line[0:index]), "*"
	}

	lower := strings.ToLower(line)

	if isDirectReference(lower) {
		return stripExtras(eggName(line)), "*"
	}

	return splitRequirementName(line)

}

func indexOutsideQuotes(line string, value rune) int {

	single := false
	double := false
	escaped := false

	for index, character := range line {

		if escaped == true {
			escaped = false
			continue
		}

		if character == '\\' {
			escaped = true
			continue
		}

		if character == '\'' && double == false {
			single = !single
			continue
		}

		if character == '"' && single == false {
			double = !double
			continue
		}

		if character == value && single == false && double == false {
			return index
		}

	}

	return -1

}

func isDirectReference(value string) bool {

	if value == "." || value == ".." {
		return true
	}

	prefixes := []string{
		"git+",
		"hg+",
		"svn+",
		"bzr+",
		"http://",
		"https://",
		"ftp://",
		"file:",
		"./",
		"../",
		"/",
	}

	for p := 0; p < len(prefixes); p++ {

		if strings.HasPrefix(value, prefixes[p]) {
			return true
		}

	}

	return false

}

func eggName(value string) string {

	markers := []string{"#egg=", "&egg=", "?egg="}

	for m := 0; m < len(markers); m++ {

		if index := strings.Index(value, markers[m]); index != -1 {

			name := value[index+len(markers[m]):]

			if end := strings.IndexAny(name, "&#"); end != -1 {
				name = name[0:end]
			}

			if decoded, err := url.QueryUnescape(name); err == nil {
				name = decoded
			}

			return strings.TrimSpace(name)

		}

	}

	return ""

}

func splitRequirementName(line string) (string, string) {

	index := len(line)

	for i, character := range line {

		if character == '[' ||
			character == '(' ||
			character == '<' ||
			character == '>' ||
			character == '=' ||
			character == '!' ||
			character == '~' ||
			unicode.IsSpace(character) {
			index = i
			break
		}

	}

	name := strings.TrimSpace(line[0:index])
	rest := strings.TrimSpace(line[index:])

	if strings.HasPrefix(rest, "[") {

		if end := strings.Index(rest, "]"); end != -1 {
			rest = strings.TrimSpace(rest[end+1:])
		}

	}

	return stripExtras(name), normalizeSpecifier(rest)

}

func stripExtras(name string) string {

	name = strings.TrimSpace(name)

	if index := strings.Index(name, "["); index != -1 {
		name = name[0:index]
	}

	return strings.TrimSpace(name)

}

func normalizeSpecifier(value string) string {

	value = strings.TrimSpace(value)

	if value == "" || value == "*" {
		return "*"
	}

	if strings.HasPrefix(value, "(") && strings.HasSuffix(value, ")") {
		value = strings.TrimSpace(value[1 : len(value)-1])
	}

	parts := strings.Split(value, ",")
	normalized := make([]string, 0)

	for p := 0; p < len(parts); p++ {

		part := strings.TrimSpace(parts[p])
		part = strings.TrimPrefix(part, "(")
		part = strings.TrimSuffix(part, ")")
		part = strings.TrimSpace(part)

		if part == "" {
			continue
		}

		normalized = append(normalized, normalizeOperator(part))

	}

	if len(normalized) == 0 {
		return "*"
	}

	return strings.Join(normalized, ", ")

}

func normalizeOperator(value string) string {

	if strings.HasPrefix(value, "===") {
		return "= " + strings.TrimSpace(value[3:])
	} else if strings.HasPrefix(value, "==") {
		return "= " + strings.TrimSpace(value[2:])
	} else if strings.HasPrefix(value, ">=") {
		return ">= " + strings.TrimSpace(value[2:])
	} else if strings.HasPrefix(value, "<=") {
		return "<= " + strings.TrimSpace(value[2:])
	} else if strings.HasPrefix(value, "!=") {
		return "!= " + strings.TrimSpace(value[2:])
	} else if strings.HasPrefix(value, "~=") {
		return "~= " + strings.TrimSpace(value[2:])
	} else if strings.HasPrefix(value, ">") {
		return "> " + strings.TrimSpace(value[1:])
	} else if strings.HasPrefix(value, "<") {
		return "< " + strings.TrimSpace(value[1:])
	}

	return value

}

func setManifestDatetime(file string, result *structs.Package) {

	if stat, err := os.Stat(file); err == nil {
		result.SetDatetime(stat.ModTime().Format(time.RFC3339))
	} else {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

}
