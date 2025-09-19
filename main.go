package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan Packet)
	//go clientThread(ch)
	//go serverThread(ch)

	go server("tcp", ":8080")
	time.Sleep(100 * time.Millisecond)
	go client("tcp", "127.0.0.1:8080")

	for {
		time.Sleep(100 * time.Millisecond)
	}

}

func clientThread(ch chan Packet) {
	request := Packet{1234, 1234, 100, 0, 0, 0, 1, 0}
	fmt.Printf("Seq=%d, Ack=%d \n", request.SeqNo, request.AckNo)
	ch <- request
	response := <-ch
	if response.SYN == 1 && response.ACK == 1 {
		request = Packet{1234, 1234, response.AckNo, response.SeqNo + 1, 1, 0, 0, 0}
		fmt.Printf("Seq=%d, Ack=%d \n", request.SeqNo, request.AckNo)
		ch <- request
	}

}

func serverThread(ch chan Packet) {
	connEstablished := false

	for !connEstablished {
		request := <-ch
		if request.SYN == 1 {
			response := Packet{1234, 1234, 300, request.SeqNo + 1, 1, 0, 1, 0}
			fmt.Printf("Seq=%d, Ack=%d \n", response.SeqNo, response.AckNo)
			ch <- response
		} else if request.ACK == 1 {
			connEstablished = true
		} else {
			break
		}
	}

	fmt.Println("Connection with client established!")

}
