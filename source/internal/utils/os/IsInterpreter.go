package os

import "os"
import "path/filepath"
import "strings"

func IsInterpreter(file string) bool {

	command := ""

	stat, err1 := os.Lstat(file)

	if err1 == nil {

		if stat.Mode()&os.ModeSymlink != 0 {

			real_file, err2 := filepath.EvalSymlinks(file)

			if err2 == nil {
				file = real_file
			}

		}

		buffer, err2 := os.ReadFile(file)

		if err2 == nil && len(buffer) > 0 {

			shebang := strings.TrimSpace(strings.Split(string(buffer), "\n")[0])

			if strings.HasPrefix(shebang, "#!") {
				command = ""
			} else {
				command = file
			}

		}

	}

	if command != "" {

		for repl_name := range interpreters {

			repl_command := ToProgram(repl_name)

			if repl_command != "" && repl_command == command {
				return true
			}

		}

	}

	return false

}
