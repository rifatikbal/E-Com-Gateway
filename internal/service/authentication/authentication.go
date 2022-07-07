package authentication

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/rifatikbal/E-Com-Gateway/internal/service"
	"time"
)

type Authentication struct {
	ID        uint64
	Email     string
	SecretKey string
	Duration  time.Duration
}

type JWTClaim struct {
	ID    uint64 `json:"ID"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func New(id *uint64, email *string, secretKey *string, duration *time.Duration) service.AuthenticationSvc {
	return &Authentication{
		ID:        *id,
		Email:     *email,
		SecretKey: *secretKey,
		Duration:  *duration,
	}
}

func (auth *Authentication) NewToken() (*string, error) {
	expirationTime := time.Now().Add(auth.Duration * time.Second)
	claims := JWTClaim{
		ID:    auth.ID,
		Email: auth.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES512, claims).SignedString([]byte(auth.SecretKey))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (auth *Authentication) ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.SecretKey), nil
	})

	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err := errors.New("could not parse claims")
		return err
	}

	if claims.Email != auth.Email || claims.ID != auth.ID {
		err := errors.New("unauthorised entity")
		return err
	}

	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		err := errors.New("token expired")
		return err
	}
	return nil
}
