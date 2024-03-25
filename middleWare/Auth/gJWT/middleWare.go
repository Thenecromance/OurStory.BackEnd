package gJWT

import (
	"github.com/Thenecromance/OurStories/base/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	tokenExpireTime = time.Hour * time.Duration(1) // 1 hour
	secret          = "secert.www.ourstories.com"
)

// when user login, we will sign a token for him
func SignTokenToUser(c *gin.Context, usr string, uid int) {
	//sign a new token to user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim(usr, uid))
	tokenStr, err := token.SignedString([]byte(secret))

	// if sign token failed, return 401
	if err != nil {
		responseUnauthorized(c, "Sign token failed")
	}

	// using Set-Cookie to store token in user client
	StoreInCookie(c, tokenStr)
}

// add Set-Cookie to store token in user client
func StoreInCookie(c *gin.Context, token string) {
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
}

// auth token sequence
func authUserTokenIsValid(c *gin.Context) {

	// get token from header
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		logger.Get().Errorf("%s token is empty ", c.Request.RemoteAddr)
		responseUnauthorized(c, "token is empty")
		return
	}

	//start to parse token and check if it is valid
	token_, err := jwt.ParseWithClaims(tokenStr, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || token_.Valid {
		responseUnauthorized(c, "Invalid token")
		return
	}

	claims, ok := token_.Claims.(*JWTClaim)
	if !ok {
		responseUnauthorized(c, "Invalid token")
		return
	}

	if claims.ExpiresAt.Unix() <= time.Now().Unix() {
		responseUnauthorized(c, "Token is expired")
		return
	}

	//allow request to continue
	c.Next()

}

// middleware to check if token is valid
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authUserTokenIsValid(c)
	}
}
