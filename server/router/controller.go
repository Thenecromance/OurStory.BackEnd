package router

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/gin-gonic/gin"
)

type controller struct {
	gin   *gin.Engine
	proxy map[string]Interface.RouterProxy
}

func (c *controller) RegisterRouter(routerProxy Interface.RouterProxy) error {
	//TODO implement me
	panic("implement me")
}

func (c *controller) Close() error {
	//TODO implement me
	panic("implement me")
}

func (c *controller) ApplyRouter() error {
	for _, proxy := range c.proxy {
		if proxy.IsRESTFUL() {
			restful := proxy.(*REST)
			c.gin.GET(restful.Path, restful.GETHandler)
			c.gin.POST(restful.Path, restful.POSTHandler)
			c.gin.PUT(restful.Path, restful.PUTHandler)
			c.gin.DELETE(restful.Path, restful.DELETEHandler)
		} else {
			router := proxy.(*Router)
			c.gin.Handle(router.Method, router.Path, router.Handler)
		}

	}
}

func NewController() Interface.RouterController {
	return &controller{}
}
