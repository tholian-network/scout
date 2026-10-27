package systemd

import "tholian-core/structs"

// import "tholian-scout/adapters/incidents/systemd/auth"
import "tholian-scout/adapters/incidents/systemd/dovecot"
import "tholian-scout/adapters/incidents/systemd/gogs"
import "tholian-scout/adapters/incidents/systemd/pam"
import "tholian-scout/adapters/incidents/systemd/postfix"
import "tholian-scout/adapters/incidents/systemd/sshd"
import "tholian-scout/adapters/incidents/systemd/sudo"

var Classifiers map[string]func(string, string, string, string, string) structs.Incident = map[string]func(string, string, string, string, string) structs.Incident{
	// "auth":                auth.Classify,
	// "auth-worker":         auth.Classify,
	"dovecot":             dovecot.Classify,
	"gogs":                gogs.Classify,
	"pam":                 pam.Classify,
	"ssh":                 sshd.Classify,
	"sshd":                sshd.Classify,
	"su":                  sudo.Classify,
	"sudo":                sudo.Classify,
	"postfix":             postfix.Classify,
	"postfix/anvil":       postfix.Classify,
	"postfix/bounce":      postfix.Classify,
	"postfix/cleanup":     postfix.Classify,
	"postfix/lmtp":        postfix.Classify,
	"postfix/master":      postfix.Classify,
	"postfix/qmgr":        postfix.Classify,
	"postfix/scache":      postfix.Classify,
	"postfix/smtpd":       postfix.Classify,
	"postfix/smtps/smtpd": postfix.Classify,
	"postfix/pickup":      postfix.Classify,
	"postfix/postmap":     postfix.Classify,
}
