package delivery

import (
	"github.com/go-chi/chi/v5"
	"github.com/rifatikbal/E-Com-Gateway/domain"
	"net/http"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func New(r *chi.Mux, userUseCase domain.UserUseCase) {
	handler := UserHandler{
		userUseCase: userUseCase,
	}
	r.Group(func(r chi.Router) {
		r.Post("/users", handler.CreateUserHandler)

	})
}

func (u *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

}
