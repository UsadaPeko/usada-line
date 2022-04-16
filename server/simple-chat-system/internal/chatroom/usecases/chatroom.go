package usecases

import "github.com/UsadaPeko/usadaline/simplechatsystem/internal/chatroom/domain"

type ChatRoomUseCases struct {
	rr domain.RoomRepository
	ir domain.InboxRepository
	bc *domain.MessageBroadcasting
}

func NewChatRoomUseCases(
	rr domain.RoomRepository,
	ir domain.InboxRepository,
	bc *domain.MessageBroadcasting,
) *ChatRoomUseCases {
	return &ChatRoomUseCases{rr: rr, ir: ir, bc: bc}
}

func (cu *ChatRoomUseCases) CreateNewChatRoom(userID, roomName string) (*domain.Room, error) {
	newChatRoom := domain.NewRoom(roomName, domain.NewUser(userID))
	err := cu.rr.Save(newChatRoom)
	return newChatRoom, err
}

func (cu *ChatRoomUseCases) MyRooms(userID string) ([]*domain.Room, error) {
	rooms, err := cu.rr.MyRooms(domain.NewUser(userID))
	return rooms, err
}

func (cu *ChatRoomUseCases) InviteChatRoom(inviter, invitee *domain.User, chatRoomID string) error {
	chatRoom, err := cu.rr.Get(chatRoomID)
	if err != nil {
		return err
	}

	err = chatRoom.Invite(inviter, invitee)
	if err != nil {
		return err
	}

	return cu.rr.Save(chatRoom)
}

func (cu *ChatRoomUseCases) SendMessage(user *domain.User, chatRoomID, message string) error {
	chatRoom, err := cu.rr.Get(chatRoomID)
	if err != nil {
		return err
	}
	return cu.bc.Broadcast(chatRoom, user, message)
}

func (cu *ChatRoomUseCases) DetectNewUser(user *domain.User) error {
	inbox := domain.MakeUserInbox(user)
	return cu.ir.Save(inbox)
}
