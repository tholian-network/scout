package opkg

import "tholian-core/structs"
import "tholian-core/types"
import "os"
import "strings"

const status_file = "/usr/lib/opkg/status"
const info_directory = "/usr/lib/opkg/info"

func readStatus() ([]byte, error) {
	return os.ReadFile(status_file)
}

func parseStatus(buffer string) []structs.Package {

	var collected []structs.Package

	blocks := strings.Split(strings.ReplaceAll(string(buffer), "\r\n", "\n"), "\n\n")

	for b := 0; b < len(blocks); b++ {

		block := strings.TrimSpace(blocks[b])

		if block == "" {
			continue
		}

		fields := parseControlFields(block)

		if isInstalledStatus(fields["Status"]) == false {
			continue
		}

		pkg := structs.NewPackage("opkg")
		ParsePackage(block, &pkg)

		if pkg.Name != "" && pkg.Version.IsValid() {
			collected = append(collected, pkg)
		}

	}

	for c := 0; c < len(collected); c++ {
		collected[c].ResolveDependencies(collected)
	}

	return collected

}

func collectArchitectures(buffer string) map[string]string {

	architectures := make(map[string]string)
	blocks := strings.Split(strings.ReplaceAll(string(buffer), "\r\n", "\n"), "\n\n")

	for b := 0; b < len(blocks); b++ {

		block := strings.TrimSpace(blocks[b])

		if block == "" {
			continue
		}

		fields := parseControlFields(block)

		if isInstalledStatus(fields["Status"]) == false {
			continue
		}

		name := strings.TrimSpace(fields["Package"])

		if name != "" {
			architectures[name] = types.ToArchitecture(fields["Architecture"]).String()
		}

	}

	return architectures

}

func attachFiles(pkg *structs.Package) {

	if pkg.Name == "" {
		return
	}

	buffer, err := os.ReadFile(info_directory + "/" + pkg.Name + ".list")

	if err == nil {

		lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

		for l := 0; l < len(lines); l++ {

			file := strings.TrimSpace(lines[l])

			if file != "" {
				pkg.AddFilesystem(file)
			}

		}

	}

}
