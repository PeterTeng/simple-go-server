package main

import (
	"fmt"
	"net"
	"os"
)

const (
	connectionHost = "localhost"
	connectionPort = "8888"
	connectionType = "tcp"
	connectionAddr = connectionHost + ":" + connectionPort
)

func main() {
	l, err := net.Listen(connectionType, connectionAddr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("Listening on " + connectionAddr)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	conn.Write([]byte("Message received.\n"))
	conn.Close()
}
