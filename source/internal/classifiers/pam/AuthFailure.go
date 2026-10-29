package pam

import "strconv"
import "strings"

type AuthFailure struct {
	Logname    string `json:"logname"`
	UID        uint16 `json:"uid"`
	EUID       uint16 `json:"euid"`
	TTY        string `json:"tty"`
	RemoteUser string `json:"ruser"`
	RemoteHost string `json:"rhost"`
	User       string `json:"user"`
}

func NewAuthFailure(message string) AuthFailure {

	var auth_failure AuthFailure

	auth_failure.Parse(message)

	return auth_failure

}

func (auth_failure *AuthFailure) Parse(message string) {

	if strings.HasPrefix(message, "authentication failure;") {

		tmp := splitLine(strings.TrimSpace(message[23:]), " ")

		for t := 0; t < len(tmp); t++ {

			var chunk = strings.Split(tmp[t], "=")

			if len(chunk) == 2 {

				if chunk[0] == "logname" {

					auth_failure.Logname = strings.TrimSpace(chunk[1])

				} else if chunk[0] == "uid" {

					num, err := strconv.ParseUint(chunk[1], 10, 16)

					if err == nil {
						auth_failure.UID = uint16(num)
					}

				} else if chunk[0] == "euid" {

					num, err := strconv.ParseUint(chunk[1], 10, 16)

					if err == nil {
						auth_failure.EUID = uint16(num)
					}

				} else if chunk[0] == "tty" {

					auth_failure.TTY = strings.TrimSpace(chunk[1])

				} else if chunk[0] == "ruser" {

					auth_failure.RemoteUser = strings.TrimSpace(chunk[1])

				} else if chunk[0] == "rhost" {

					auth_failure.RemoteHost = strings.TrimSpace(chunk[1])

				} else if chunk[0] == "user" {

					auth_failure.User = strings.TrimSpace(chunk[1])

				}

			}

		}

	}

}
