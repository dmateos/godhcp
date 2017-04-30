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

func (listener UDPListener) GetPacket() DHCPPacket {
	buffer := make([]byte, 1024)
	listener.connection.ReadFromUDP(buffer[:])
	return NewDHCPPacket(buffer)
}

func (listener UDPListener) SendPacket() bool {
	return false
}
