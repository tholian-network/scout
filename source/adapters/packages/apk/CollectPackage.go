package apk

import "tholian-core/structs"
import "os"
import "os/exec"

func CollectPackage(name string) structs.Package {

	var result structs.Package = structs.NewPackage("apk")

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("apk", "info", "--installed", "--all", name)
		buffer, err := cmd.Output()

		if err == nil {
			ParsePackage(string(buffer), &result)
		}

	}

	return result

}
