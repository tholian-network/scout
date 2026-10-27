package apk

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "strconv"
import "strings"
import "time"

func parsePackageIndex(buffer string, result *structs.Package) {

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	var last_file string
	var last_folder string

	// F: sets the last folder (incrementally)
	// M: sets folder attributes
	// R: sets file inside last folder
	// a: sets file attributes
	// Z: compares file hashes

	// F: etc
	// F: etc/whatever
	// M: 0:0:644
	// R: program.cfg
	// a: 0:0:755
	// Z: Q1<sha1 hash>

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if strings.HasPrefix(line, "A:") {

			// Architecture (e.g. "A:x86_64")
			val := strings.TrimSpace(line[2:])

			if val != "" {
				result.SetArchitecture(val)
			}

		} else if strings.HasPrefix(line, "C:") {
			// Checksum: Do Nothing
		} else if strings.HasPrefix(line, "D:") {

			// Conflicts and Dependencies
			// (e.g. "D:!vlan so:libc.whatever-x86_64.so.1 package>=1.2.3-r4")
			chunks := strings.Split(strings.TrimSpace(line[2:]), " ")

			for c := 0; c < len(chunks); c++ {

				chunk := strings.TrimSpace(chunks[c])

				if chunk != "" {

					if strings.HasPrefix(chunk, "!") {

						if strings.HasPrefix(chunk, "!cmd:") {

							conflict := matchers.ToPackage(strings.TrimSpace(chunk[5:]))
							conflict.SetManager("apk")
							result.AddConflict(conflict)

						} else if strings.HasPrefix(chunk, "!so:") {

							conflict := matchers.ToPackage(strings.TrimSpace(chunk[4:]))
							conflict.SetManager("apk")
							result.AddConflict(conflict)

						} else if strings.HasPrefix(chunk, "!/") {

							conflict := matchers.NewPackage()
							conflict.SetName(strings.TrimSpace(chunk[1:]))
							result.AddConflict(conflict)

						} else {

							conflict := matchers.ToPackage(strings.TrimSpace(chunk[1:]))
							conflict.SetManager("apk")
							result.AddConflict(conflict)

						}

					} else {

						if strings.HasPrefix(chunk, "cmd:") {

							dependency := matchers.ToPackage(strings.TrimSpace(chunk[4:]))
							dependency.SetManager("apk")
							result.AddDependency(dependency)

						} else if strings.HasPrefix(chunk, "so:") {

							dependency := matchers.ToPackage(strings.TrimSpace(chunk[3:]))
							dependency.SetManager("apk")
							result.AddDependency(dependency)

						} else if strings.HasPrefix(chunk, "/") {

							dependency := matchers.NewPackage()
							dependency.SetName(chunk)
							dependency.SetManager("apk")
							result.AddDependency(dependency)

						} else {

							dependency := matchers.ToPackage(chunk)
							dependency.SetManager("apk")
							result.AddDependency(dependency)

						}

					}

				}

			}

		} else if strings.HasPrefix(line, "F:") {

			// Folder Path (e.g. "etc/whatever")
			val := strings.TrimSpace(line[2:])

			if val != "" {
				last_folder = "/" + val
			}

		} else if strings.HasPrefix(line, "I:") {
			// Installed Size: Do Nothing
		} else if strings.HasPrefix(line, "L:") {
			// License: Do Nothing
		} else if strings.HasPrefix(line, "M:") {

			// Folder ACL (uid:gid:chmod): Do Nothing

			if last_folder != "" {
				// Do Nothing
			}

		} else if strings.HasPrefix(line, "P:") {

			// Name (e.g. "package-name")
			val := strings.TrimSpace(line[2:])

			if val != "" {
				result.SetName(val)
			}

		} else if strings.HasPrefix(line, "R:") {

			// File Name (e.g. "program.cfg")
			val := strings.TrimSpace(line[2:])

			if last_folder != "" {

				last_file = strings.TrimSpace(last_folder + "/" + val)

				if strings.HasPrefix(last_file, "/") {
					result.AddFilesystem(last_file)
				}

			}

		} else if strings.HasPrefix(line, "S:") {
			// Compressed Size: Do Nothing
		} else if strings.HasPrefix(line, "T:") {
			// Description: Do Nothing
		} else if strings.HasPrefix(line, "U:") {

			// URL
			val := strings.TrimSpace(line[2:])

			if val != "" {
				result.SetURL(val)
			}

		} else if strings.HasPrefix(line, "V:") {

			// Version
			val := strings.TrimSpace(line[2:])

			if val != "" {
				result.SetVersion(val)
			}

		} else if strings.HasPrefix(line, "Z:") {

			// File Hash: Do Nothing

			if last_folder != "" && last_file != "" {
				// Do Nothing
			}

		} else if strings.HasPrefix(line, "a:") {

			// File ACL (uid:gid:chmod): Do Nothing

			if last_folder != "" && last_file != "" {
				// Do Nothing
			}

		} else if strings.HasPrefix(line, "c:") {
			// Commit Hash: Do Nothing
		} else if strings.HasPrefix(line, "f:") {
			// Broken Items: Do Nothing
		} else if strings.HasPrefix(line, "i:") {
			// Install If: Do Nothing
		} else if strings.HasPrefix(line, "k:") {
			// Provider Priority: Do Nothing
		} else if strings.HasPrefix(line, "m:") {

			// Maintainer (e.g. "Prename Surname <email@server.tld>")
			val := strings.TrimSpace(line[2:])

			if val != "" {
				result.AddMaintainer(types.ToMaintainer(val))
			}

		} else if strings.HasPrefix(line, "o:") {
			// Origin: Do Nothing
		} else if strings.HasPrefix(line, "p:") {

			// Provides
			// (e.g. "p:so:libc.whatever-x86_64.so.1 cmd:command=1.2.3-r4 package=1.2.3-r4")
			chunks := strings.Split(strings.TrimSpace(line[2:]), " ")

			for c := 0; c < len(chunks); c++ {

				chunk := strings.TrimSpace(chunks[c])

				if chunk != "" {

					if strings.HasPrefix(chunk, "cmd:") {

						provide := matchers.ToPackage(strings.TrimSpace(chunk[4:]))
						provide.SetManager("apk")
						result.AddProvide(provide)

					} else if strings.HasPrefix(chunk, "so:") {

						provide := matchers.ToPackage(strings.TrimSpace(chunk[3:]))
						provide.SetManager("apk")
						result.AddProvide(provide)

					} else if strings.HasPrefix(chunk, "/") {

						provide := matchers.NewPackage()
						provide.SetName(chunk)
						provide.SetManager("apk")
						result.AddProvide(provide)

					} else {

						provide := matchers.ToPackage(chunk)
						provide.SetManager("apk")
						result.AddProvide(provide)

					}

				}

			}

		} else if strings.HasPrefix(line, "q:") {
			// Priority: Do Nothing
		} else if strings.HasPrefix(line, "r:") {

			// Replaces
			// (e.g. "r:package1 package2=1.2.3-r4")
			chunks := strings.Split(strings.TrimSpace(line[2:]), " ")

			for c := 0; c < len(chunks); c++ {

				chunk := strings.TrimSpace(chunks[c])

				if chunk != "" {

					if strings.HasPrefix(chunk, "cmd:") {
						// Never happens
					} else if strings.HasPrefix(chunk, "so:") {
						// Never happens
					} else if strings.HasPrefix(chunk, "/") {
						// Never happens
					} else {

						replace := matchers.ToPackage(chunk)
						replace.SetManager("apk")
						result.AddReplace(replace)

					}

				}

			}

		} else if strings.HasPrefix(line, "s:") {
			// Repository Tag: Never Happens
		} else if strings.HasPrefix(line, "t:") {

			// Build Time in milliseconds
			val, err := strconv.ParseInt(strings.TrimSpace(line[2:]), 10, 64)

			if err == nil {
				result.SetDatetime(time.Unix(val, 0).Format(time.RFC3339))
			}

		}

	}

}
