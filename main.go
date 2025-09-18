package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	ch := make(chan string)
	go client(ch)
	go server(ch)
}

func client(ch chan string) {

}

func server(ch chan string) {

}
