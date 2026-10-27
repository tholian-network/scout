package os

import "os"
import "path/filepath"
import "strings"

func toProgramArguments(value string) []string {

	var arguments []string

	if strings.HasPrefix(value, "#!") {

		tmp := strings.Split(strings.TrimSpace(value[2:]), " ")

		if len(tmp) > 1 {

			if tmp[0] == "/usr/bin/env" || tmp[0] == "/bin/env" {

				// #!/usr/bin/env python2 --flag
				if len(tmp) >= 2 && !strings.HasPrefix(tmp[1], "/") {

					resolved := ToProgram(tmp[1])

					if resolved != "" {
						arguments = append(arguments, resolved)
					} else {
						arguments = append(arguments, tmp[1])
					}

					for t := 2; t < len(tmp); t++ {
						arguments = append(arguments, tmp[t])
					}

				}

			} else if strings.HasPrefix(tmp[0], "/") {

				// #!/bin/bash --flag
				resolved, err1 := filepath.EvalSymlinks(tmp[0])

				if err1 == nil {
					arguments = append(arguments, resolved)
				} else {
					arguments = append(arguments, tmp[0])
				}

				if len(tmp) > 1 {

					for t := 1; t < len(tmp); t++ {
						arguments = append(arguments, tmp[t])
					}

				}

			}

		} else if len(tmp) == 1 {

			if strings.HasPrefix(tmp[0], "/") {

				// #!/bin/bash
				resolved, err1 := filepath.EvalSymlinks(tmp[0])

				if err1 == nil {
					arguments = append(arguments, resolved)
				}

			} else {

				// #!bash
				resolved := ToProgram(tmp[0])

				if resolved != "" {
					arguments = append(arguments, resolved)
				} else {
					arguments = append(arguments, tmp[0])
				}

			}

		}

	}

	return arguments

}

func ToInterpreter(file string) (string, []string) {

	var command string
	var arguments []string

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

				arguments = toProgramArguments(shebang)

				if len(arguments) > 0 {
					command = arguments[0]
				}

			} else {
				arguments = []string{file}
				command = file
			}

		}

	}

	if command != "" {

		for repl_name, repl_args := range interpreters {

			repl_command := ToProgram(repl_name)

			if repl_command != "" && repl_command == command {

				if len(arguments) == 0 {

					for a := 0; a < len(repl_args); a++ {

						argument := repl_args[a]

						if argument == "%file%" {
							arguments = append(arguments, file)
						} else {
							arguments = append(arguments, argument)
						}

					}

				} else {
					arguments = append(arguments, file)
				}

				break

			}

		}

	}

	return command, arguments

}
