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

/*
	func (ctrl *controller) SetRootGroup(group *gin.RouterGroup) {
		// parent group is  /api/
		ctrl.ParentGroup = group
		//setup self group as /api/user
		ctrl.Group = group.Group("/" + ctrl.Name())

		ctrl.node.Group(ctrl.node.Name)

}
*/

func (ctrl *controller) LoadChildren(sub ...Interface.Controller) {
	ctrl.Children = append(ctrl.Children, sub...)
	//setup children groups
}

func (ctrl *controller) Use(middleware ...gin.HandlerFunc) {
	/*if ctrl.Group == nil {
		return
	}
	ctrl.Group.Use(middleware...)*/
	ctrl.Use(middleware...)

}
func (ctrl *controller) BuildRoutes() {

}

func NewController(i ...Interface.Controller) Interface.Controller {
	ctrl := &controller{}
	ctrl.RouteNode = Interface.NewNode("/", ctrl.Name())
	ctrl.LoadChildren(i...)
	return ctrl
}
func NewControllerWithGroup(group *gin.RouterGroup, i ...Interface.Controller) Interface.Controller {
	ctrl := &controller{}
	ctrl.RouteNode = Interface.NewNode("/", ctrl.Name())
	//ctrl.SetRootGroup(group)
	ctrl.LoadChildren(i...)
	return ctrl
}
