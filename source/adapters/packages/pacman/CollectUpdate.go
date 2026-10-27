package pacman

import "tholian-core/structs"
import "os"
import "os/exec"

func CollectUpdate(name string) structs.Update {

	os.Setenv("TZ", "Europe/Greenwich")
	os.Setenv("LC_TIME", "en_US")

	var result structs.Update = structs.NewUpdate("pacman")

	cmd := exec.Command("pacman", "-Si", "--noconfirm", name)
	buffer, err := cmd.Output()

	if err == nil {
		ParseUpdate(string(buffer), &result)
	}

	return result

}
