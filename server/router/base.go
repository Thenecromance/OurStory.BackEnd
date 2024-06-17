package router

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Path   string
	Method string

	MiddleWare gin.HandlersChain
	Handler    gin.HandlerFunc
}

func (r *Router) GetPath() string {
	return r.Path
}

func (r *Router) GetMethod() string {
	return r.Method
}

func (r *Router) IsRESTFUL() bool {
	return false
}

func (r *Router) GetMiddleWare() gin.HandlersChain {
	return r.MiddleWare
}

func (r *Router) GetHandler() gin.HandlerFunc {
	return r.Handler
}

func New(path, method string) Interface.RouterProxy {
	return &Router{
		Path:   path,
		Method: method,

		Handler: defaultFunc,
	}
}
