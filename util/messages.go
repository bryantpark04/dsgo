package util

type BaseMessage struct {
	sender string
}

func (m BaseMessage) Sender() string { return m.sender }

func BaseMessageFrom(sender string) BaseMessage { return BaseMessage{sender} }

type Message interface {
	Sender() string
}

func Send[Msg Message](to chan Msg, message Msg) {
	to <- message
}

type Directory map[string]chan Message

// using some kind of global map of node ids to chan and calling Message.send(from, to) would be nice
