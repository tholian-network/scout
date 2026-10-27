package rpm

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectPackages() []structs.Package {

	var collected []structs.Package

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		queryformat := []string{
			"Name : %{name}",
			"Architecture : %{arch}",
			"Version : %{evr}",
			"Buildtime : %{buildtime}",
			"URL : %{url}",
			"Vendor : %{vendor}",
			"Provides : [\n\t%{provides}]",
			"Requires : [\n\t%{requires}]",
			"Obsoletes : [\n\t%{obsoletes}]",
			"Conflicts : [\n\t%{conflicts}]",
			"Files : [\n\t%{filenames}]",
		}

		cmd := exec.Command("rpm", "-qa", "--queryformat", "\\n\\n"+strings.Join(queryformat, "\\n"), "--noplugins")
		buffer, err := cmd.Output()

		if err == nil {

			blocks := strings.Split("\n\n"+strings.TrimSpace(string(buffer)), "\n\nName")

			for b := 0; b < len(blocks); b++ {

				block := strings.TrimSpace(blocks[b])

				if block != "" {

					pkg := structs.NewPackage("rpm")
					ParsePackage("Name "+block, &pkg)

					if pkg.Name != "" && pkg.Version.IsValid() {
						collected = append(collected, pkg)
					}

				}

			}

		}

	}

	return collected

}
