package domain

import (
	"fmt"
)

type MessageBroadcasting struct {
	ir InboxRepository
	rm MessageBroadcastRetryManager
}

func NewMessageBroadcasting(ir InboxRepository, rm MessageBroadcastRetryManager) *MessageBroadcasting {
	return &MessageBroadcasting{ir: ir, rm: rm}
}

func (mb *MessageBroadcasting) Broadcast(room *Room, user *User, messageText string) error {
	isAvailable := room.AvailableToSendMessage(user)
	if !isAvailable {
		return fmt.Errorf("user(%v) send message to invalid room(%v)", user.ID, room.ID)
	}

	for _, v := range room.ParticipateUsers {
		inbox, err := mb.ir.Get(v)
		if err != nil {
			mb.rm.RetryBackgroundTask(room.ID, v, messageText)
		}

		inbox.PutMessage(user, room.ID, messageText)

		err = mb.ir.Save(inbox)
		if err != nil {
			mb.rm.RetryBackgroundTask(room.ID, v, messageText)
		}
	}
	return nil
}
