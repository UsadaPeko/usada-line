package domain

type User struct {
	ID string
}

func NewUser(ID string) *User {
	return &User{ID: ID}
}
