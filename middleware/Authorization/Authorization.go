package Authorization

import "github.com/gin-gonic/gin"

type IAuth interface {
	AuthorizeUser(claim interface{}) (string, error)

	AuthorizeToken(token string) (interface{}, error)

	TokenExpired(token string) (bool, error)

	RefreshToken(oldToken string) (string, error)

	MarkTokenExpired(token string) error

	Middleware() gin.HandlerFunc
}
