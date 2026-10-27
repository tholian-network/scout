package dnf

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdates() []structs.Update {

	var collected []structs.Update

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		queryformat := []string{
			"Name : %{name}",
			"Architecture : %{arch}",
			"Version : %{evr}",
			"URL : %{url}",
		}

		cmd := exec.Command("dnf", "repoquery", "--upgrades", "--cacheonly", "--queryformat", "\\n\\n"+strings.Join(queryformat, "\\n"), "--noplugins")
		buffer, err := cmd.Output()

		if err == nil {

			blocks := strings.Split("\n\n"+strings.TrimSpace(string(buffer)), "\n\nName")

			for b := 0; b < len(blocks); b++ {

				block := strings.TrimSpace(blocks[b])

				if block != "" {

					update := structs.NewUpdate("dnf")
					ParseUpdate("Name "+block, &update)

					if update.Name != "" && update.Version.IsValid() {
						collected = append(collected, update)
					}

				}

			}

		}

	}

	return collected

}
