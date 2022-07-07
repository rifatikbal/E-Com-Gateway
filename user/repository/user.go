package repository

import (
	"github.com/rifatikbal/E-Com-Gateway/domain"
	"github.com/rifatikbal/E-Com-Gateway/internal/conn"
)

type User struct {
	*conn.DB
}

func New(db *conn.DB) domain.UserRepo {
	return &User{
		db,
	}
}

func (u *User) Store(m *domain.User) error {
	if err := u.GormDB.Create(m).Error; err != nil {
		return err
	}
	return nil
}
