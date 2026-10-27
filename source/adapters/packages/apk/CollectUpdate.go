package apk

import "tholian-core/structs"
import "os"
import "os/exec"
import "strings"

func CollectUpdate(name string) structs.Update {

	var result structs.Update = structs.NewUpdate("apk")

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		cmd := exec.Command("apk", "info", "--all", name)
		buffer, err := cmd.Output()

		if err == nil {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")
			blocks := make([]string, 0)
			block := make([]string, 0)

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if strings.HasPrefix(line, name) && strings.HasSuffix(line, "description:") {

					if len(block) > 0 {
						blocks = append(blocks, strings.TrimSpace(strings.Join(block, "\n")))
					}

					block = []string{line}

				} else {
					block = append(block, line)
				}

			}

			if len(block) > 0 {
				blocks = append(blocks, strings.TrimSpace(strings.Join(block, "\n")))
			}

			if len(blocks) > 0 {
				ParseUpdate(blocks[len(blocks)-1], &result)
			}

		}

	}

	return result

}
