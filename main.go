package main

import (
	"fmt"
)

func handle_packet(packet *DHCPPacket, data []uint8) {
	if !packet.IsValid() {
		fmt.Println("false")
		return
	}

	packet.Print()
	optionParser := DHCPOptionParser{}
	option := optionParser.ParseOptions(data, DHCP_PACKET_END)
	fmt.Println(option)
}

func main() {
	dhcpListener := NewUDPListener()

	for {
		data, _ := dhcpListener.GetPacket()
		packet := NewDHCPPacket(data)
		go handle_packet(&packet, data)
	}
}
