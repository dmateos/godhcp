package main

import (
	"net"
)

type UDPListener struct {
	serverAddr  *net.UDPAddr
	connection  *net.UDPConn
	isConnected bool
}

func NewUDPListener() (UDPListener, error) {
	listener := UDPListener{}
	listener.isConnected = false

	serverAddr, err := net.ResolveUDPAddr("udp", ":67")

	if err != nil {
		return listener, err
	}

	ln, err := net.ListenUDP("udp", serverAddr)

	if err != nil {
		return listener, err
	}

	listener.serverAddr = serverAddr
	listener.connection = ln
	listener.isConnected = true
	return listener, nil
}

func (listener *UDPListener) GetPacket() ([]uint8, *net.UDPAddr, error) {
	buffer := make([]byte, 2048)
	n, addr, err := listener.connection.ReadFromUDP(buffer[:])

	if err != nil {
		return nil, nil, err
	}

	//So we get the correct sized buffer
	//Unless this can be done better?
	new_buffer := make([]byte, n)
	copy(new_buffer, buffer)

	return new_buffer, addr, nil
}

func (listener *UDPListener) SendPacket(data []uint8, addr *net.UDPAddr) {
	listener.connection.WriteToUDP(data, addr)
}
