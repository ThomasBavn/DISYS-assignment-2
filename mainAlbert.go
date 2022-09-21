package main

import (
	"fmt"
	"math/rand"
	"os"
)

var serverChan chan Message
var clientChan chan Message

var clientAck int
var clientSyn int

var serverAck int
var serverSyn int

var messagesToSent int

type Message struct {
	sec  int
	ack  int
	data *int
}

func main() {
	messagesToSent = 10
	serverChan = make(chan Message)
	clientChan = make(chan Message)
	go server()
	go clientInitialising()
	for {

	}
}

func server() {

	m := <-clientChan
	if m.data == nil {
		fmt.Println("Server acknoledged first messaged: ")
	} else {
		fmt.Println("Server acknoledged messaged: ", *m.data, " (sec:", m.sec, ")")
	}
	serverChan <- Message{m.sec + 1, m.ack + 1, nil}

	go server()

}

func clientInitialising() {
	fmt.Println("Sent first request", " (sec: ", 1, ")")
	clientChan <- Message{1, 0, nil}
	m := <-serverChan
	fmt.Println("connection established")
	toSent := rand.Intn(40)
	fmt.Println("client sent messaged: ", toSent, " (sec: ", m.sec+1, ")")
	clientChan <- Message{m.sec + 1, m.ack + 1, &toSent}
	messagesToSent--
	if messagesToSent == 0 {
		os.Exit(3)
	}
	go client()
}

func client() {
	m := <-serverChan
	toSent := rand.Intn(40)
	fmt.Println("client sent messaged: ", toSent, " (syn: ", m.sec+1, ")")
	clientChan <- Message{m.sec + 1, m.ack + 1, &toSent}
	messagesToSent--
	if messagesToSent == 0 {
		os.Exit(3)
	}

	go client()
}
