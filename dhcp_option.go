package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	DHO_PAD = 0
	DHO_END = 255
)

func (parser DHCPOptionParser) Print(options []DHCPOption) {
	fmt.Println(options)
}

func (parser DHCPOptionParser) Parse(data []byte, offset int) []DHCPOption {
	var optionArray []DHCPOption

	for offset+2 < len(data) {
		var option, length uint8

		//Option/message type is first byte (first at 240 of DHCP message proper)
		err := binary.Read(bytes.NewBuffer(data[offset:offset+1]), binary.BigEndian, &option)

		if err != nil {
			log.Fatal(err)
		}

		if option == DHO_END {
			break
		}

		if option == DHO_PAD {
			offset += 1
			continue
		}

		//Length of the message data is second byte
		err = binary.Read(bytes.NewBuffer(data[offset+1:offset+2]), binary.BigEndian, &length)

		if err != nil {
			log.Fatal(err)
		}

		//Read data
		optionData := make([]uint8, length)
		if length > 0 {
			err = binary.Read(bytes.NewBuffer(data[offset+2:offset+2+int(length)]), binary.BigEndian, &optionData)
		}

		if err != nil {
			log.Fatal(err)
		}

		dhcpOption := DHCPOption{}
		dhcpOption.Type = option
		dhcpOption.Length = length
		dhcpOption.Data = optionData

		optionArray = append(optionArray, dhcpOption)

		//Get past the option, length and data for next option in the packet
		offset += (2 + int(length))
	}

	return optionArray
}
