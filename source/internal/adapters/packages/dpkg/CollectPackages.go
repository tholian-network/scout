package dpkg

import "tholian-core/structs"
import "os"
import "strings"

const status_file = "/var/lib/dpkg/status"
const info_directory = "/var/lib/dpkg/info"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true {

		buffer, err := os.ReadFile(status_file)

		if err == nil && len(buffer) > 0 {

			blocks := strings.Split(strings.ReplaceAll(string(buffer), "\r\n", "\n"), "\n\n")

			for b := 0; b < len(blocks); b++ {

				block := strings.TrimSpace(blocks[b])

				if block == "" {
					continue
				}

				fields := parseControlFields(block)

				if strings.TrimSpace(fields["Status"]) != "install ok installed" {
					continue
				}

				pkg := structs.NewPackage("dpkg")
				ParsePackage(block, &pkg)

				if pkg.Name != "" && pkg.Version.IsValid() {

					attachFiles(&pkg, fields["Architecture"])

					collected = append(collected, pkg)

				}

			}

		}

		for c := 0; c < len(collected); c++ {
			collected[c].ResolveDependencies(collected)
		}

	}

	return collected

}

func attachFiles(pkg *structs.Package, architecture string) {

	candidates := make([]string, 0)

	if pkg.Name != "" {
		candidates = append(candidates, pkg.Name)
	}

	if pkg.Name != "" && architecture != "" {
		candidates = append(candidates, pkg.Name+":"+architecture)
	}

	for c := 0; c < len(candidates); c++ {

		buffer, err := os.ReadFile(info_directory + "/" + candidates[c] + ".list")

		if err == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

			for l := 0; l < len(lines); l++ {

				file := strings.TrimSpace(lines[l])

				if strings.HasPrefix(file, "/") {
					pkg.AddFilesystem(file)
				}

			}

			return

		}

	}

}
