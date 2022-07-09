package delivery

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rifatikbal/E-Com-Gateway/internal/config"
	svc "github.com/rifatikbal/E-Com-Gateway/internal/service"
	"log"
	"net/http"
)

type PMSHandler struct {
	authSvc svc.AuthenticationSvc
	pmsCfg  config.PMSCfg
}

func New(r *chi.Mux, authSvc svc.AuthenticationSvc, pmsCfg config.PMSCfg) {
	handler := &PMSHandler{
		authSvc: authSvc,
		pmsCfg:  pmsCfg,
	}

	r.Group(func(r chi.Router) {

		r.Route("/pms", func(r chi.Router) {
			r.Get("/products/{name}", handler.GetProducts)
			r.Post("/products", handler.StoreProducts)
		})
	})
}

func (p *PMSHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		for _, c := range r.Cookies() {
			fmt.Println(c, " this is not what you think it is ")
		}

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})

		return
	}

	_, err = p.authSvc.ValidateToken(token.Value)
	if err != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{
			"error": "unauthorized",
		})

		return
	}

	http.Redirect(w, r, fmt.Sprintf(p.pmsCfg.Url+"/api/product/%v", chi.URLParam(r, "name")), http.StatusPermanentRedirect)
}

func (p *PMSHandler) StoreProducts(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		log.Println("Couldn't parse token")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "coudn't parse token",
		})

		return
	}

	_, err = p.authSvc.ValidateToken(token.Value)
	if err != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{
			"error": "unauthorized",
		})

		return
	}

	http.Redirect(w, r, p.pmsCfg.Url+"/api/product", http.StatusPermanentRedirect)
}
