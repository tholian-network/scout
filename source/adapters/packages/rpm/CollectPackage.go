package rpm

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectPackage(name string) structs.Package {

	var result structs.Package = structs.NewPackage("rpm")

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

		cmd := exec.Command("rpm", "-q", "--queryformat", "\\n\\n"+strings.Join(queryformat, "\\n"), "--noplugins", name)
		buffer, err := cmd.Output()

		if err == nil {
			ParsePackage(string(buffer), &result)
		}

	}

	return result

}
