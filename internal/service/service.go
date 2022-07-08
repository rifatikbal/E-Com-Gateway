package service

import "github.com/rifatikbal/E-Com-Gateway/domain/dto"

type AuthenticationSvc interface {
	NewToken() (*string, error)
	ValidateToken(signedToken string) (*dto.JWTClaim, error)
}

type PMSSvc interface {
}
