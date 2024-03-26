package gJWT

import (
	"github.com/Thenecromance/OurStories/base/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var (
	option Options
)

//==================Options for the token==================

type Options struct {
	ExpireTime time.Duration
	Key        string
	Claim      interface{}
}

type Option func(*Options)

// WithExpireTime can set the expire time for the token
func WithExpireTime(expireTime time.Duration) Option {
	return func(o *Options) {
		o.ExpireTime = expireTime
	}
}

// WithKey can set the secret for the token
func WithKey(key string) Option {
	return func(o *Options) {
		o.Key = key
	}
}

//=========================================================

type DemoUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Claim is a struct that will be used to store the user information
type Claim struct {
	UserInfo             DemoUser
	jwt.RegisteredClaims // embedded unmodified `jwt.RegisteredClaims`
}

type gJWT struct {
	options Options
	token   *jwt.Token
}

// SignedToken should only be called by the place where you want to sign the token to user
func SignedToken(arg interface{}) (string, error) {

	claim := Claim{
		UserInfo: arg.(DemoUser),
	}
	claim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(option.ExpireTime)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(option.Key))
}

// AuthToken will authenticate the token is valid or not
func AuthToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(option.Key), nil
	})
	return err
}

// UnauthorizedResponse is a helper function to return a unauthorized response (due to most situation is api auth, so we return 200)
func UnauthorizedResponse(ctx *gin.Context) {
	ctx.Abort()
	ctx.JSON(http.StatusOK, gin.H{
		"error": "Unauthorized operation",
	})
}

func New(opts ...Option) {
	for _, opt := range opts {
		opt(&option)
	}
}

// NewMiddleware is a middleware for gin to authenticate the token
func NewMiddleware(opts ...Option) gin.HandlerFunc {

	for _, opt := range opts {
		opt(&option)
	}

	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization") // get the token from the header
		if len(auth) == 0 {
			UnauthorizedResponse(ctx)
			return
		}
		err := AuthToken(auth)
		if err != nil {
			UnauthorizedResponse(ctx)
		}
		logger.Get().Debug("auth success")
		ctx.Next()
	}
}
