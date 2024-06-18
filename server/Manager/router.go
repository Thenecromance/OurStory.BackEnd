package Manager

import (
	"fmt"
	"github.com/Thenecromance/OurStories/Interface"
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
	for _, router := range r.routeMap {

		if router.IsRESTFUL() {
			r.gin.Handle(http.MethodGet, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[0])...)
			r.gin.Handle(http.MethodPost, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[1])...)
			r.gin.Handle(http.MethodPut, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[2])...)
			r.gin.Handle(http.MethodDelete, router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[3])...)
		} else {
			r.gin.Handle(router.GetMethod(), router.GetPath(), append(router.GetMiddleWare(), router.GetHandler()[0])...)
		}

	}
	return nil
}

func (r *routerManager) Close() error {
	//TODO implement me
	panic("implement me")
	return nil
}

func NewRouterManager() Interface.RouterController {
	return &routerManager{}
}
