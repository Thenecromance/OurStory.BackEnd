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
		router Interface.Router
		loaded bool
	}
*/
type routerManager struct {
	gin      *gin.Engine
	routeMap map[string]Interface.Router
}

func (r *routerManager) RegisterRouter(routerProxy Interface.Router) error {

	r.routeMap[routerProxy.GetPath()] = routerProxy
	return nil
}

func (r *routerManager) GetRouter(name string) (Interface.Router, error) {
	router, ok := r.routeMap[name]
	if !ok {
		return nil, fmt.Errorf("router %s not found", name)
	}
	return router, nil
}

func (r *routerManager) ApplyRouter() error {
	Log.Info(r.routeMap)
	for _, router := range r.routeMap {
		Log.Infof("Registering router %s", router.GetPath())
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
		routeMap: make(map[string]Interface.Router),
	}
}
