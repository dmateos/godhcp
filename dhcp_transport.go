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

type UDPPacket struct {
	data    []byte
	address *net.UDPAddr
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

func (listener UDPListener) GetPacket() ([]byte, *net.UDPAddr) {
	buffer := make([]byte, 1024)
	n, addr, err := listener.connection.ReadFromUDP(buffer[:])

	if n > 1 {

	}

	if err != nil {

	}

	return buffer, addr
}

func (listener UDPListener) SendPacket(p Packet, addr *net.UDPAddr) {
	listener.connection.WriteToUDP(p.ToBinary(), addr)
}
