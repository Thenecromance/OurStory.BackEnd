package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/route"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteControl struct {
	route Interface.IRoute
}

func (r *RouteControl) Initialize() {
	r.SetRoutes()
}

func (r *RouteControl) Name() string {
	//TODO implement me
	panic("implement me")
}

func (r *RouteControl) SetRoutes() {
	/*	//TODO implement me
		panic("implement me")*/
	log.Info("start to set RouteControl routes")
	r.route = route.NewRouter("/dash/control", http.MethodGet)
	{
		r.route.SetHandler(r.getRouteHandler)
	}

	log.Info("RouteControl routes set")

}

func (r *RouteControl) GetRoutes() []Interface.IRoute {
	//TODO implement me
	panic("implement me")
}

func (r *RouteControl) getRouteHandler(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewRouteControl() Interface.IController {
	return &RouteControl{}
}
