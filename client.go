package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
	"unsafe"
)

func client(network, addr string) {
	if network == "" {
		network = "tcp"
	}
	Dial, err := net.Dial(network, addr)

	if err != nil {
		fmt.Printf("Error in Dial: %s \n", err)
	}

	connected := false
	retries := 0

	for !connected && retries < 3 {
		connected = handshakeClient(Dial)
		if connected {
			fmt.Println("Client: Connected successfully!")
		} else {
			fmt.Println("Client: Connection attempt failed, retrying...")
			retries++
			time.Sleep(1 * time.Second)
		}
	}

	if !connected {
		fmt.Println("Client: Connection could not be established after 3 retries.")
		return
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}

}

func handshakeClient(conn net.Conn) bool {
	fmt.Println("Client: Trying handshake...")
	packet := Packet{1234, 1234, 100, 0, 0, 0, 1, 0} //SYN
	sendPacket(conn, packet)

	var data Packet
	buf := make([]byte, int(unsafe.Sizeof(Packet{})))
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		fmt.Printf("Client: Error happened when reading %s \n", err)
		return false
	} else {
		binary.Read(
			bytes.NewReader(buf),
			binary.BigEndian,
			&data,
		)
		fmt.Printf("Client: Received from server: %+v\n", data)
	}

	//Validate SYN-ACK flags
	if data.SYN == 1 && data.ACK == 1 {
		fmt.Println("Client: SYN-ACK received correctly")
		packet = Packet{1234, 1234, data.AckNo, data.SeqNo + 1, 1, 0, 0, 0} //ACK
		sendPacket(conn, packet)
	} else {
		fmt.Println("Client: Unexpected packet flags")
		return false
	}

	defer fmt.Printf("Client: handshake complete! \n")
	return true
}

func sendPacket(conn net.Conn, pkt Packet) {
	err := binary.Write(conn, binary.BigEndian, pkt)
	if err != nil {
		fmt.Printf("Error happened when sending %s \n", err)
	}
}
