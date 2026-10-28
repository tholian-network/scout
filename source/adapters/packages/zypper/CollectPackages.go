package zypper

import "tholian-core/structs"
import "tholian-scout/adapters/packages/rpm"
import "os"
import "os/exec"
import "strings"
import "time"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		rpm_packages := rpm.CollectPackages()
		rpm_index := make(map[string]structs.Package)

		for r := 0; r < len(rpm_packages); r++ {
			key := rpm_packages[r].Name + "." + rpm_packages[r].Architecture.String()
			rpm_index[key] = rpm_packages[r]
		}

		cmd := exec.Command("zypper", "packages", "--installed-only")
		buffer, err := cmd.Output()

		if err == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")
			start := -1

			for l := 0; l < len(lines)-1; l++ {

				check1 := strings.TrimSpace(lines[l])
				check2 := strings.TrimSpace(lines[l+1])

				if strings.Contains(check1, "|") &&
					strings.HasPrefix(check2, "---") && strings.HasSuffix(check2, "---") {
					start = l + 2
					break
				}

			}

			if start != -1 {

				index := make(map[string]bool)

				for l := start; l < len(lines); l++ {

					line := strings.TrimSpace(lines[l])

					if line != "" {

						if strings.Contains(line, "---") {

							// Do Nothing

						} else if strings.Contains(line, "|") {

							chunks := strings.Split(line, "|")

							if len(chunks) == 5 {

								status := strings.TrimSpace(chunks[0])
								name := strings.TrimSpace(chunks[2])
								version := strings.TrimSpace(chunks[3])
								architecture := strings.TrimSpace(chunks[4])
								key := name + "." + architecture

								if (status == "i" || status == "i+") && index[key] == false {

									index[key] = true

									pkg := structs.NewPackage("zypper")
									pkg.SetName(name)
									pkg.SetVersion(version)
									pkg.SetArchitecture(architecture)

									if other, ok := rpm_index[key]; ok {
										pkg.URL = other.URL
										pkg.Vendor = other.Vendor
										pkg.Datetime = other.Datetime
										pkg.Maintainers = other.Maintainers
										pkg.Filesystem = other.Filesystem
										pkg.Conflicts = other.Conflicts
										pkg.Dependencies = other.Dependencies
										pkg.Provides = other.Provides
										pkg.Replaces = other.Replaces
										pkg.Unresolved = other.Unresolved
									}

									if pkg.Datetime.IsValid() == false {
										pkg.SetDatetime(time.Now().Format(time.RFC3339))
									}

									if pkg.Name != "" && pkg.Version.IsValid() {
										collected = append(collected, pkg)
									}

								}

							}

						}

					}

				}

			}

		}

	}

	return collected

}
