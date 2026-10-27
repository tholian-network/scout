package pacman

import "tholian-core/structs"
import "bytes"
import "errors"
import "os/exec"
import "strings"

func Update(updates []structs.Update) ([]structs.Update, error) {

	var updated []structs.Update
	var err error = nil

	if SUPPORTED == true {

		cmd1 := exec.Command("pacman", "-Sy", "--noconfirm")
		err1 := cmd1.Run()

		if err1 == nil {

			// TODO: Use arguments :=?
			// TODO: -Sw should be -S
			var arguments []string = []string{"-S", "--noconfirm"}

			for u := 0; u < len(updates); u++ {
				arguments = append(arguments, updates[u].Name)
			}

			var stdout2 bytes.Buffer
			var stderr2 bytes.Buffer

			cmd2 := exec.Command("pacman", arguments...)
			cmd2.Dir = "/tmp"

			cmd2.Stdout = &stdout2
			cmd2.Stderr = &stderr2

			err2 := cmd2.Run()

			if err2 == nil {

				// block_suffix can be "Total Installed Size:" or "Total Download Size:"
				tmp2 := strings.TrimSpace(string(stdout2.Bytes()))
				block_prefix := "\n\nPackages "
				block_suffix := "\n\nTotal"

				index_prefix := strings.Index(tmp2, block_prefix)
				index_suffix := strings.Index(tmp2, block_suffix)

				if index_prefix != -1 && index_suffix != -1 && index_suffix > index_prefix {

					block := strings.TrimSpace(tmp2[index_prefix+len(block_prefix):index_suffix])

					if strings.HasPrefix(block, "(") && strings.Contains(block, ") ") {
						block = strings.TrimSpace(block[strings.Index(block, ") ")+2:])
					}

					lines := strings.Split(block, "\n")

					for l := 0; l < len(lines); l++ {

						packages := strings.Split(strings.TrimSpace(lines[l]), "  ")

						for p := 0; p < len(packages); p++ {

							name, version := splitShortName(packages[p])

							update := structs.NewUpdate("pacman")
							update.SetName(name)
							update.SetVersion(version)

							updated = append(updated, update)

						}

					}

				}

			} else {

				stderr_message := strings.TrimSpace(string(stderr2.Bytes()))
				err = errors.New("Could not synchronize Packages: " + stderr_message)

			}

		} else {
			err = errors.New("Could not synchronize Package Database")
		}

	}

	return updated, err

}
