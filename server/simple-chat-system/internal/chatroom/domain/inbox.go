package domain

type UserInbox struct {
	Owner *User

	Messages []Message
}

func MakeUserInbox(owner *User) *UserInbox {
	return &UserInbox{Owner: owner, Messages: []Message{}}
}

func (ui *UserInbox) PutMessage(sender *User, chatRoomID, messageText string) {
	message := Message{
		ChatRoomID:       chatRoomID,
		From:             sender,
		LamportTimestamp: 0, // TODO
		Text:             messageText,
	}
	ui.Messages = append(ui.Messages, message)
}

type InboxRepository interface {
	Get(*User) (*UserInbox, error)
	Save(inbox *UserInbox) error
}

type Message struct {
	ChatRoomID       string
	From             *User
	LamportTimestamp int64

	Text string
}
