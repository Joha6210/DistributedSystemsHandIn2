package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func server(network, port string) {
	if network == "" {
		network = "tcp"
	}
	Listener, err := net.Listen(network, port)

	if err != nil {
		fmt.Printf("Server: Error in listening: %s", err)
	} else {
		fmt.Printf("Server: listening on %s... \n", port)
	}

	for {
		connection, err := Listener.Accept()
		if err != nil {
			fmt.Printf("Server: Error in connection: %s", err)
		}
		result := handshakeServer(connection)
		if result  {
			fmt.Println("Server: Connection established")
			go handleClient(connection)
		}
		

	}
}

func handleClient(connection net.Conn) {
	for{
		time.Sleep(100*time.Millisecond)
	}
}

func handshakeServer(conn net.Conn) bool {

		var data Packet
		err := binary.Read(conn, binary.BigEndian, &data)
		if err != nil {
			fmt.Println("Server: Error decoding response:", err)
			return false
		}
		fmt.Printf("Server: Received from client: %+v\n", data)

		if data.SYN != 1 {
			fmt.Println("Server: Error expected SYN package")
		}

		synAck := Packet{
			SourcePort:      data.DestinationPort,
			DestinationPort: data.SourcePort,
			SeqNo:           300,
			AckNo:           data.SeqNo + 1,
			SYN:             1,
			ACK:             1,
		}
		sendPacket(conn, synAck)
		fmt.Println("Server: Sent SYN-ACK")

		err = binary.Read(conn, binary.BigEndian, &data)
		if err != nil {
			fmt.Println("Server: Error decoding response:", err)
			return false
		}
		fmt.Printf("Server: Received from client: %+v\n", data)		

		if data.ACK == 1{
			fmt.Println("Server: Received ACK from client")	
			return true
		}else{
			return false
		}

}
