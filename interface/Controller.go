package Interface

import "github.com/gin-gonic/gin"

type Controller interface {
	SetRootGroup(group *gin.RouterGroup)

	LoadChildren(sub ...Controller)

	Use(middleware ...gin.HandlerFunc)

	BuildRoutes()

	Name() string
}

type ControllerBase struct {
	ParentGroup *gin.RouterGroup
	Group       *gin.RouterGroup
	Children    []Controller
}

func (c *ControllerBase) ChildrenSetGroup(group *gin.RouterGroup) {
	for _, child := range c.Children {
		child.SetRootGroup(group)
	}
}
func (c *ControllerBase) ChildrenBuildRoutes() {
	for _, child := range c.Children {
		child.BuildRoutes()
	}
}
