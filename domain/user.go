package domain

type User struct {
	ID       uint64 `gorm:"primary_key; auto_increment;" json:"id"`
	Name     string `gorm:"not null;type:varchar(255);"json:"name"`
	Email    string `gorm:"not null;type:varchar(255);uniqueIndex" json:"email"`
	Phone    string `gorm:"not null;type:varchar(255);" json:"phone"`
	Address  string `gorm:"not null;type:varchar(255);" json:"address"`
	Password string `gorm:"not null;type:varchar(255);" json:"password"`
}

type UserCriteria struct {
	Email *string
}

type UserRepo interface {
	Store(m *User) error
	GetUser(ctr *UserCriteria) (*User, error)
}

type UserUseCase interface {
	Store(m *User) error
	GetUser(ctr *UserCriteria) (*User, error)
}
