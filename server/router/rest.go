package router

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/gin-gonic/gin"
)

type rest struct {
	path string

	middleWare gin.HandlersChain

	handlers []gin.HandlerFunc // GET ,POST , PUT , DELETE

	available bool
}

func (r *rest) getHandler(c *gin.Context) {
	if r.handlers[0] != nil || !r.available {
		r.handlers[0](c)
	} else {
		r.handlers[0](c)
	}
}
func (r *rest) postHandler(c *gin.Context) {
	if r.handlers[1] != nil || !r.available {
		r.handlers[1](c)
	} else {
		r.handlers[1](c)
	}
}
func (r *rest) putHandler(c *gin.Context) {
	if r.handlers[2] != nil || !r.available {
		r.handlers[2](c)
	} else {
		r.handlers[2](c)
	}
}
func (r *rest) deleteHandler(c *gin.Context) {
	if r.handlers[3] != nil || !r.available {
		r.handlers[3](c)
	} else {
		r.handlers[3](c)
	}
}

func (r *rest) IsRESTFUL() bool {
	return true
}

func (r *rest) Enable() {
	r.available = true
}

func (r *rest) Disable() {
	r.available = false
}

func (r *rest) SetPath(path string) {
	r.path = path
}

func (r *rest) SetMethod(method string) {
	// empty method for rest
}

func (r *rest) SetMiddleWare(middleware gin.HandlersChain) {
	r.middleWare = middleware
}

func (r *rest) SetHandler(handler []gin.HandlerFunc) {
	for i, h := range handler {
		if h == nil {
			continue
		} else {
			r.handlers[i] = h
		}
	}
}

func (r *rest) GetPath() string {
	return r.path
}

func (r *rest) GetMethod() string {
	// empty method for rest
	return ""
}

func (r *rest) GetMiddleWare() gin.HandlersChain {
	return r.middleWare
}

func (r *rest) GetHandler() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		r.getHandler,
		r.postHandler,
		r.putHandler,
		r.deleteHandler,
	}
}

func NewREST() Interface.Router {
	return &rest{}
}
