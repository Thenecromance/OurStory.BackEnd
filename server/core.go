package server

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/gin-gonic/gin"
)

type core struct {
	routerController      Interface.RouterController
	middlerWareController Interface.MiddleWareController // all middlewares will be registered here which will be used by all routers
}

func (c *core) RegisterRouter(routerProxy Interface.RouterProxy) error {
	return c.routerController.RegisterRouter(routerProxy)
}

func (c *core) RegisterMiddleWare(name string, handler gin.HandlerFunc) {
	c.middlerWareController.RegisterMiddleWare(name, handler)
}
