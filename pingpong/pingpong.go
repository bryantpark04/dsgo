package main

import (
	"fmt"
	"util"
)

func server(self string, in chan util.Message, out util.Directory) {
	for msg := range in {
		switch t := msg.(type) {
		case Ping:
			util.Send(out[util.CLIENT], Pong{util.BaseMessageFrom(util.SERVER)})
		default:
			fmt.Println("Received unknown message: ", t)
		}
	}
}

func client(in chan util.Message, out util.Directory) {
	for {
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)

		switch input {
		case "ping":
			util.Send(out[util.SERVER], Ping{util.BaseMessageFrom(util.CLIENT)})
			msg := <-in
			switch t := msg.(type) {
			case Pong:
				fmt.Println("pong")
			default:
				fmt.Println("Received unknown message: ", t)
			}
		case "", "quit":
			return
		default:
			fmt.Println("Command not recognized. \"quit\" to quit.")
		}
	}
}

func main() {
	directory := util.Directory{
		util.SERVER: make(chan util.Message),
		util.CLIENT: make(chan util.Message),
	}

	go server(util.SERVER, directory[util.SERVER], directory)

	client(directory[util.CLIENT], directory)
}
