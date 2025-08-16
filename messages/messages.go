package messages

type Message interface{}

func Send(to chan Message, message Message) {
	to <- message
}

// using some kind of global map of node ids to chan and calling Message.send(from, to) would be nice