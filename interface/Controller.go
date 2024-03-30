package Interface

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	//SetRootGroup(group *gin.RouterGroup)

	//LoadChildren(sub ...Controller)

	AddMiddleWare(middleware ...gin.HandlerFunc)

	BuildRoutes()

	Name() string

	GetNode() *RouteNode
}

type ControllerBase struct {
	*RouteNode
	Children []Controller
}

func (c *ControllerBase) ChildrenBuildRoutes() {
	for _, child := range c.Children {
		child.BuildRoutes()
	}
}

func (c *ControllerBase) GetNode() *RouteNode {
	return c.RouteNode
}
