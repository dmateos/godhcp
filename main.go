package main

import (
	"log"
)

func handle_packet(data []uint8) {
	packet, err := NewDHCPPacket(data)

	if err != nil {
		log.Print("could not parse packet")
		return
	}

	if !packet.IsValid() {
		log.Print("packet is not valid")
		return
	}

	optionParser := DHCPOptionParser{}
	options, err := optionParser.Parse(data, DHCP_PACKET_END)

	if err != nil {
		return
	}

	packet.Print()
	optionParser.Print(options)
}

func main() {
	dhcpListener, err := NewUDPListener()

	if err != nil {
		log.Fatal(err)
	}

	for {
		data, _, err := dhcpListener.GetPacket()
		if err != nil {
			log.Print(err)
			continue
		}

		go handle_packet(data)
	}
}
