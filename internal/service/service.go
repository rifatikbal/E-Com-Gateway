package service

type AuthenticationSvc interface {
	NewToken() (*string, error)
	ValidateToken(signedToken string) error
}
