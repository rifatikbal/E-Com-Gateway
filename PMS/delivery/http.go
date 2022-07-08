package delivery

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rifatikbal/E-Com-Gateway/internal/config"
	svc "github.com/rifatikbal/E-Com-Gateway/internal/service"
	"log"
	"net/http"
)

type PMSHandler struct {
	pmsSVc  svc.PMSSvc
	authSvc svc.AuthenticationSvc
	pmsCfg  config.PMSCfg
}

func New(r *chi.Mux, pmsSvc svc.PMSSvc, authSvc svc.AuthenticationSvc, pmsCfg config.PMSCfg) {
	handler := &PMSHandler{
		pmsSVc:  pmsSvc,
		authSvc: authSvc,
		pmsCfg:  pmsCfg,
	}

	r.Group(func(r chi.Router) {

		r.Route("/pms", func(r chi.Router) {
			r.Post("/products/{product_type}", handler.GetProducts)
		})
	})
}

func (p *PMSHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		log.Println("Couldn't parse token")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "coudn't parse token",
		})

		return
	}

	_, err = p.authSvc.ValidateToken(token.String())
	if err != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{
			"error": "unauthorized",
		})

		return
	}

	http.Redirect(w, r, p.pmsCfg.Url, http.StatusFound)
}
