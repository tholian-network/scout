package apk

import "tholian-core/matchers"
import "tholian-core/structs"
import "strings"
import "time"

func ParsePackage(buffer string, result *structs.Package) {

	blocks := strings.Split(strings.TrimSpace(buffer), "\n\n")

	if len(blocks) > 0 {

		name, version := toNameAndVersion(blocks[0][0:strings.Index(blocks[0], " ")])

		result.SetName(name)
		result.SetVersion(version)

		for b := 0; b < len(blocks); b++ {

			lines := strings.Split(strings.TrimSpace(blocks[b]), "\n")

			if len(lines) > 0 {

				first_line := lines[0]

				if strings.HasSuffix(first_line, " description:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " webpage:") {

					if len(lines) == 2 {
						result.SetURL(strings.TrimSpace(lines[1]))
					}

				} else if strings.HasSuffix(first_line, " installed size:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " depends on:") {

					for l := 1; l < len(lines); l++ {

						line := strings.TrimSpace(lines[l])

						if line != "" {

							if strings.HasPrefix(line, "!") {

								if strings.HasPrefix(line, "!cmd:") {

									conflict := matchers.ToPackage(strings.TrimSpace(line[5:]))
									conflict.SetManager("apk")
									result.AddConflict(conflict)

								} else if strings.HasPrefix(line, "!so:") {

									conflict := matchers.ToPackage(strings.TrimSpace(line[4:]))
									conflict.SetManager("apk")
									result.AddConflict(conflict)

								} else if strings.HasPrefix(line, "!/") {

									conflict := matchers.NewPackage()
									conflict.SetName(strings.TrimSpace(line[1:]))
									result.AddConflict(conflict)

								} else {

									conflict := matchers.ToPackage(strings.TrimSpace(line[1:]))
									conflict.SetManager("apk")
									result.AddConflict(conflict)

								}

							} else {

								if strings.HasPrefix(line, "cmd:") {

									dependency := matchers.ToPackage(strings.TrimSpace(line[4:]))
									dependency.SetManager("apk")
									result.AddDependency(dependency)

								} else if strings.HasPrefix(line, "so:") {

									dependency := matchers.ToPackage(strings.TrimSpace(line[3:]))
									dependency.SetManager("apk")
									result.AddDependency(dependency)

								} else if strings.HasPrefix(line, "/") {

									dependency := matchers.NewPackage()
									dependency.SetName(line)
									dependency.SetManager("apk")
									result.AddDependency(dependency)

								} else {

									dependency := matchers.ToPackage(line)
									dependency.SetManager("apk")
									result.AddDependency(dependency)

								}

							}

						}

					}

				} else if strings.HasSuffix(first_line, " provides:") {

					for l := 1; l < len(lines); l++ {

						line := strings.TrimSpace(lines[l])

						if line != "" {

							if strings.HasPrefix(line, "cmd:") {

								provide := matchers.ToPackage(strings.TrimSpace(line[4:]))
								provide.SetManager("apk")
								result.AddProvide(provide)

							} else if strings.HasPrefix(line, "so:") {

								provide := matchers.ToPackage(strings.TrimSpace(line[3:]))
								provide.SetManager("apk")
								result.AddProvide(provide)

							} else if strings.HasPrefix(line, "/") {

								provide := matchers.NewPackage()
								provide.SetName(line)
								provide.SetManager("apk")
								result.AddProvide(provide)

							} else {

								provide := matchers.ToPackage(line)
								provide.SetManager("apk")
								result.AddProvide(provide)

							}

						}

					}

				} else if strings.HasSuffix(first_line, " is required by:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " contains:") {

					for l := 1; l < len(lines); l++ {

						file := strings.TrimSpace(lines[l])

						if file != "" {

							if !strings.HasPrefix(file, "/") {
								file = "/" + file
							}

							result.AddFilesystem(file)

						}

					}

				} else if strings.HasSuffix(first_line, " triggers:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " has auto-install rule:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " affects auto-installation of:") {
					// Do Nothing
				} else if strings.HasSuffix(first_line, " replaces:") {

					for l := 1; l < len(lines); l++ {

						line := strings.TrimSpace(lines[l])

						if line != "" {

							if strings.HasPrefix(line, "cmd:") {
								// Never happens
							} else if strings.HasPrefix(line, "so:") {
								// Never happens
							} else if strings.HasPrefix(line, "/") {
								// Never happens
							} else {

								replace := matchers.ToPackage(line)
								replace.SetManager("apk")
								result.AddReplace(replace)

							}

						}

					}

				} else if strings.HasSuffix(first_line, " license:") {
					// Do Nothing
				}

			}

		}

	}

	if result.Datetime.IsValid() == false {
		result.SetDatetime(time.Now().Format(time.RFC3339))
	}

}
