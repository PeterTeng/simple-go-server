package main

import (
	"fmt"
	"net"
)

const (
	connectionHost = "localhost"
	connectionPort = "8888"
	connectionType = "udp"
	connectionAddr = connectionHost + ":" + connectionPort
)

func main() {
	packetConn, err := net.ListenPacket(connectionType, connectionAddr)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}

	defer packetConn.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := packetConn.ReadFrom(buf)
		if err != nil {
			fmt.Println("Error Reading:", err.Error())
		}

		go serve(packetConn, addr, buf[:n])
	}
}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	// 0 - 1: ID
	// 2: QR(1): Opcode(4)
	buf[2] |= 0x80 // Set QR bit

	pc.WriteTo(buf, addr)
}
