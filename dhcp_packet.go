package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type DHCPPacket struct {
	Op, Htype, Hlen, Hops          uint8
	Xid                            uint32
	Secs, Flags                    uint16
	Ciaddr, Yiaddr, Siaddr, Giaddr uint32
	Chaddr                         [16]uint8
	Sname                          [64]uint8
	File                           [128]uint8
	Options                        DHCPOption
}

type DHCPOption struct {
	MagicCookie [4]uint8
	MessageType [3]uint8
	LeaseTime   [6]uint8
	SubnetMask  [6]uint8
	ServerId    [6]uint8
	End         uint8
}

const (
	BOOTREQUEST = 1
	BOOTREPLY   = 2
)

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

func NewDHCPPacket(data []byte) DHCPPacket {
	packet := DHCPPacket{}
	binary.Read(bytes.NewBuffer(data), binary.BigEndian, &packet)
	return packet
}

func (packet DHCPPacket) Print() {
	str := fmt.Sprintf(
		"op: %d\nhtype:%d\nhlen:%d\nhops:%d\nxid:%d\nsecs:%d\nflags:%d",
		packet.Op, packet.Htype, packet.Hlen, packet.Hops,
		packet.Xid, packet.Secs, packet.Flags,
	)

	ipStr := fmt.Sprintf(
		"C: %s\nY: %s\nS: %s\nG: %s\n",
		packet.Int2Ip(packet.Ciaddr), packet.Int2Ip(packet.Yiaddr),
		packet.Int2Ip(packet.Siaddr), packet.Int2Ip(packet.Giaddr),
	)

	fmt.Println("DHCP data:")
	fmt.Println(str)
	fmt.Println(ipStr)
	fmt.Println("DHCP options:")

}

func (packet DHCPPacket) IsValid() bool {
	if packet.Options.MagicCookie[0] == 99 &&
		packet.Options.MagicCookie[1] == 130 &&
		packet.Options.MagicCookie[2] == 83 &&
		packet.Options.MagicCookie[3] == 99 {
		return true
	}
	return false
}

func (packet DHCPPacket) Int2Ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func (packet DHCPPacket) ToBinary() []byte {
	return make([]byte, 12)
}
