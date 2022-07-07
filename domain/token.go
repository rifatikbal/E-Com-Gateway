package domain

type Token struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

type TokenRepo interface {
	Store(m *Token) error
}

type TokenUseCase interface {
}
