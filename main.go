package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Header struct {
	sourcePort      int
	destinationPort int
	seqNo           int
	AckNo           int
	ACK             int
	RST             int
	SYN             int
	FIN             int
}

func main() {
	fmt.Println("hello")
	ch := make(chan Packet)
	ch2 := make(chan Packet)
	go middlewareThread(ch, ch2)
	go clientThread(ch)
	go serverThread(ch2)

	go server("tcp", ":8090")
	time.Sleep(100 * time.Millisecond)
	go middleware("tcp", ":8080", "127.0.0.1:8090")
	client("tcp", "127.0.0.1:8080")

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

func serverThread(ch2 chan Packet) {
	request := <-ch2
	if request.SYN == 1 {
		response := Packet{1234, 1234, 300, request.SeqNo + 1, 1, 0, 1, 0}
		fmt.Printf("Seq=%d, Ack=%d \n", response.SeqNo, response.AckNo)
		ch2 <- response
	}

}
func middlewareThread(ch chan Packet, ch2 chan Packet) {
	rand.Seed(7)
	x := rand.Int()
	request := <-ch
	switch x % 2 {
	case 0:
		time.Sleep(30 * time.Millisecond)
		ch2 <- request
	case 1:

	}

	request2 := <-ch2
	ch <- request2
}
