package main

import (
	"log"
)

func handle_packet(data []uint8) {
	handler := MessageHandler{}
	err := handler.Handle(data)

	if err != nil {
		log.Print(err)
	}
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
