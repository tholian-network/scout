package dpkg

import "tholian-core/structs"
import "os"
import "strings"

func CollectPackage(name string) structs.Package {

	var result structs.Package = structs.NewPackage("dpkg")

	if SUPPORTED == true {

		buffer, err := os.ReadFile(status_file)

		if err == nil && len(buffer) > 0 {

			blocks := strings.Split(strings.ReplaceAll(string(buffer), "\r\n", "\n"), "\n\n")
			requested := stripArchitecture(strings.TrimSpace(name))

			for b := 0; b < len(blocks); b++ {

				block := strings.TrimSpace(blocks[b])

				if block == "" {
					continue
				}

				fields := parseControlFields(block)

				if strings.TrimSpace(fields["Status"]) != "install ok installed" {
					continue
				}

				if stripArchitecture(fields["Package"]) != requested {
					continue
				}

				ParsePackage(block, &result)

				if result.Name != "" && result.Version.IsValid() {
					attachFiles(&result, fields["Architecture"])
				}

				return result

			}

		}

	}

	return result

}
