package api

import (
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type controller struct {
	Interface.ControllerBase
}

func (ctrl *controller) Name() string {
	return "api"
}

func (ctrl *controller) LoadChildren(sub ...Interface.Controller) {
	ctrl.Children = append(ctrl.Children, sub...)
	//setup children groups
}

func (ctrl *controller) PreLoadMiddleWare(middleware ...gin.HandlerFunc) {
	/*if ctrl.Group == nil {
		return
	}
	ctrl.Group.PreLoadMiddleWare(middleware...)*/

	ctrl.CachedMiddleWare = append(ctrl.CachedMiddleWare, middleware...)

}
func (ctrl *controller) BuildRoutes() {
	ctrl.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Test": "test",
		})
	})
}

func (ctrl *controller) ApplyMiddleWare() {
	ctrl.Use(ctrl.CachedMiddleWare...)
}

func NewController(i ...Interface.Controller) Interface.Controller {
	ctrl := &controller{}
	ctrl.RouteNode = Interface.NewNode("/", ctrl.Name())
	ctrl.LoadChildren(i...)
	return ctrl
}
