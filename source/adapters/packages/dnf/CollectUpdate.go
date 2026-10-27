package dnf

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdate(name string) structs.Update {

	var result structs.Update = structs.NewUpdate("dnf")

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		queryformat := []string{
			"Name : %{name}",
			"Architecture : %{arch}",
			"Version : %{evr}",
			"URL : %{url}",
		}

		cmd := exec.Command("dnf", "repoquery", "--upgrades", "--cacheonly", "--queryformat", "\\n\\n"+strings.Join(queryformat, "\\n"), "--noplugins", name)
		buffer, err := cmd.Output()

		if err == nil {
			ParseUpdate(string(buffer), &result)
		}

	}

	return result

}
