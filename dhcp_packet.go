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
	Magic                          [4]uint8
}

func NewDHCPPacket(data []byte) (DHCPPacket, error) {
	packet := DHCPPacket{}
	err := binary.Read(bytes.NewBuffer(data[:DHCP_PACKET_END]), binary.BigEndian, &packet)

	if err != nil {
		return packet, err
	}

	return packet, nil
}

func (packet *DHCPPacket) Print() {
	str := fmt.Sprintf(
		"op: %d, htype:%d, hlen:%d, hops:%d, xid:%d, secs:%d, flags:%d",
		packet.Op, packet.Htype, packet.Hlen, packet.Hops,
		packet.Xid, packet.Secs, packet.Flags,
	)

	ipStr := fmt.Sprintf(
		"C: %s, Y: %s, S: %s, G: %s",
		packet.Int2Ip(packet.Ciaddr), packet.Int2Ip(packet.Yiaddr),
		packet.Int2Ip(packet.Siaddr), packet.Int2Ip(packet.Giaddr),
	)

	fmt.Println(str)
	fmt.Println(ipStr)
}

func (packet *DHCPPacket) IsValid() bool {
	if packet.Magic[0] == 99 &&
		packet.Magic[1] == 130 &&
		packet.Magic[2] == 83 &&
		packet.Magic[3] == 99 {
		return true
	}
	return false
}

func (packet *DHCPPacket) Int2Ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func (packet *DHCPPacket) ToBinary() []byte {
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, packet)
	return buffer.Bytes()
}
