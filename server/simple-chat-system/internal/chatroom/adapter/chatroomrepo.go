package adapter

import (
	"fmt"
	"github.com/UsadaPeko/usadaline/simplechatsystem/internal/chatroom/domain"
)

type InMemoryChatRoomRepository struct {
	cache map[string]*domain.Room
}

func NewInMemoryChatRoomRepository() *InMemoryChatRoomRepository {
	return &InMemoryChatRoomRepository{
		cache: map[string]*domain.Room{},
	}
}

func (i *InMemoryChatRoomRepository) Get(roomID string) (*domain.Room, error) {
	room, ok := i.cache[roomID]
	if !ok {
		return nil, fmt.Errorf("not found chat room(%v)", roomID)
	}
	return room, nil
}

func (i *InMemoryChatRoomRepository) Save(room *domain.Room) error {
	i.cache[room.ID] = room
	return nil
}

func (i *InMemoryChatRoomRepository) MyRooms(user *domain.User) ([]*domain.Room, error) {
	var myRooms []*domain.Room
	for key, v := range i.cache {
		if v.OwnerUser.ID == user.ID {
			myRooms = append(myRooms, i.cache[key])
		}
	}
	return myRooms, nil
}
