package main

import (
	"net"
)

type Packet interface {
	ToBinary() []byte
}

type Listener interface {
	GetPacket() (Packet, *net.UDPAddr)
	SendPacket(Packet, *net.UDPAddr)
}
