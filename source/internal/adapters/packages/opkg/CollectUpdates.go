package opkg

import "tholian-core/structs"
import "os/exec"
import "strings"

func CollectUpdates() []structs.Update {

	var collected []structs.Update

	if SUPPORTED == true {

		architectures := make(map[string]string)

		buffer, err := readStatus()

		if err == nil && len(buffer) > 0 {
			architectures = collectArchitectures(string(buffer))
		}

		exec.Command("opkg", "update").Run()

		cmd := exec.Command("opkg", "list-upgradable")
		buffer, err = cmd.Output()

		if err == nil && len(buffer) > 0 {
			collected = parseUpgradable(string(buffer), architectures)
		}

	}

	return collected

}

func parseUpgradable(buffer string, architectures map[string]string) []structs.Update {

	var collected []structs.Update

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if line == "" {
			continue
		}

		chunks := strings.Split(line, " - ")

		if len(chunks) < 3 {
			continue
		}

		update := structs.NewUpdate("opkg")
		ParseUpdate(line, &update)

		if architecture, ok := architectures[update.Name]; ok && architecture != "" {
			update.SetArchitecture(architecture)
		}

		if update.Name != "" && update.Version.IsValid() {
			collected = append(collected, update)
		}

	}

	return collected

}
