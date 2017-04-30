package main

import (
	"fmt"
)

func main() {
	dhcpListener := NewUDPListener()

	for {
		packet := dhcpListener.GetPacket()
		packet.Print()

		if packet.IsValid() {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
