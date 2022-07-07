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

func (u *User) GetUser(ctr *domain.UserCriteria) (*domain.User, error) {
	user, err := u.userRepo.GetUser(ctr)
	if err != nil {
		return nil, err
	}

	return user, nil

}
