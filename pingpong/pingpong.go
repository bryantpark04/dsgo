package main

import (
	"fmt"
	"util"
)

func server(in chan util.Message, out util.Directory) {
	<-in
	// switch on message type
	out[util.CLIENT] <- Pong{util.BaseMessageFrom(util.SERVER)}
}

func client(in chan util.Message, out util.Directory) {
	out[util.SERVER] <- Ping{util.BaseMessageFrom(util.CLIENT)}
	fmt.Println("Ping!")
	<-in
	fmt.Println("Pong")
}

func main() {
	directory := util.Directory{
		util.SERVER: make(chan util.Message),
		util.CLIENT: make(chan util.Message),
	}

	go server(directory[util.SERVER], directory)

	client(directory[util.CLIENT], directory)
}
