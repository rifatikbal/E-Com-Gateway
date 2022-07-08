package dto

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	ID    uint64 `json:"ID"`
	Email string `json:"email"`
	jwt.StandardClaims
}
