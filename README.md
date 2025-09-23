# Handin 2 - TCP/IP Simulator in Go

- (1) Is implemented
- (2) Is implemented
- (3) Is implemented on (2), not working on (1)

### a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
For (1) we use:
`"fmt",
"math/rand" and 
"time"`

for (2) we use: 
`
	"bytes",
	"encoding/binary",
	"fmt",
	"io",
	"net",
	"time",
	"unsafe" and
	"math/rand"
`

For meta-data transfer we are using a `Packet` struct which is defined in Packet.go

`type Packet struct {
	SourcePort      uint16
	DestinationPort uint16
	SeqNo           uint32
	AckNo           uint32
	ACK             uint8
	RST             uint8
	SYN             uint8
	FIN             uint8
}`

### b) Does your implementation use threads or processes? Why is it not realistic to use threads?
Our implementation (both (1) and (2)) uses threads. Threads are not independent and system resources like memory etc, this can cause that multiple threads are trying to access the same memory address, whereas processes gets allocated their own memory. Threads are smaller, takes less time to setup and teardown again. 

### c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
Based on sequence number and acknowledgment number, the ordering can be identified and in the case of network changes re-ordering can be done. 

### d) In case messages can be delayed or lost, how does your implementation handle message loss?
Our implementation of (1) uses panic which terminates the program.

In our implementation of (2), it uses timeout while waiting for packets, if no packet is received within 3 tries it then terminates with an error.

### e) Why is the 3-way handshake important?
To establish an 2-way communication that makes sure that the server is receiving from and can send to the client, and that the client can do the same to the server.
