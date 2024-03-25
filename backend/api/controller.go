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

func (ctrl *controller) SetRootGroup(group *gin.RouterGroup) {
	// parent group is  /api/
	ctrl.ParentGroup = group
	//setup self group as /api/user
	ctrl.Group = group.Group("/" + ctrl.Name())

}

func (ctrl *controller) LoadChildren(sub ...Interface.Controller) {
	ctrl.Children = append(ctrl.Children, sub...)
	//setup children groups
	ctrl.ChildrenSetGroup(ctrl.Group)

}
func (ctrl *controller) Use(middleware ...gin.HandlerFunc) {
	if ctrl.Group == nil {
		return
	}
	ctrl.Group.Use(middleware...)
}
func (ctrl *controller) BuildRoutes() {
	ctrl.ChildrenBuildRoutes()
}

func NewController() Interface.Controller {
	ctrl := &controller{}
	return ctrl
}
