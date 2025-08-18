package main

import (
	"fmt"
	"util"
)

func server(in chan util.Message, out map[string]chan util.Message) {
	<-in
	// switch on message type
	out[util.CLIENT] <- Pong{util.BaseMessageFrom(util.SERVER)}
}

func client(in chan util.Message, out map[string]chan util.Message) {
	out[util.SERVER] <- Ping{util.BaseMessageFrom(util.CLIENT)}
	fmt.Println("Ping!")
	<-in
	fmt.Println("Pong")
}

func main() {
	directory := map[string]chan util.Message{
		util.SERVER: make(chan util.Message),
		util.CLIENT: make(chan util.Message),
	}

	go server(directory[util.SERVER], directory)

	client(directory[util.CLIENT], directory)
}