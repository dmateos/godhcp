package main

import (
	"fmt"
)

func handle_packet(data []uint8) {
	packet := NewDHCPPacket(data)

	if !packet.IsValid() {
		fmt.Println("false")
		return
	}

	optionParser := DHCPOptionParser{}
	options := optionParser.Parse(data, DHCP_PACKET_END)

	packet.Print()
	optionParser.Print(options)
}

func main() {
	dhcpListener := NewUDPListener()

	for {
		data, _ := dhcpListener.GetPacket()
		go handle_packet(data)
	}
}
