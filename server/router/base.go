package router

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/gin-gonic/gin"
)

type Router struct {
	path   string
	method string

	middleWare  gin.HandlersChain
	realHandler gin.HandlerFunc

	available bool
}

func (r *Router) IsRESTFUL() bool {
	return false
}

func (r *Router) Enable() {
	r.available = true
}

func (r *Router) Disable() {
	r.available = false
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

func (r *Router) handler(c *gin.Context) {
	if r.realHandler == nil || !r.available {
		defaultFunc(c)
	} else {
		r.realHandler(c)
	}
}

func NewRouter() Interface.Router {
	return &Router{}
}

type handler struct {
	h gin.HandlerFunc
}
