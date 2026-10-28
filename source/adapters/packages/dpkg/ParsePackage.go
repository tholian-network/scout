package dpkg

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "strings"
import "time"

func ParsePackage(buffer string, result *structs.Package) {

	fields := parseControlFields(buffer)

	result.SetManager("dpkg")

	if value, ok := fields["Package"]; ok {
		result.SetName(stripArchitecture(value))
	}

	if value, ok := fields["Version"]; ok {
		result.SetVersion(value)
	}

	if value, ok := fields["Architecture"]; ok {
		result.SetArchitecture(value)
	}

	if value, ok := fields["Homepage"]; ok {
		result.SetURL(value)
	}

	if value, ok := fields["Maintainer"]; ok {
		result.AddMaintainer(types.ToMaintainer(value))
	}

	if value, ok := fields["Original-Maintainer"]; ok {
		result.AddMaintainer(types.ToMaintainer(value))
	}

	if value, ok := fields["Depends"]; ok {

		relations := make([]string, 0)
		chunks := strings.Split(value, ",")

		for c := 0; c < len(chunks); c++ {

			chunk := strings.TrimSpace(chunks[c])

			if chunk != "" {
				relations = append(relations, chunk)
			}

		}

		for r := 0; r < len(relations); r++ {

			relation := relations[r]

			if strings.Contains(relation, " | ") {

				candidates := strings.Split(relation, " | ")

				for c := 0; c < len(candidates); c++ {

					candidate := strings.TrimSpace(candidates[c])
					name := candidate
					rest := ""
					index := strings.IndexAny(candidate, " (")

					if index > 0 {
						name = candidate[0:index]
						rest = candidate[index:]
					}

					candidates[c] = stripArchitecture(name) + rest

				}

				relation = strings.Join(candidates, " | ")

			} else {

				relation = strings.TrimSpace(relation)
				name := relation
				rest := ""
				index := strings.IndexAny(relation, " (")

				if index > 0 {
					name = relation[0:index]
					rest = relation[index:]
				}

				relation = stripArchitecture(name) + rest

			}

			if strings.Contains(relation, " | ") {

				result.Unresolved = append(result.Unresolved, matchers.ToUnresolved(relation))

			} else {

				dependency := matchers.ToPackage(relation)
				dependency.SetManager("dpkg")
				result.AddDependency(dependency)

			}

		}

	}

	if value, ok := fields["Provides"]; ok {

		relations := make([]string, 0)
		chunks := strings.Split(value, ",")

		for c := 0; c < len(chunks); c++ {

			chunk := strings.TrimSpace(chunks[c])

			if chunk != "" {
				relations = append(relations, chunk)
			}

		}

		for r := 0; r < len(relations); r++ {

			relation := relations[r]

			if strings.Contains(relation, " | ") {

				candidates := strings.Split(relation, " | ")

				for c := 0; c < len(candidates); c++ {

					candidate := strings.TrimSpace(candidates[c])
					name := candidate
					rest := ""
					index := strings.IndexAny(candidate, " (")

					if index > 0 {
						name = candidate[0:index]
						rest = candidate[index:]
					}

					candidates[c] = stripArchitecture(name) + rest

				}

				relation = strings.Join(candidates, " | ")

			} else {

				relation = strings.TrimSpace(relation)
				name := relation
				rest := ""
				index := strings.IndexAny(relation, " (")

				if index > 0 {
					name = relation[0:index]
					rest = relation[index:]
				}

				relation = stripArchitecture(name) + rest

			}

			provide := matchers.ToPackage(relation)
			provide.SetManager("dpkg")
			result.AddProvide(provide)

		}

	}

	if value, ok := fields["Replaces"]; ok {

		relations := make([]string, 0)
		chunks := strings.Split(value, ",")

		for c := 0; c < len(chunks); c++ {

			chunk := strings.TrimSpace(chunks[c])

			if chunk != "" {
				relations = append(relations, chunk)
			}

		}

		for r := 0; r < len(relations); r++ {

			relation := relations[r]

			if strings.Contains(relation, " | ") {

				candidates := strings.Split(relation, " | ")

				for c := 0; c < len(candidates); c++ {

					candidate := strings.TrimSpace(candidates[c])
					name := candidate
					rest := ""
					index := strings.IndexAny(candidate, " (")

					if index > 0 {
						name = candidate[0:index]
						rest = candidate[index:]
					}

					candidates[c] = stripArchitecture(name) + rest

				}

				relation = strings.Join(candidates, " | ")

			} else {

				relation = strings.TrimSpace(relation)
				name := relation
				rest := ""
				index := strings.IndexAny(relation, " (")

				if index > 0 {
					name = relation[0:index]
					rest = relation[index:]
				}

				relation = stripArchitecture(name) + rest

			}

			replaced := matchers.ToPackage(relation)
			replaced.SetManager("dpkg")
			result.AddReplace(replaced)

		}

	}

	if value, ok := fields["Conflicts"]; ok {

		relations := make([]string, 0)
		chunks := strings.Split(value, ",")

		for c := 0; c < len(chunks); c++ {

			chunk := strings.TrimSpace(chunks[c])

			if chunk != "" {
				relations = append(relations, chunk)
			}

		}

		for r := 0; r < len(relations); r++ {

			relation := relations[r]

			if strings.Contains(relation, " | ") {

				candidates := strings.Split(relation, " | ")

				for c := 0; c < len(candidates); c++ {

					candidate := strings.TrimSpace(candidates[c])
					name := candidate
					rest := ""
					index := strings.IndexAny(candidate, " (")

					if index > 0 {
						name = candidate[0:index]
						rest = candidate[index:]
					}

					candidates[c] = stripArchitecture(name) + rest

				}

				relation = strings.Join(candidates, " | ")

			} else {

				relation = strings.TrimSpace(relation)
				name := relation
				rest := ""
				index := strings.IndexAny(relation, " (")

				if index > 0 {
					name = relation[0:index]
					rest = relation[index:]
				}

				relation = stripArchitecture(name) + rest

			}

			conflict := matchers.ToPackage(relation)
			conflict.SetManager("dpkg")
			result.AddConflict(conflict)

		}

	}

	if value, ok := fields["Breaks"]; ok {

		relations := make([]string, 0)
		chunks := strings.Split(value, ",")

		for c := 0; c < len(chunks); c++ {

			chunk := strings.TrimSpace(chunks[c])

			if chunk != "" {
				relations = append(relations, chunk)
			}

		}

		for r := 0; r < len(relations); r++ {

			relation := relations[r]

			if strings.Contains(relation, " | ") {

				candidates := strings.Split(relation, " | ")

				for c := 0; c < len(candidates); c++ {

					candidate := strings.TrimSpace(candidates[c])
					name := candidate
					rest := ""
					index := strings.IndexAny(candidate, " (")

					if index > 0 {
						name = candidate[0:index]
						rest = candidate[index:]
					}

					candidates[c] = stripArchitecture(name) + rest

				}

				relation = strings.Join(candidates, " | ")

			} else {

				relation = strings.TrimSpace(relation)
				name := relation
				rest := ""
				index := strings.IndexAny(relation, " (")

				if index > 0 {
					name = relation[0:index]
					rest = relation[index:]
				}

				relation = stripArchitecture(name) + rest

			}

			conflict := matchers.ToPackage(relation)
			conflict.SetManager("dpkg")
			result.AddConflict(conflict)

		}

	}

	if result.Datetime.IsValid() == false {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

}
