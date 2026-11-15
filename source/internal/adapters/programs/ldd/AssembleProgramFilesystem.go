package ldd

import "tholian-core/structs"
import "os/exec"
import "strings"

func AssembleProgramFilesystem(result *structs.Program) {

	if result.Command != "" && strings.HasPrefix(result.Command, "/") {

		cmd := exec.Command("ldd", result.Command)
		buffer, err := cmd.Output()

		if err == nil {

			lines := strings.Split(string(buffer), "\n")

			for l := 0; l < len(lines); l++ {

				line := strings.TrimSpace(lines[l])

				if strings.Contains(line, " => ") && strings.Contains(line, "(0x") && strings.HasSuffix(line, ")") {

					// unresolved := strings.TrimSpace(line[0:strings.Index(line, " => ")])
					resolved := strings.TrimSpace(line[strings.Index(line, " => ")+4:])

					if strings.Contains(resolved, "(0x") && strings.HasSuffix(resolved, ")") {
						resolved = strings.TrimSpace(resolved[0:strings.Index(resolved, "(0x")])
					}

					if strings.HasPrefix(resolved, "/") {
						result.AddFilesystem(resolved)
					}

				}

			}

		}

	}

}
