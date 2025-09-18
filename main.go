package main

import (
	"fmt"
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
	go client(ch)
	go server(ch)

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

func server(ch chan Header) {
	request := <-ch
	if request.SYN == 1 {
		response := Header{1234, 1234, 300, request.seqNo + 1, 1, 0, 1, 0}
		fmt.Printf("Seq=%d, Ack=%d \n", response.seqNo, response.AckNo)
		ch <- response
	}

}
