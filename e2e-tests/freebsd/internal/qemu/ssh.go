package qemu

import "fmt"
import "os"
import "os/exec"
import "path/filepath"
import "strings"
import "time"

type SSH struct {
	User string
	Key  string
	Port string
	Sudo bool
}

func WaitForSSH(port string, key string, timeout time.Duration) (SSH, error) {

	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {

		for _, user := range []string{"freebsd", "root"} {

			client := SSH{
				User: user,
				Key:  key,
				Port: port,
				Sudo: user != "root",
			}

			if _, err := client.Output("true"); err == nil {
				return client, nil
			}

		}

		time.Sleep(2 * time.Second)

	}

	return SSH{}, fmt.Errorf("timed out waiting for SSH on port %s", port)

}

func (client SSH) args() []string {

	return []string{
		"-o", "BatchMode=yes",
		"-o", "StrictHostKeyChecking=no",
		"-o", "UserKnownHostsFile=/dev/null",
		"-o", "ConnectTimeout=10",
		"-i", client.Key,
		"-p", client.Port,
		client.User + "@127.0.0.1",
	}

}

func (client SSH) command(value string) string {

	if client.Sudo == true {
		return "sudo -n " + value
	}

	return value

}

func (client SSH) Command(value string) *exec.Cmd {

	arguments := client.args()
	arguments = append(arguments, client.command(value))

	cmd := exec.Command("ssh", arguments...)
	cmd.Env = append(os.Environ(), "LC_ALL=C")

	return cmd

}

func (client SSH) Run(value string) error {

	cmd := client.Command(value)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func (client SSH) Output(value string) ([]byte, error) {

	cmd := client.Command(value)
	cmd.Env = append(os.Environ(), "LC_ALL=C")

	return cmd.Output()

}

func (client SSH) Copy(names ...string) error {

	arguments := []string{
		"-o", "BatchMode=yes",
		"-o", "StrictHostKeyChecking=no",
		"-o", "UserKnownHostsFile=/dev/null",
		"-i", client.Key,
		"-P", client.Port,
	}

	arguments = append(arguments, names...)
	arguments = append(arguments, client.User+"@127.0.0.1:/tmp/")

	cmd := exec.Command("scp", arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func (client SSH) CopyAs(name string, remote string) error {

	arguments := []string{
		"-o", "BatchMode=yes",
		"-o", "StrictHostKeyChecking=no",
		"-o", "UserKnownHostsFile=/dev/null",
		"-i", client.Key,
		"-P", client.Port,
		name,
		client.User + "@127.0.0.1:" + remote,
	}

	cmd := exec.Command("scp", arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func GenerateSSHKey(dir string) (string, string, error) {

	private := filepath.Join(dir, "id_ed25519")
	public := private + ".pub"

	cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-N", "", "-q", "-f", private)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", "", err
	}

	buffer, err := os.ReadFile(public)

	if err != nil {
		return "", "", err
	}

	return private, strings.TrimSpace(string(buffer)), nil

}
