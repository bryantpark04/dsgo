package main

import (
	"fmt"
	"messages"
)

func server(in chan messages.Message, out map[string]chan messages.Message) {
	<-in
	// switch on message type
	out[messages.CLIENT] <- Pong{messages.BaseMessageFrom(messages.SERVER)}
}

func client(in chan messages.Message, out map[string]chan messages.Message) {
	out[messages.SERVER] <- Ping{messages.BaseMessageFrom(messages.CLIENT)}
	fmt.Println("Ping!")
	<-in
	fmt.Println("Pong")
}

func main() {
	directory := map[string]chan messages.Message{
		messages.SERVER: make(chan messages.Message),
		messages.CLIENT: make(chan messages.Message),
	}

	go server(directory[messages.SERVER], directory)

	client(directory[messages.CLIENT], directory)
}