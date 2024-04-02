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

// Claim is a struct that will be used to store the user information
type Claim struct {
	UserInfo             interface{} `json:"info"` // the user information
	jwt.RegisteredClaims                           // embedded unmodified `jwt.RegisteredClaims`
}

type gJWT struct {
	options Options
	token   *jwt.Token
}

// SignedToken should only be called by the place where you want to sign the token to user
func SignedToken(ctx *gin.Context, arg interface{}) (string, error) {

	claim := Claim{
		UserInfo: arg,
	}
	claim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(option.ExpireTime)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(option.Key))
	ctx.SetCookie("Authorization", tokenString, 1296000, "/", "localhost:8080", false, false)
	return tokenString, err
}

// AuthToken will authenticate the token is valid or not if return not nil,means the token is invalid otherwise it is valid
func AuthToken(tokenString string) error {
	_, err := GetObjectFromToken(tokenString)
	return err
}

func AuthorizeToken(tokenString string) bool {
	_, err := GetObjectFromToken(tokenString)
	if err != nil {
		return false
	}
	return true
}

// GetObjectFromToken will return the object from the token, if the token is invalid, it will return an error and the object will be nil
func GetObjectFromToken(tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(option.Key), nil
	})

	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*Claim)
	if !ok {
		return nil, err
	}
	return claim.UserInfo, nil
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
	New(opts...)
	return func(ctx *gin.Context) {
		auth, err := ctx.Cookie("Authorization")

		if len(auth) == 0 || err != nil {
			UnauthorizedResponse(ctx)
			return
		}
		err = AuthToken(auth)
		if err != nil {
			UnauthorizedResponse(ctx)
		}
		logger.Get().Debug("auth success")
		ctx.Next()
	}
}
