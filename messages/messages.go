package messages

type BaseMessage struct{
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

// using some kind of global map of node ids to chan and calling Message.send(from, to) would be nice