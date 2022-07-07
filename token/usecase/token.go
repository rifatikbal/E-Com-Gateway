package usecase

import "github.com/rifatikbal/E-Com-Gateway/domain"

type Token struct {
	tokenRepo domain.TokenRepo
}

func New(ur domain.TokenRepo) domain.TokenUseCase {
	return &Token{
		tokenRepo: ur,
	}
}

func (u *Token) Store(m *domain.Token) error {
	if err := u.tokenRepo.Store(m); err != nil {
		return err
	}
	return nil
}
