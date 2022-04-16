package usecases

import "github.com/UsadaPeko/usadaline/simplechatsystem/internal/userinfo/domain"

type UserInfoUseCases struct {
	ur domain.UserRepository
}

func NewUserInfoUseCases(ur domain.UserRepository) *UserInfoUseCases {
	return &UserInfoUseCases{ur: ur}
}

func (uc *UserInfoUseCases) GetUserInfo(userID string) (*domain.User, error) {
	user, err := uc.ur.Get(userID)
	return user, err
}

func (uc *UserInfoUseCases) CreateNewUser(userName string) (*domain.User, error) {
	user := domain.CreateNewUser(userName)
	err := uc.ur.Save(user)
	return user, err
}
