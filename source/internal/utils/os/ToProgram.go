package os

import "os"
import "path/filepath"
import "strings"

func ToProgram(name string) string {

	if name == "systemd" || strings.HasPrefix(name, "systemd-") {

		stat, err1 := os.Stat("/usr/lib/systemd/" + name)

		if err1 == nil && !stat.IsDir() {

			resolved, err2 := filepath.EvalSymlinks("/usr/lib/systemd/" + name)

			if err2 == nil {
				return resolved
			}

		}

	} else if strings.HasPrefix(name, "libvbox") || strings.HasPrefix(name, "vbox") || strings.HasPrefix(name, "VBox") || strings.HasPrefix(name, "VirtualBox") {

		stat, err1 := os.Stat("/usr/lib/virtualbox/" + name)

		if err1 == nil && !stat.IsDir() {

			resolved, err2 := filepath.EvalSymlinks("/usr/lib/virtualbox/" + name)

			if err2 == nil {
				return resolved
			}

		}

	} else if name == "polkitd" || strings.HasPrefix(name, "polkit-") {

		stat, err1 := os.Stat("/usr/lib/polkit-1/" + name)

		if err1 == nil && !stat.IsDir() {

			resolved, err2 := filepath.EvalSymlinks("/usr/lib/polkit-1/" + name)

			if err2 == nil {
				return resolved
			}

		}

	} else if name == "sftp-server" || strings.HasPrefix(name, "ssh-") {

		stat, err1 := os.Stat("/usr/lib/ssh/" + name)

		if err1 == nil && !stat.IsDir() {

			resolved, err2 := filepath.EvalSymlinks("/usr/lib/ssh/" + name)

			if err2 == nil {
				return resolved
			}

		}

	} else if name == "udisksd" {

		stat, err1 := os.Stat("/usr/lib/udisks2/udisksd")

		if err1 == nil && !stat.IsDir() {
			return "/usr/lib/udisks2/udisksd"
		}

	}

	paths := []string{
		"/sbin",
		"/bin",
		"/lib",
		"/usr/local/sbin",
		"/usr/local/bin",
		"/usr/local/lib",
		"/usr/sbin",
		"/usr/bin",
		"/usr/lib",
		"/var/lib/flatpak/exports/bin",
		"/usr/lib/jvm/default/bin",
		"/usr/bin/site_perl",
		"/usr/bin/vendor_perl",
		"/usr/bin/core_perl",
	}

	for p := 0; p < len(paths); p++ {

		stat, err1 := os.Stat(paths[p] + "/" + name)

		if err1 == nil && !stat.IsDir() {

			resolved, err2 := filepath.EvalSymlinks(paths[p] + "/" + name)

			if err2 == nil {
				return resolved
			}

		}

	}

	return ""

}
