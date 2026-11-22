package pip

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "strings"
import version_utils "tholian-scout/internal/utils/versions"

type version_bound struct {
	operator string
	version  string
}

func addDependency(result *structs.Package, name string, value string) {

	alternatives := splitAlternatives(value)

	if len(alternatives) > 1 {

		candidates := make([]matchers.Package, 0)

		for a := 0; a < len(alternatives); a++ {

			if version, ok := normalizeVersion(alternatives[a]); ok == true {
				candidates = append(candidates, newMatcher(name, version))
			}

		}

		if len(candidates) > 0 {
			result.AddUnresolved(matchers.Unresolved{Candidates: candidates})
			return
		}

		result.AddDependency(newMatcher(name, "any"))
		return

	}

	version, ok := normalizeVersion(value)

	if ok == false {
		version = "any"
	}

	result.AddDependency(newMatcher(name, version))

}

func normalizeVersion(value string) (string, bool) {

	value = strings.TrimSpace(value)

	if value == "" {
		return "", false
	}

	lower := strings.ToLower(value)

	if lower == "any" || lower == "all" || lower == "latest" || lower == "next" {
		return "", false
	}

	if index := strings.Index(value, " - "); index > 0 {

		if version, ok := version_utils.Normalize(strings.TrimSpace(value[0:index])); ok == true {
			return version_utils.Constraint(">=", version)
		}

		return "", false

	}

	tokens := tokenizeConstraint(value)
	lower_bounds := make([]version_bound, 0)
	upper_bounds := make([]version_bound, 0)
	exacts := make([]string, 0)

	for t := 0; t < len(tokens); t++ {

		token := tokens[t]
		operator := ""
		version := token

		switch token {
		case "===", "==", "!=", "~=", ">=", "<=", ">", "<", "=":

			operator = token

			if t+1 < len(tokens) {
				t++
				version = tokens[t]
			} else {
				continue
			}

		}

		normalized, ok := version_utils.Normalize(version)

		if ok == false {
			continue
		}

		switch operator {
		case "!=":
			continue
		case "~=":
			lower_bounds = append(lower_bounds, version_bound{">=", normalized})
		case ">=", ">":
			lower_bounds = append(lower_bounds, version_bound{operator, normalized})
		case "<=", "<":
			upper_bounds = append(upper_bounds, version_bound{operator, normalized})
		case "=", "==", "===":

			if version_utils.HasWildcard(version) == true {
				lower_bounds = append(lower_bounds, version_bound{">=", normalized})
			} else {
				exacts = append(exacts, normalized)
			}

		default:

			if version_utils.HasWildcard(version) == true {
				lower_bounds = append(lower_bounds, version_bound{">=", normalized})
			} else {
				exacts = append(exacts, normalized)
			}

		}

	}

	return chooseVersion(exacts, lower_bounds, upper_bounds)

}

func chooseVersion(exacts []string, lower_bounds []version_bound, upper_bounds []version_bound) (string, bool) {

	if len(exacts) > 0 {

		value := exacts[0]

		for e := 1; e < len(exacts); e++ {

			if exacts[e] != value {
				return "", false
			}

		}

		return version_utils.Constraint("=", value)

	}

	if len(lower_bounds) > 0 {

		best := lower_bounds[0]

		for b := 1; b < len(lower_bounds); b++ {

			compare := compareVersions(lower_bounds[b].version, best.version)

			if compare > 0 || (compare == 0 && lower_bounds[b].operator == ">" && best.operator == ">=") {
				best = lower_bounds[b]
			}

		}

		return version_utils.Constraint(best.operator, best.version)

	}

	if len(upper_bounds) > 0 {

		best := upper_bounds[0]

		for b := 1; b < len(upper_bounds); b++ {

			compare := compareVersions(upper_bounds[b].version, best.version)

			if compare < 0 || (compare == 0 && upper_bounds[b].operator == "<" && best.operator == "<=") {
				best = upper_bounds[b]
			}

		}

		return version_utils.Constraint(best.operator, best.version)

	}

	return "", false

}

func compareVersions(value_a string, value_b string) int {

	a := types.ToVersion(value_a)
	b := types.ToVersion(value_b)

	if a.IsAfter(b) {
		return 1
	} else if a.IsBefore(b) {
		return -1
	}

	return 0

}

func tokenizeConstraint(value string) []string {

	value = strings.ReplaceAll(value, ",", " ")

	replacer := strings.NewReplacer(
		"===", " === ",
		"==", " == ",
		"!=", " != ",
		"~=", " ~= ",
		">=", " >= ",
		"<=", " <= ",
		">", " > ",
		"<", " < ",
		"=", " = ",
	)

	return strings.Fields(replacer.Replace(value))

}

func splitAlternatives(value string) []string {

	value = strings.ReplaceAll(value, "||", "|")
	parts := strings.Split(value, "|")
	result := make([]string, 0)

	for p := 0; p < len(parts); p++ {

		part := strings.TrimSpace(parts[p])

		if part != "" {
			result = append(result, part)
		}

	}

	return result

}

func newMatcher(name string, version string) matchers.Package {

	dependency := matchers.NewPackage()

	dependency.SetName(name)
	dependency.SetVersion(version)
	dependency.SetArchitecture("any")
	dependency.SetManager("pip")
	dependency.SetVendor("pip")

	return dependency

}
