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

func (u *User) GetUser(ctr *domain.UserCriteria) (*domain.User, error) {
	var user *domain.User
	q := u.GormDB.
		Table("users").
		Where(
			"email = ?",
			*ctr.Email,
		)

	if err := q.Scan(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
