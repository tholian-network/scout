package dnf

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

			os.Setenv("TZ", "Europe/Greenwich")
			os.Setenv("LC_TIME", "en_US")

			arguments := []string{"--assumeyes", "upgrade", "--refresh"}
			arguments = append(arguments, names...)

			var stdout bytes.Buffer
			var stderr bytes.Buffer

			cmd := exec.Command("dnf", arguments...)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err1 := cmd.Run()

			if err1 == nil {
				updated = updates
			} else {
				stderr_message := strings.TrimSpace(string(stderr.Bytes()))
				err = errors.New("Could not update Packages: " + stderr_message)
			}

		}

	}

	return updated, err

}
