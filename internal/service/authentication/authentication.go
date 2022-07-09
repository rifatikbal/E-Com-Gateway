package authentication

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rifatikbal/E-Com-Gateway/domain/dto"
	"github.com/rifatikbal/E-Com-Gateway/internal/service"
	"log"
	"time"
)

type Authentication struct {
	ID        uint64
	Email     string
	SecretKey string
	Duration  time.Duration
}

func New(id *uint64, email *string, secretKey *string, duration *time.Duration) service.AuthenticationSvc {
	authentication := Authentication{}

	if id != nil {
		authentication.ID = *id
	}
	if email != nil {
		authentication.Email = *email
	}
	if secretKey != nil {
		authentication.SecretKey = *secretKey
	}
	if duration != nil {
		authentication.Duration = *duration
	}

	return &authentication
}

func (auth *Authentication) NewToken() (*string, error) {
	expirationTime := time.Now().Add(auth.Duration)
	claims := dto.JWTClaim{
		Email: auth.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	log.Println(claims)

	log.Println(expirationTime)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(auth.SecretKey))
	if err != nil {
		log.Println("failed to log in here: ", err)
		return nil, err
	}

	return &token, nil
}

func (auth *Authentication) ValidateToken(signedToken string) (*dto.JWTClaim, error) {
	claims := &dto.JWTClaim{}
	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.SecretKey), nil
	})

	if err != nil {
		log.Println("buggy: ", err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(*dto.JWTClaim)
	if !ok {
		log.Println("xyz")
		err := errors.New("could not parse claims")
		return nil, err
	}

	fmt.Println("buggybot ", claims.Email, " ", claims.ExpiresAt)

	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		log.Println("here")
		err := errors.New("token expired")
		return nil, err
	}
	return claims, nil
}
