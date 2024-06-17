package Interface

import "github.com/gin-gonic/gin"

type RouterProxy interface {
	GetPath() string
	GetMethod() string
	IsRESTFUL() bool

	GetMiddleWare() gin.HandlersChain
	GetHandler() gin.HandlerFunc
}
