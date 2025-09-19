package main

type Packet struct {
	SourcePort      uint16
	DestinationPort uint16
	SeqNo           uint32
	AckNo           uint32
	ACK             uint8
	RST             uint8
	SYN             uint8
	FIN             uint8
}
