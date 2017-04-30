package main

type Packet interface {
	ToBinary() []byte
}

type Listener interface {
	GetPacket() Packet
	SendPacket(Packet)
}
