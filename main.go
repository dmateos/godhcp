package main

import (
	"fmt"
)

func handle_packet(packet DHCPPacket) {
	packet.Print()

	if packet.IsValid() {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func main() {
	dhcpListener := NewUDPListener()

	for {
		packet := dhcpListener.GetPacket()
		go handle_packet(packet)
	}
}
