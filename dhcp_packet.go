package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type DHCPPacket struct {
	Op, Htype, Hlen, Hops          uint8
	Xid                            uint32
	Secs, Flags                    uint16
	Ciaddr, Yiaddr, Siaddr, Giaddr uint32
	Chaddr                         [16]uint8
	Sname                          [64]uint8
	File                           [128]uint8
	Options                        [312]uint8
}

const (
	BOOTREQUEST = iota
	BOOTREPLY   = iota
)

func NewDHCPPacket(data []byte) DHCPPacket {
	packet := DHCPPacket{}
	binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &packet)
	return packet
}

func (packet DHCPPacket) Print() {
	str := fmt.Sprintf(
		"op: %d\nhtype:%d\nhlen:%d\nhops:%d\nxid:%d\nsecs:%d\nflags:%d",
		packet.Op, packet.Htype, packet.Hlen, packet.Hops, packet.Xid, packet.Secs, packet.Flags,
	)

	fmt.Println("Packet data:")
	fmt.Println(str)
}

func (packet DHCPPacket) IsValid() bool {
	if packet.Options[0] == 99 && packet.Options[1] == 130 &&
		packet.Options[2] == 83 && packet.Options[3] == 99 {
		return true
	}
	return false
}
