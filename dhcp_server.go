package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":67")
	ln, err := net.ListenUDP("udp", ServerAddr)

	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 1024)
		ln.ReadFromUDP(buf[:])
		packet := NewDhcpPacket(buf)
		packet.Print()

		if packet.IsValid() {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
