package Interface

import "github.com/gin-gonic/gin"

type IMiddleWareController interface {
	RegisterMiddleWare(name string, handler gin.HandlerFunc)
	GetMiddleWare(name string) (gin.HandlerFunc, error)
}
