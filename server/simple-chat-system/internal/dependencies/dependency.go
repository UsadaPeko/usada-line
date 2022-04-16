package dependencies

import (
	chatRoomAdapter "github.com/UsadaPeko/usadaline/simplechatsystem/internal/chatroom/adapter"
	chatRoomDomain "github.com/UsadaPeko/usadaline/simplechatsystem/internal/chatroom/domain"
	chatRoomUseCases "github.com/UsadaPeko/usadaline/simplechatsystem/internal/chatroom/usecases"

	userInfoAdapter "github.com/UsadaPeko/usadaline/simplechatsystem/internal/userinfo/adapter"
	userInfoUseCases "github.com/UsadaPeko/usadaline/simplechatsystem/internal/userinfo/usecases"
)

var d *Dependencies = nil

func init() {
	d = NewDependencies()
}

func Use() *Dependencies {
	if d == nil {
		d = NewDependencies()
	}
	return d
}

type Dependencies struct {
	cuc *chatRoomUseCases.ChatRoomUseCases
	uuc *userInfoUseCases.UserInfoUseCases
}

func NewDependencies() *Dependencies {
	// Chat Room
	roomRepo := chatRoomAdapter.NewInMemoryChatRoomRepository()
	inboxRepo := chatRoomAdapter.NewInMemoryInboxRepository()
	retryManager := chatRoomAdapter.NewLogRetryManager()
	broadcasting := chatRoomDomain.NewMessageBroadcasting(inboxRepo, retryManager)

	cuc := chatRoomUseCases.NewChatRoomUseCases(roomRepo, inboxRepo, broadcasting)

	// User Info
	userInfoRepo := userInfoAdapter.NewInMemoryUserRepository()

	uuc := userInfoUseCases.NewUserInfoUseCases(userInfoRepo)

	// Build
	return &Dependencies{
		cuc: cuc,
		uuc: uuc,
	}
}

func (d *Dependencies) ChatRoomUseCases() *chatRoomUseCases.ChatRoomUseCases {
	return d.cuc
}

func (d *Dependencies) UserInfoUseCases() *userInfoUseCases.UserInfoUseCases {
	return d.uuc
}
