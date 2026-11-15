package pciconf

import "tholian-core/structs"
import "fmt"
import "os/exec"
import "strconv"
import "strings"

func CollectPCI() []structs.Device {

	var collected []structs.Device

	if SUPPORTED == true {

		cmd := exec.Command("pciconf", "-lv")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			var current *structs.Device

			names := make(map[string]string)
			ids := make(map[string]string)

			lines := strings.Split(strings.ReplaceAll(string(buffer), "\r\n", "\n"), "\n")

			for l := 0; l < len(lines); l++ {

				line := lines[l]
				trimmed := strings.TrimSpace(line)

				if trimmed == "" {
					continue
				}

				if strings.Contains(trimmed, "@pci") {

					if flushed, ok := flushDevice(current, names, ids); ok == true {
						collected = append(collected, flushed)
					}

					current = nil
					names = make(map[string]string)
					ids = make(map[string]string)

					device := structs.NewDevice("pci")
					current = &device

					fields := strings.Fields(trimmed)

					for f := 0; f < len(fields); f++ {

						chunks := strings.SplitN(fields[f], "=", 2)

						if len(chunks) == 2 {

							key := strings.TrimSpace(chunks[0])
							value := strings.TrimSpace(chunks[1])

							if strings.HasPrefix(value, "0x") {

								tmp_id := strings.ToLower(strings.TrimSpace(value))

								if tmp_num, err := strconv.ParseUint(tmp_id, 16, 64); err == nil {
									ids[key] = fmt.Sprintf("%04x", tmp_num)
								}

							}

						}

					}

				} else if current != nil && strings.Contains(trimmed, "=") {

					chunks := strings.SplitN(trimmed, "=", 2)

					if len(chunks) == 2 {

						key := strings.TrimSpace(chunks[0])
						value := strings.TrimSpace(chunks[1])

						if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") && len(value) >= 2 {
							value = value[1 : len(value)-1]
						}

						names[key] = value

					}

				}

			}

			if flushed, ok := flushDevice(current, names, ids); ok == true {
				collected = append(collected, flushed)
			}

		}

	}

	return collected

}
