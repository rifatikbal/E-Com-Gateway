package domain

type User struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UserRepo interface {
}

type UserUseCase interface {
}
