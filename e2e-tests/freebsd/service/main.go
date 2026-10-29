package main

import "flag"
import "fmt"
import "net"
import "os"

func main() {

	listen := flag.String("listen", ":8080", "Listen address")

	flag.Parse()

	listener, err := net.Listen("tcp", *listen)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	for {

		connection, err := listener.Accept()

		if err != nil {
			continue
		}

		connection.Close()

	}

}
