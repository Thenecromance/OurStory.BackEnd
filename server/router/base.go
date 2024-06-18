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

// real router
type ProxyRouter struct {
	Router *Router
}

func (p *ProxyRouter) GetPath() string {
	//return p.Router.GetPath()
	if p.Router == nil {
		return ""
	}
	return p.Router.GetPath()
}

func (p *ProxyRouter) GetMethod() string {

	if p.Router == nil {
		return ""
	}

	return p.Router.GetMethod()
}

func (p *ProxyRouter) IsRESTFUL() bool {
	if p.Router == nil {
		return false
	}
	return p.Router.IsRESTFUL()
}

func (p *ProxyRouter) GetMiddleWare() gin.HandlersChain {
	return p.Router.GetMiddleWare()
}

func (p *ProxyRouter) GetHandler() gin.HandlerFunc {
	return p.Router.GetHandler()
}

func NewProxyRouter(router *Router) Interface.RouterProxy {
	return &ProxyRouter{
		Router: router,
	}
}
