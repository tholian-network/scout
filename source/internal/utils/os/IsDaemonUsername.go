package os

func IsDaemonUsername(username string) bool {

	daemon_usernames := []string{
		"apache",
		"avahi",
		"backup",
		"bind",
		"colord",
		"daemon",
		"dovecot",
		"dovenull",
		"exim",
		"mail",
		"ftp",
		"flatpak",
		"http",
		"httpd",
		"git",
		"lp",
		"mysql",
		"nginx",
		"opendkim",
		"openvpn",
		"postfix",
		"postgres",
		"proxy",
		"samba",
		"saned",
		"sshd",
		"tor",
		"unbound",
		"www",
		"www-data",
	}

	for d := 0; d < len(daemon_usernames); d++ {

		other := daemon_usernames[d]

		if other == username {
			return true
		}

	}

	return false

}
