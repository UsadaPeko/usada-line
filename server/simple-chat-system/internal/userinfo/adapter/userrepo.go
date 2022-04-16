package adapter

import (
	"fmt"
	"github.com/UsadaPeko/usadaline/simplechatsystem/internal/userinfo/domain"
)

type InMemoryUserRepository struct {
	cache map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{cache: map[string]*domain.User{}}
}

func (i *InMemoryUserRepository) Get(id string) (*domain.User, error) {
	user, ok := i.cache[id]
	if !ok {
		return nil, fmt.Errorf("not found user(%v)", id)
	}
	return user, nil
}

func (i *InMemoryUserRepository) Save(user *domain.User) error {
	i.cache[user.ID] = user
	return nil
}
