package main

import ()

type MessageHandler struct {
	packet  *DHCPPacket
	options []DHCPOption
}

func (handler *MessageHandler) Handle(data []uint8) error {
	err := handler.BuildPacket(data)
	if err != nil {
		return err
	}

	for _, e := range handler.options {
		switch e.Type {
		case DHCP_MESSAGE_TYPE:
			switch e.Data[0] {
			case DHCP_DISCOVER:
			case DHCP_REQUEST:
			}
		}
	}

	return nil
}

func (handler *MessageHandler) BuildPacket(data []uint8) error {
	packet, err := NewDHCPPacket(data)
	handler.packet = &packet

	if err != nil {
		return err
	}

	if !packet.IsValid() {

	}

	optionParser := DHCPOptionParser{}
	handler.options, err = optionParser.Parse(data, DHCP_PACKET_END)

	if err != nil {
		return err
	}

	packet.Print()
	optionParser.Print(handler.options)

	return nil
}
