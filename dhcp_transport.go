package main

import (
	"net"
)

type UDPListener struct {
	serverAddr  *net.UDPAddr
	connection  *net.UDPConn
	err         error
	isConnected bool
}

func NewUDPListener() UDPListener {
	listener := UDPListener{}
	listener.isConnected = false

	serverAddr, err := net.ResolveUDPAddr("udp", ":67")

	if err != nil {
		listener.err = err
		return listener
	}

	ln, err := net.ListenUDP("udp", serverAddr)

	if err != nil {
		listener.err = err
		return listener
	}

	listener.serverAddr = serverAddr
	listener.connection = ln
	listener.isConnected = true
	return listener
}

func (listener UDPListener) GetPacket() []byte {
	buffer := make([]byte, 1024)
	listener.connection.ReadFromUDP(buffer[:])
	return buffer
}

func (listener UDPListener) SendPacket(p Packet) {
	listener.connection.Write(p.ToBinary())
}
