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

		var data Packet
		connection.SetReadDeadline(time.Now().Add(3 * time.Second))
		err = binary.Read(connection, binary.BigEndian, &data)
		if err != nil {
			fmt.Println("Server: Error decoding response:", err)
		}

		if data.SYN != 1 {
			fmt.Println("Server: Error expected SYN package")
		} else {
			var retries int
			fmt.Println("Server: Received from SYN from client, trying handshake...")
			for retries < 3 {
				result := handshakeServer(connection, data)
				if result {
					fmt.Println("Server: Handshake complete, connection established!")
					go handleClient(connection)
					return
				} else {
					retries++
				}
			}
			if retries > 3 {
				fmt.Println("Server: Connection not established!")
			}

		}

	}
}

func handleClient(connection net.Conn) {
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func handshakeServer(conn net.Conn, initPacket Packet) bool {

	var data Packet

	data = Packet{
		SourcePort:      initPacket.DestinationPort,
		DestinationPort: initPacket.SourcePort,
		SeqNo:           300,
		AckNo:           initPacket.SeqNo + 1,
		SYN:             1,
		ACK:             1,
	}
	sendPacket(conn, data) //SYN-ACK
	fmt.Println("Server: Sent SYN-ACK")

	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	err := binary.Read(conn, binary.BigEndian, &data)
	if err != nil {
		fmt.Println("Server: Error decoding response:", err)
		return false
	}

	if data.ACK == 1 {
		fmt.Println("Server: Received ACK from client")
		return true
	} else {
		return false
	}

}
