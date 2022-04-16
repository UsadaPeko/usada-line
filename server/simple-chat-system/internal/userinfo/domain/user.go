package domain

import "github.com/google/uuid"

type User struct {
	ID   string
	Name string
}

type UserRepository interface {
	Get(id string) (*User, error)
	Save(user *User) error
}

func CreateNewUser(name string) *User {
	return &User{
		ID:   uuid.NewString(),
		Name: name,
	}
}
