package usecase

import "github.com/rifatikbal/E-Com-Gateway/domain"

type User struct {
	userRepo domain.UserRepo
}

func New(ur domain.UserRepo) domain.UserUseCase {
	return &User{
		userRepo: ur,
	}
}

func (u *User) Store(m *domain.User) error {
	if err := u.userRepo.Store(m); err != nil {
		return err
	}
	return nil
}
