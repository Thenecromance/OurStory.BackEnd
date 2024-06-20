package router

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/gin-gonic/gin"
)

type Router struct {
	path   string
	method string

	middleWare gin.HandlersChain

	realHandler gin.HandlerFunc // the real handler of the router

	// control the router's active status
	active bool
}

func (r *Router) IsRESTFUL() bool {
	return false
}

func (r *Router) Enable() {
	r.active = true
}

func (r *Router) Disable() {
	r.active = false
}

func (r *Router) SetPath(path string) {
	r.path = path
}

func (r *Router) SetMethod(method string) {
	r.method = method
}

func (r *Router) SetMiddleWare(middleware gin.HandlersChain) {
	r.middleWare = middleware
}

func (r *Router) SetHandler(handler ...gin.HandlerFunc) {
	r.realHandler = handler[0]
}

func (r *Router) GetPath() string {
	return r.path
}

func (r *Router) GetMethod() string {
	return r.method
}

func (r *Router) GetMiddleWare() gin.HandlersChain {
	return r.middleWare
}

func (r *Router) GetHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{r.handler}
}

func (r *Router) handler(ctx *gin.Context) {
	if r.realHandler == nil || !r.active {
		DefaultHandler()(ctx)
	} else {
		r.realHandler(ctx)
	}
}

func NewDefaultRouter() Interface.Router {
	return &Router{
		realHandler: DefaultHandler(),
		active:      true,
	}
}

func NewRouter(path_, method_ string) Interface.Router {
	return &Router{
		path:   path_,
		method: method_,

		realHandler: DefaultHandler(),
		active:      true,
	}
}
