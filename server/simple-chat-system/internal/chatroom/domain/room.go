package domain

import "fmt"

type Room struct {
	ID   string
	Name string

	OwnerUser *User

	ParticipateUsers []*User
}

func (r *Room) Invite(inviter *User, invitee *User) error {
	if r.OwnerUser.ID != inviter.ID {
		return fmt.Errorf("user %v can't invite other user %v", inviter.ID, invitee.ID)
	}
	r.ParticipateUsers = append(r.ParticipateUsers, invitee)
	return nil
}

func (r *Room) AvailableToSendMessage(user *User) (available bool) {
	available = false
	for _, v := range r.ParticipateUsers {
		if v.ID == user.ID {
			available = true
		}
	}
	return
}

func NewRoom(name string, ownerUser *User) *Room {
	return &Room{Name: name, OwnerUser: ownerUser}
}

type RoomRepository interface {
	Get(roomID string) (*Room, error)
	Save(*Room) error
	MyRooms(user *User) ([]*Room, error)
}
