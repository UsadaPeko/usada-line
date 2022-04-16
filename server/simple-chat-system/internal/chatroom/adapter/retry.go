package adapter

import (
	"github.com/UsadaPeko/usadaline/simplechatsystem/internal/chatroom/domain"
	"log"
)

type LogRetryManager struct {
}

func NewLogRetryManager() *LogRetryManager {
	return &LogRetryManager{}
}

func (l LogRetryManager) RetryBackgroundTask(roomID string, user *domain.User, message string) {
	log.Printf("fail to send message. room(%v) -> user(%v): %v", roomID, user.ID, message)
}
