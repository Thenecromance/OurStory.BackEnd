package user

import (
	"fmt"
	"github.com/Thenecromance/OurStories/backend/UserV2/data"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var (
	KEY = "secret"
)

// Claim is a struct that will be used to store the user information
type Claim struct {
	UserInfo             data.UserClaim `json:"info"` // the user information
	jwt.RegisteredClaims                // embedded unmodified `jwt.RegisteredClaims`
}

func getRealToken(token string) string {
	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	token = splitToken[1]
	return token
}

func (auth *Authorization) ValidByToken(token string) bool {
	token = getRealToken(token)
	if auth.tokenCache.hasToken(token) {
		auth.tokenCache.storeToken(token)
	}

	_, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})
	return err == nil
}

func (auth *Authorization) SignedToken(UsrClaim data.UserClaim) string {
	claim := Claim{
		UserInfo: UsrClaim,
	}
	claim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte(KEY))
	log.Debug("token: ", tokenString)
	return tokenString
}

func (auth *Authorization) ParseToken(token string) (*data.UserClaim, error) {
	token = getRealToken(token)
	if auth.tokenCache.hasToken(token) {
		return nil, nil
	}

	obj, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(KEY), nil
	})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if claim, ok := obj.Claims.(*Claim); ok {
		log.Debug("parse token success")
		return &claim.UserInfo, nil
	}
	log.Error("parse token failed")
	return nil, nil
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
