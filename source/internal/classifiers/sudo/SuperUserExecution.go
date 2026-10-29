package sudo

import "strings"

type SuperUserExecution struct {
	TSID      string   `json:"tsid"`
	TTY       string   `json:"tty"`
	Folder    string   `json:"folder"`
	User      string   `json:"user"`
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

func NewSuperUserExecution(message string) SuperUserExecution {

	var suexec SuperUserExecution

	suexec.Parse(message)

	return suexec

}

func (suexec *SuperUserExecution) Parse(message string) {

	if strings.Contains(message, " ; ") {

		chunks := strings.Split(message, " ; ")

		for c := 0; c < len(chunks); c++ {

			chunk := strings.TrimSpace(chunks[c])

			if strings.HasPrefix(chunk, "COMMAND=") {

				tmp1 := strings.TrimSpace(chunk[8:])

				if strings.Contains(tmp1, " ") {

					tmp2 := strings.Split(tmp1, " ")

					suexec.Command = tmp2[0]
					suexec.Arguments = tmp2[1:]

				} else {
					suexec.Command = tmp1
					suexec.Arguments = []string{}
				}

			} else if strings.HasPrefix(chunk, "PWD=") {
				suexec.Folder = strings.TrimSpace(chunk[4:])
			} else if strings.HasPrefix(chunk, "TSID=") {
				suexec.TSID = strings.TrimSpace(chunk[5:])
			} else if strings.HasPrefix(chunk, "TTY=") {
				suexec.TTY = strings.TrimSpace(chunk[4:])
			} else if strings.HasPrefix(chunk, "USER=") {
				suexec.User = strings.TrimSpace(chunk[5:])
			}

		}

	}

}

func (suexec *SuperUserExecution) ToProgramName() string {

	var result string

	if strings.Contains(suexec.Command, "/") {

		tmp := strings.Split(suexec.Command, "/")

		if len(tmp) > 0 {
			result = tmp[len(tmp)-1]
		}

	} else {

		result = suexec.Command

	}

	return result

}
