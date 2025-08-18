package main

import "messages"

type Ping struct {
	messages.BaseMessage
}

type Pong struct {
	messages.BaseMessage
}