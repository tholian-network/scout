package incidents

import "tholian-core/structs"
import "sort"
import "strings"

// Identity returns a stable key for an incident that can be used to merge the
// same event reported by different collectors. Evidence is intentionally
// excluded because journald and syslog may attach different evidence for the
// same event.
func ToIdentity(value structs.Incident) string {

	values := []string{
		value.System,
		value.Type,
		value.Datetime.String(),
	}

	for _, connection := range value.Connections {
		values = append(values, "c:"+connection.Hash())
	}

	for _, distribution := range value.Distributions {
		values = append(values, "d:"+distribution.Hash())
	}

	for _, network := range value.Networks {
		values = append(values, "n:"+network.Hash())
	}

	for _, pkg := range value.Packages {
		values = append(values, "p:"+pkg.Hash())
	}

	for _, product := range value.Products {
		values = append(values, "r:"+product.Hash())
	}

	for _, program := range value.Programs {
		values = append(values, "g:"+program.Hash())
	}

	for _, user := range value.Users {
		values = append(values, "u:"+user.Hash())
	}

	for _, vulnerability := range value.Vulnerabilities {
		values = append(values, "v:"+vulnerability.Hash())
	}

	for _, weakness := range value.Weaknesses {
		values = append(values, "w:"+weakness.Hash())
	}

	sort.Strings(values)

	return strings.Join(values, "|")

}
