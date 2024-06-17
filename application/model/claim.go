package model

import "github.com/golang-jwt/jwt/v5"

var (
	KEY = "secret"
)

// Claim is a struct that will be used to store the user information
type Claim struct {
	UserInfo             UserClaim `json:"info"` // the user information
	jwt.RegisteredClaims           // embedded unmodified `jwt.RegisteredClaims`
}
