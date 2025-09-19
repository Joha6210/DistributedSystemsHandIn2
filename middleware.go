package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"
)

func middleware(network, port, destAddress string) {
	if network == "" {
		network = "tcp"
	}

	// Listen for client
	listener, err := net.Listen(network, port)
	if err != nil {
		fmt.Printf("Middleware: Error in listening: %s\n", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Middleware: Listening on %s\n", port)

	// Accept client connections
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Middleware: Error accepting connection: %s\n", err)
			continue
		}

		// Handle each client connection concurrently
		go handleConnection(clientConn, destAddress, network)
	}
}

func handleConnection(client net.Conn, serverAddr, network string) {
	defer client.Close()

	// Dial to actual server
	server, err := net.Dial(network, serverAddr)
	if err != nil {
		fmt.Printf("Middleware: Error dialing to server: %s\n", err)
		return
	}
	defer server.Close()

	fmt.Println("Middleware: Connected to server.")

	// Bidirectional forwarding
	go forwardPackets(client, server, "Client → Server")
	forwardPackets(server, client, "Server → Client")
}

func forwardPackets(src net.Conn, dst net.Conn, label string) {
	for {
		var packet Packet

		err := binary.Read(src, binary.BigEndian, &packet)
		if err != nil {
			if err == io.EOF {
				fmt.Println(label + ": Connection closed.")
			} else {
				fmt.Printf(label+": Error reading packet: %s\n", err)
			}
			return
		}

		// Simulate packet delay
		delay := time.Duration(rand.Intn(20)+10) * time.Millisecond // 10–30 ms Wifi or LAN-like
		fmt.Printf("Middleware: Packet delay %v\n", delay)
		time.Sleep(delay)

		// Simulate packet loss (e.g., 5% chance to drop)
		if rand.Float32() < 0.10 {
			fmt.Printf(label+": Dropped packet (simulated loss): %+v\n", packet)
			continue
		}

		fmt.Printf(label+": Forwarding packet: %+v\n", packet)
		err = binary.Write(dst, binary.BigEndian, packet)
		if err != nil {
			fmt.Printf(label+": Error writing packet: %s\n", err)
			return
		}
	}
}
