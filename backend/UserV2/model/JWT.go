package model

import (
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var (
	KEY = "secret"
)

// Claim is a struct that will be used to store the user information
type Claim struct {
	UserInfo             interface{} `json:"info"` // the user information
	jwt.RegisteredClaims                           // embedded unmodified `jwt.RegisteredClaims`
}

func (auth *Authorization) ValidByToken(token string) bool {

	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return false
	}
	token = splitToken[1]

	if auth.tokenCache.hasToken(token) {
		auth.tokenCache.storeToken(token)
	}

	_, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})
	return err == nil
}

func (auth *Authorization) SignedToken(UsrClaim interface{}) string {
	claim := Claim{
		UserInfo: UsrClaim,
	}
	claim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte(KEY))
	return tokenString
}

func (auth *Authorization) userInfo(token string) interface{} {
	claim := &Claim{}
	_, _ = jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})
	return claim.UserInfo
}

func (auth *Authorization) storeToken(token string) {
	auth.tokenCache.storeToken(token)
}

func (auth *Authorization) markTokenExpired(token string) {
	auth.tokenCache.markTokenExpired(token)
}
