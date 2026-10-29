package proc

import "tholian-core/structs"
import "os"
import "os/exec"
import "path/filepath"
import "strings"

func Init(console *structs.Console) bool {

	var result bool
	var name string = "tholian-guard"
	var kernel string

	cmd := exec.Command("uname", "-s")
	buffer, err0 := cmd.Output()

	if err0 == nil {
		kernel = strings.TrimSpace(string(buffer))
	}

	exe, err1 := os.Executable()

	if err1 == nil {
		name = filepath.Base(exe)
	} else {
		name = os.Args[0]
	}

	if kernel == "FreeBSD" {

		if os.Geteuid() == 0 {
			return true
		}

		console.Error("")
		console.Error(name + ": Please execute \"" + name + "\" as \"root\" user")
		console.Error("")

		return false

	}

	_, err2 := os.Stat("/proc")

	if err2 == nil {

		if os.Getenv("USER") == "root" {
			result = true
		} else {

			console.Error("")
			console.Error(name + ": Please execute \"" + name + "\" as \"root\" user")
			console.Error("")

		}

	} else {

		if os.Getenv("USER") == "root" {

			err3 := os.Mkdir("/proc", os.ModePerm)

			if err3 == nil {

				if kernel == "BSD" {

					cmd := exec.Command("mount", "-t", "procfs", "proc", "/proc")
					_, err4 := cmd.Output()

					if err4 == nil {
						result = true
					} else {

						console.Error(name + ": Cannot mount procfs to /proc")
						console.Error(name + ": Please execute this before running \"" + name + "\" again:")
						console.Error(name + ":     mount -t procfs proc /proc;")

					}

				} else if kernel == "Linux" {

					cmd := exec.Command("mount", "-t", "proc", "proc", "/proc")
					_, err4 := cmd.Output()

					if err4 == nil {
						result = true
					} else {

						console.Error(name + ": Cannot mount proc to /proc")
						console.Error(name + ": Please execute this before running \"" + name + "\" again:")
						console.Error(name + ":     mount -t proc proc /proc;")

					}

				} else {
					console.Error(name + ": Not a POSIX Kernel!?")
				}

			} else {

				if kernel == "BSD" {

					console.Error(name + ": Cannot mount procfs to /proc")
					console.Error(name + ": Please execute this before running \"" + name + "\" again:")
					console.Error(name + ":     mkdir /proc;")
					console.Error(name + ":     mount -t procfs proc /proc;")

				} else if kernel == "Linux" {

					console.Error(name + ": Cannot mount proc to /proc")
					console.Error(name + ": Please execute this before running \"" + name + "\" again:")
					console.Error(name + ":     mkdir /proc;")
					console.Error(name + ":     mount -t proc proc /proc;")

				} else {
					console.Error(name + ": Not a POSIX Kernel!?")
				}

			}

		} else {

			console.Error("")
			console.Error(name + ": Please execute \"" + name + "\" as \"root\" user")
			console.Error("")

		}

	}

	return result

}
