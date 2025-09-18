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
	ch := make(chan Header)
	ch2 := make(chan Header)
	go middleware(ch, ch2)
	go client(ch)
	go server(ch2)
	for {

	}

}

func client(ch chan Header) {
	request := Header{1234, 1234, 100, 0, 0, 0, 1, 0}
	fmt.Printf("Seq=%d, Ack=%d \n", request.seqNo, request.AckNo)
	ch <- request
	response := <-ch
	if response.SYN == 1 && response.ACK == 1 {
		request = Header{1234, 1234, response.AckNo, response.seqNo + 1, 1, 0, 0, 0}
		fmt.Printf("Seq=%d, Ack=%d \n", request.seqNo, request.AckNo)
		ch <- request
	}

}

func server(ch2 chan Header) {
	request := <-ch2
	if request.SYN == 1 {
		response := Header{1234, 1234, 300, request.seqNo + 1, 1, 0, 1, 0}
		fmt.Printf("Seq=%d, Ack=%d \n", response.seqNo, response.AckNo)
		ch2 <- response
	}

}
func middleware(ch chan Header, ch2 chan Header) {
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
