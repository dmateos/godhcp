package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

type DHCPOptionParser struct {
}

type DHCPOption struct {
	Type   uint8
	Length uint8
	Data   []uint8
}

const (
	DHCP_DISCOVER = 1
	DHCP_OFFER    = 2
	DHCP_REQUEST  = 3
	DHCP_DECLINE  = 4
	DHCP_ACK      = 5
	DHCP_NAK      = 6
	DHCP_RELEASE  = 7
	DHCP_INFORM   = 8
)

func (parser DHCPOptionParser) ParseOptions(data []byte, offset uint8) []DHCPOption {
	var optionArray []DHCPOption

	for offset+2 < uint8(len(data)) {
		var option, length uint8

		//Option/message type is first byte (first at 240 of DHCP message proper)
		err := binary.Read(bytes.NewBuffer(data[offset:offset+1]), binary.BigEndian, &option)

		if err != nil {
			log.Fatal(err)
		}

		//Length of the message data is second byte
		err = binary.Read(bytes.NewBuffer(data[offset+1:offset+2]), binary.BigEndian, &length)

		if err != nil {
			log.Fatal(err)
		}

		//Read data
		optionData := make([]uint8, length)
		err = binary.Read(bytes.NewBuffer(data[offset+2:offset+2+length]), binary.BigEndian, &optionData)

		if err != nil {
			log.Fatal(err)
		}

		dhcpOption := DHCPOption{}
		dhcpOption.Type = option
		dhcpOption.Length = length
		dhcpOption.Data = optionData

		optionArray = append(optionArray, dhcpOption)

		//Get past the option, length and data for next option in the packet
		offset += (2 + length)
	}

	return optionArray
}
