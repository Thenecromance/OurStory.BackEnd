package Manager

import (
	"fmt"
	"github.com/Thenecromance/OurStories/Interface"
	Log "github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	type entry struct {
		route Interface.Route
		loaded bool
	}
*/
type routerManager struct {
	gin      *gin.Engine
	routeMap map[string]Interface.Route
}

func (r *routerManager) RegisterRouter(routerProxy ...Interface.Route) error {

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

func (r *routerManager) GetRouter(name string) (Interface.Route, error) {
	router, ok := r.routeMap[name]
	if !ok {
		return nil, fmt.Errorf("route %s not found", name)
	}
	return router, nil
}

func (r *routerManager) ApplyRouter() error {
	Log.Info(r.routeMap)
	for _, router := range r.routeMap {
		//Log.Infof("Registering route %s", route.GetPath())
		if router.IsRESTFUL() {
			r.gin.Handle(http.MethodGet, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[0])...)
			r.gin.Handle(http.MethodPost, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[1])...)
			r.gin.Handle(http.MethodPut, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[2])...)
			r.gin.Handle(http.MethodDelete, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[3])...)
		} else {
			r.gin.Handle(router.GetMethod(), router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[0])...)
		}

	}
	Log.Infof("All routers registered")
	return nil
}

func (r *routerManager) Close() error {
	//TODO implement me
	panic("implement me")
	return nil
}

func NewRouterManager(g *gin.Engine) Interface.RouterController {
	return &routerManager{
		gin:      g,
		routeMap: make(map[string]Interface.Route),
	}
}
