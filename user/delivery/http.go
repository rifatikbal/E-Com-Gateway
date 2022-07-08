package delivery

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rifatikbal/E-Com-Gateway/domain"
	"github.com/rifatikbal/E-Com-Gateway/domain/dto"
	"github.com/rifatikbal/E-Com-Gateway/internal/config"
	authSvc "github.com/rifatikbal/E-Com-Gateway/internal/service/authentication"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

const (
	hashCost = 14
	jwtKey   = "supersecretkey"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
	authCfg     config.AuthCfg
}

func New(r *chi.Mux, userUseCase domain.UserUseCase, authCfg config.AuthCfg) {
	handler := UserHandler{
		userUseCase: userUseCase,
		authCfg:     authCfg,
	}
	log.Println(authCfg)
	r.Group(func(r chi.Router) {

		r.Route("/user", func(r chi.Router) {
			r.Post("/register", handler.registerUser)
			r.Post("/login", handler.loginUser)

		})
	})
}

func (u *UserHandler) registerUser(w http.ResponseWriter, r *http.Request) {
	userReq := domain.User{}

	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}
	log.Println(userReq.Password)
	hashedPassword, err := generatePasswordHash(userReq.Password)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}

	log.Println(*hashedPassword)

	user := &domain.User{
		ID:       userReq.ID,
		Name:     userReq.Name,
		Email:    userReq.Email,
		Phone:    userReq.Phone,
		Address:  userReq.Address,
		Password: *hashedPassword,
	}

	if err := u.userUseCase.Store(user); err != nil {
		log.Println(err)

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return

	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"success": "true",
	})
}

func (u *UserHandler) loginUser(w http.ResponseWriter, r *http.Request) {
	loginReq := dto.LoginUserReq{}
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})

		return
	}

	ctr := &domain.UserCriteria{
		Email: &loginReq.Email,
	}

	user, err := u.userUseCase.GetUser(ctr)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{
			"error": "unauthorized user ",
		})

		return
	}

	auth := authSvc.New(&user.ID, &user.Email, &u.authCfg.Secret, &u.authCfg.Duration)
	log.Println(auth)

	token, err := auth.NewToken()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to generate token",
		})

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   *token,
		Expires: time.Now().Add(u.authCfg.Duration),
	})

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"success": "true",
	})
}

func generatePasswordHash(password string) (*string, error) {
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(hashedPasswordByte)

	return &hashedPassword, nil
}
