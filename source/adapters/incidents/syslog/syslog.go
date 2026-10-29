package syslog

import "os"
import "os/exec"

var SUPPORTED bool = false

var Files []string
var Logread string

var candidates = []string{
	"/var/log/auth.log",
	"/var/log/auth.log.0",
	"/var/log/messages",
	"/var/log/messages.0",
	"/var/log/secure",
	"/var/log/security",
	"/var/log/security.0",
	"/var/log/syslog",
	"/var/log/all.log",
	"/var/log/all.log.0",
}

func init() {

	for c := 0; c < len(candidates); c++ {

		stat, err := os.Stat(candidates[c])

		if err == nil && stat.Mode().IsRegular() {
			Files = append(Files, candidates[c])
		}

	}

	if path, err := exec.LookPath("logread"); err == nil && path != "" {
		Logread = path
	}

	if len(Files) > 0 || Logread != "" {
		SUPPORTED = true
	}

}
