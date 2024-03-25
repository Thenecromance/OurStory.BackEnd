package gJWT

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type token struct {
	User string `json:"usr"`
	Uid  int    `json:"uid"`
}

type JWTClaim struct {
	token
	jwt.RegisteredClaims
}

func Claim(user string, uid int) JWTClaim {
	claims := JWTClaim{
		token{
			User: user,
			Uid:  uid,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return claims
}
