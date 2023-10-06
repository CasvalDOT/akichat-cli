package presenters

import (
	"fmt"

	core "github.com/CasvalDOT/akichat-core"
)

type messagesPresenter struct {
	messages []core.Message
}

func (m *messagesPresenter) Print() {
	for _, message := range m.messages {
		fmt.Println(message.Author.Name + " - " + message.Time)
		fmt.Println(message.Content)
	}
}

func NewMessagesPresenter(messages []core.Message) Presenter {
	return &messagesPresenter{
		messages,
	}
}
