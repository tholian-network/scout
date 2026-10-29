package classifiers

import "tholian-core/structs"

import "tholian-scout/internal/classifiers/dropbear"
import "tholian-scout/internal/classifiers/dovecot"
import "tholian-scout/internal/classifiers/gogs"
import "tholian-scout/internal/classifiers/pam"
import "tholian-scout/internal/classifiers/postfix"
import "tholian-scout/internal/classifiers/sshd"
import "tholian-scout/internal/classifiers/sudo"

var Classifiers map[string]func(string, string, string, string, string) structs.Incident = map[string]func(string, string, string, string, string) structs.Incident{
	// "auth":                auth.Classify,
	// "auth-worker":         auth.Classify,
	"dovecot":             dovecot.Classify,
	"dropbear":            dropbear.Classify,
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
