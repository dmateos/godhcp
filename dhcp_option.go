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

func (parser DHCPOptionParser) ParseOptions(data []byte, offset uint8) *DHCPOption {

	for offset <= uint8(len(data)) {
		var option, length uint8
	}

	var option, length uint8

	err := binary.Read(bytes.NewBuffer(data[offset:offset+1]), binary.BigEndian, &option)

	if err != nil {
		log.Fatal(err)
	}

	err = binary.Read(bytes.NewBuffer(data[offset+1:offset+2]), binary.BigEndian, &length)

	if err != nil {
		log.Fatal(err)
	}

	optionData := make([]uint8, length)
	err = binary.Read(bytes.NewBuffer(data[offset+2:offset+2+length]), binary.BigEndian, &optionData)

	if err != nil {

	}

	dhcpOption := DHCPOption{}
	dhcpOption.Type = option
	dhcpOption.Length = length
	dhcpOption.Data = optionData

	return &dhcpOption
}
