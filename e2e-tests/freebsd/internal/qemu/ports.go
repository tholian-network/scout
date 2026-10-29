package qemu

import "net"

func FreePort() (string, error) {

	listener, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		return "", err
	}

	defer listener.Close()

	_, port, err := net.SplitHostPort(listener.Addr().String())

	if err != nil {
		return "", err
	}

	return port, nil

}
