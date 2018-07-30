package main

import (
	"bufio"
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
	conn, _ := net.Dial(connectionType, connectionAddr)
	fmt.Print("Text to send: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	fmt.Fprintf(conn, text+"\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)
	conn.Close()
}
