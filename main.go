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
	go serverThread(ch2)
	go middlewareThread(ch, ch2)
	clientThread(ch)

	// go server("tcp", ":8090")
	// time.Sleep(100 * time.Millisecond)
	// go middleware("tcp", ":8080", "127.0.0.1:8090")
	// client("tcp", "127.0.0.1:8080")

}

func clientThread(ch chan Packet) {
	response := Packet{}
	request := Packet{1234, 1234, 100, 0, 0, 0, 1, 0}
	fmt.Printf("Client: Seq=%d, Ack=%d \n", request.SeqNo, request.AckNo)
	ch <- request
	timeOut := time.Now().Add(1 * time.Second)
	for ok := true; ok; ok = timeOut.After(time.Now()) {
		if len(ch) > 0 {
			response = <-ch
			break
		}
	}
	if response.SourcePort == 0 {
		panic("Client: packet not recieved")
	} else if response.SYN == 1 && response.ACK == 1 {
		request = Packet{1234, 1234, response.SeqNo + 1, response.AckNo + 1, 1, 0, 0, 0}
		fmt.Printf("Client: Seq=%d, Ack=%d \n", request.SeqNo, request.AckNo)
		ch <- request
	}

}

func serverThread(ch2 chan Packet) {
	request := <-ch2
	//prevSeq := request.SeqNo + 1
	if request.SYN == 1 {
		response := Packet{1234, 1234, 300, request.SeqNo + 1, 1, 0, 1, 0}
		fmt.Printf("Server: Seq=%d, Ack=%d \n", response.SeqNo, response.AckNo)
		ch2 <- response

		timeOut := time.Now().Add(1 * time.Second)
		for ok := true; ok; ok = timeOut.After(time.Now()) {
			if len(ch2) > 0 {
				response = <-ch2
				break
			}
		}
		if response.SourcePort == 0 {
			panic("Server: packet not recieved")
		} else if response.ACK == 1 {
			fmt.Printf("Server: Seq=%d, Ack=%d \n", response.SeqNo, response.AckNo)
		}
	}

}
func middlewareThread(ch chan Packet, ch2 chan Packet) {

	var ch1Avail, ch2Avail bool

	for ok := true; ok; ok = !(ch1Avail || ch2Avail) {
		if len(ch) > 0 {
			ch1Avail = true
		}
		if len(ch2) > 0 {
			ch2Avail = true
		}

	}
	x := rand.Float32()
	if ch1Avail {
		request := <-ch
		switch {
		case x < 0.33:
			time.Sleep(10 * time.Millisecond)
			ch2 <- request
		case x >= 0.33 && x < 0.90:
			ch2 <- request
		default:
			return
		}
	} else if ch2Avail {
		request := <-ch2
		switch {
		case x < 0.33:
			time.Sleep(10 * time.Millisecond)
			ch <- request
		case x >= 0.33 && x < 0.90:
			ch <- request
		default:
			return
		}
	}
}
