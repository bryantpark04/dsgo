package main

import (
	"fmt"
	"messages"
)

func main() {
	client := make(chan messages.Message)
	server := make(chan messages.Message)

	go func() {
		<-server
		// switch on message type
		client <- Pong{}
	}()

	server <- Ping{}
	<-client
	fmt.Println("Pong")
}