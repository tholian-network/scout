package pkg

import "tholian-core/structs"
import "bytes"
import "errors"
import "os"
import "os/exec"
import "strings"

func Update(updates []structs.Update) ([]structs.Update, error) {

	var updated []structs.Update
	var err error = nil

	if SUPPORTED == true {

		names := make([]string, 0)

		for u := 0; u < len(updates); u++ {

			if updates[u].Name != "" {
				names = append(names, updates[u].Name)
			}

		}

		if len(names) > 0 {

			environment := append(os.Environ(), "LC_ALL=C", "ASSUME_ALWAYS_YES=yes")

			var stdout1 bytes.Buffer
			var stderr1 bytes.Buffer

			cmd1 := exec.Command("pkg", "update")
			cmd1.Env = environment
			cmd1.Stdout = &stdout1
			cmd1.Stderr = &stderr1

			err1 := cmd1.Run()

			if err1 == nil {

				arguments := []string{"upgrade", "-y"}
				arguments = append(arguments, names...)

				var stdout2 bytes.Buffer
				var stderr2 bytes.Buffer

				cmd2 := exec.Command("pkg", arguments...)
				cmd2.Env = environment
				cmd2.Stdout = &stdout2
				cmd2.Stderr = &stderr2

				err2 := cmd2.Run()

				if err2 == nil {

					updated = updates

				} else {

					stderr_message := strings.TrimSpace(string(stderr2.Bytes()))
					err = errors.New("Could not update Packages: " + stderr_message)

				}

			} else {

				stderr_message := strings.TrimSpace(string(stderr1.Bytes()))
				err = errors.New("Could not update Package Database: " + stderr_message)

			}

		}

	}

	return updated, err

}
