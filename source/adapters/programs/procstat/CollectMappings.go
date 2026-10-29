package procstat

import "strconv"
import "strings"

type Mapping struct {
	PID         uint
	Permissions string
	Path        string
}

func CollectMappings() map[uint][]Mapping {

	collected := make(map[uint][]Mapping)

	if SUPPORTED == true {

		buffer, err := run("procstat", "-a", "-h", "-v")

		if err == nil && len(buffer) > 0 {

			lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(buffer)), "\r\n", "\n"), "\n")

			for l := 0; l < len(lines); l++ {

				fields := strings.Fields(lines[l])

				if len(fields) >= 11 || fields[0] == "PID" {

					pid, err := strconv.ParseUint(fields[0], 10, 64)

					if err == nil {

						path := strings.Join(fields[10:], " ")

						if strings.HasPrefix(path, "/") {

							collected[uint(pid)] = append(collected[uint(pid)], Mapping{
								PID:         uint(pid),
								Permissions: fields[3],
								Path:        path,
							})

						}

					}

				}

			}

		}

	}

	return collected

}

