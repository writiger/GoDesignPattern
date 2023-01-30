package mediator

import "fmt"

type Country interface {
	SendMessage(message string)
	GetMessage(message string)
}
type USA struct {
	UnitedNations
}

func (usa USA) SendMessage(message string) {
	usa.UnitedNations.ForwardMessage(message, usa)
}

func (usa USA) GetMessage(message string) {
	fmt.Printf("美国收到对方消息: %s\n", message)
}

type Iraq struct {
	UnitedNations
}

func (iraq Iraq) SendMessage(message string) {
	iraq.UnitedNations.ForwardMessage(message, iraq)
}

func (iraq Iraq) GetMessage(message string) {
	fmt.Printf("伊拉克收到对方消息: %s\n", message)
}
