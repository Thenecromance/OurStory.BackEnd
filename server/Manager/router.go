package Manager

import (
	"fmt"
	"github.com/Thenecromance/OurStories/Interface"
	Log "github.com/Thenecromance/OurStories/utility/log"
	"github.com/Thenecromance/OurStories/utility/mq"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	type entry struct {
		route Interface.IRoute
		loaded bool
	}
*/
type routeMgr struct {
	gin      *gin.Engine
	routeMap map[string]Interface.IRoute
}

func (r *routeMgr) RegisterRouter(routerProxy ...Interface.IRoute) error {

	//r.routeMap[routerProxy.GetPath()] = routerProxy
	for _, router := range routerProxy {
		_, ok := r.routeMap[router.GetPath()]
		if ok {
			return fmt.Errorf("route %s already registered", router.GetPath())
		}
		r.routeMap[router.GetPath()] = router
	}
	return nil
}

func (r *routeMgr) GetRouter(name string) (Interface.IRoute, error) {
	router, ok := r.routeMap[name]
	if !ok {
		return nil, fmt.Errorf("route %s not found", name)
	}
	return router, nil
}

func (r *routeMgr) ApplyRouter() error {
	Log.Info(r.routeMap)
	for _, router := range r.routeMap {
		//Logger.Infof("Registering route %s", route.GetPath())
		if router.IsRESTFUL() {
			if router.GetHandler()[0] != nil {
				r.gin.Handle(http.MethodGet, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[0])...)
			}
			if router.GetHandler()[1] != nil {
				r.gin.Handle(http.MethodPost, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[1])...)
			}
			if router.GetHandler()[2] != nil {
				r.gin.Handle(http.MethodPut, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[2])...)
			}
			if router.GetHandler()[3] != nil {
				r.gin.Handle(http.MethodDelete, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[3])...)
			}

		} else {
			r.gin.Handle(router.GetMethod(), router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[0])...)
		}
	}
	Log.Infof("All routers registered")
	return nil
}

func (r *routeMgr) Close() error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (r *routeMgr) Run() {
	mq.Publish(mq.DevRoute, r.routeMap)
}

func NewRouterManager(g *gin.Engine) Interface.IRouterController {
	return &routeMgr{
		gin:      g,
		routeMap: make(map[string]Interface.IRoute),
	}
}
