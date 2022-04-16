package adapter

import (
	"fmt"
	"github.com/UsadaPeko/usadaline/simplechatsystem/internal/chatroom/domain"
)

type InMemoryInboxRepository struct {
	cache map[string]*domain.UserInbox
}

func NewInMemoryInboxRepository() *InMemoryInboxRepository {
	return &InMemoryInboxRepository{
		cache: map[string]*domain.UserInbox{},
	}
}

func (i *InMemoryInboxRepository) Get(user *domain.User) (*domain.UserInbox, error) {
	inbox, ok := i.cache[user.ID]
	if !ok {
		return nil, fmt.Errorf("not found user(%v) inbox", user.ID)
	}
	return inbox, nil
}

func (i *InMemoryInboxRepository) Save(inbox *domain.UserInbox) error {
	i.cache[inbox.Owner.ID] = inbox
	return nil
}
