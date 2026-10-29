package procstat

import "strconv"
import "strings"

type file_info struct {
	Folder string
	Files  []string
}

func parseProcstatFiles(buffer string) map[uint]file_info {

	collected := make(map[uint]file_info)
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(buffer), "\r\n", "\n"), "\n")

	for l := 0; l < len(lines); l++ {

		fields := strings.Fields(lines[l])

		if fields[0] != "PID" && len(fields) >= 10 {

			pid, err := strconv.ParseUint(fields[0], 10, 64)

			if err == nil {

				entry := collected[uint(pid)]
				fd := fields[2]
				typ := fields[3]
				name := strings.Join(fields[9:], " ")

				if fd == "cwd" {
					entry.Folder = name
				} else if typ == "v" && strings.HasPrefix(name, "/") {
					entry.Files = append(entry.Files, name)
				}

				collected[uint(pid)] = entry

			}

		}

	}

	return collected

}
