package repository

import (
	"github.com/rifatikbal/E-Com-Gateway/domain"
	"github.com/rifatikbal/E-Com-Gateway/internal/conn"
)

type Token struct {
	*conn.DB
}

func New(db *conn.DB) domain.TokenRepo {
	return &Token{
		db,
	}
}

func (tkn *Token) Store(m *domain.Token) error {
	if err := tkn.GormDB.Create(m).Error; err != nil {
		return err
	}
	return nil
}
