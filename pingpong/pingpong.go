package main

import (
	"fmt"
	"util"
)

func server(self string, in chan util.Message, out util.Directory) {
	msg := <-in
	switch t := msg.(type) {
	case Ping:
		out[util.CLIENT] <- Pong{util.BaseMessageFrom(util.SERVER)}
	default:
		fmt.Println("Received unknown message: ", t)
	}
}

func client(in chan util.Message, out util.Directory) {
	out[util.SERVER] <- Ping{util.BaseMessageFrom(util.CLIENT)}
	fmt.Println("Ping!")
	msg := <-in
	switch t := msg.(type) {
	case Pong:
		fmt.Println("Pong")
	default:
		fmt.Println("Received unknown message: ", t)
	}
	// TODO loop
}

func main() {
	directory := util.Directory{
		util.SERVER: make(chan util.Message),
		util.CLIENT: make(chan util.Message),
	}

	go server(util.SERVER, directory[util.SERVER], directory)

	client(directory[util.CLIENT], directory)
}
