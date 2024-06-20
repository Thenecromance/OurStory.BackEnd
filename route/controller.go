package route

import (
	"fmt"
	"github.com/Thenecromance/OurStories/Interface"
	Log "github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type entry struct {
	router            Interface.IRoute
	hasBeenRegistered bool
}

type controller struct {
	gin   *gin.Engine
	proxy map[string]entry
}

func (c *controller) GetRouter(name string) (Interface.IRoute, error) {
	if router, ok := c.proxy[name]; ok {
		return router.router, nil
	}
	return nil, fmt.Errorf("route %s not found", name)
}

func (c *controller) RegisterRouter(routerProxy ...Interface.IRoute) error {
	/*c.proxy[routerProxy.GetPath()] = entry{
		route:            routerProxy,
		hasBeenRegistered: false,
	}*/

	for _, router := range routerProxy {
		_, ok := c.proxy[router.GetPath()]
		if ok {
			return fmt.Errorf("route %s already registered", router.GetPath())
		}
		c.proxy[router.GetPath()] = entry{
			router:            router,
			hasBeenRegistered: false,
		}
	}

	return nil
}

func (c *controller) Close() error {
	return nil
}

func (c *controller) ApplyRouter() error {
	Log.Debugf("Applying all routers to the server")
	for _, r := range c.proxy {
		if r.hasBeenRegistered {
			Log.Debugf("IRoute %s has already been registered", r.router.GetPath())
			continue
		}

		router := r.router
		c.gin.Handle(router.GetMethod(), router.GetPath(), router.GetHandler()...)
		r.hasBeenRegistered = true
	}

	return nil
}

func NewController() Interface.IRouterController {
	return &controller{}
}
