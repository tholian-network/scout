package dnf

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectPackage(name string) structs.Package {

	var result structs.Package = structs.NewPackage("dnf")

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		queryformat := []string{
			"Name : %{name}",
			"Architecture : %{arch}",
			"Version : %{evr}",
			"Buildtime : %{buildtime}",
			"URL : %{url}",
			"Provides : %{provides}",
			"Requires : %{requires}",
			"Obsoletes : %{obsoletes}",
			"Conflicts : %{conflicts}",
			"Files : %{files}",
		}

		cmd := exec.Command("dnf", "repoquery", "--installed", "--cacheonly", "--queryformat", "\\n\\n"+strings.Join(queryformat, "\\n"), "--noplugins", name)
		buffer, err := cmd.Output()

		if err == nil {
			ParsePackage(string(buffer), &result)
		}

	}

	return result

}
