package delivery

import (
	"github.com/go-chi/chi/v5"
	"github.com/rifatikbal/E-Com-Gateway/domain"
	"net/http"
)

type TokenHandler struct {
	tokenUseCase domain.TokenUseCase
}

func New(r *chi.Mux, TokenUseCase domain.TokenUseCase) {
	handler := TokenHandler{
		tokenUseCase: TokenUseCase,
	}
	r.Group(func(r chi.Router) {
		r.Post("/tokens", handler.CreateTokenHandler)

	})
}

func (u *TokenHandler) CreateTokenHandler(w http.ResponseWriter, r *http.Request) {

}
