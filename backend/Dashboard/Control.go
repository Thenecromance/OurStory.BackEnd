package Dashboard

import (
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Interface.ControllerBase
	//model Model

	resource DynamicResource
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) Name() string {
	return "agronDash"
}

func (c *Controller) SetRootGroup(group *gin.RouterGroup) {
	// parent group is  /api/
	c.ParentGroup = group
	//setup self group as /api/user
	c.Group = group.Group("/" + c.Name())
}

func (c *Controller) LoadChildren(children ...Interface.Controller) {
	c.Children = append(c.Children, children...)
	//setup children groups
	c.ChildrenSetGroup(c.Group)
}

// Use adds middleware to the Controller's group
func (c *Controller) Use(middleware ...gin.HandlerFunc) {
	c.Group.Use(middleware...)
}

func (c *Controller) BuildRoutes() {
	c.Group.GET("/title", c.getTitle)
	c.ChildrenBuildRoutes()
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getTitle(ctx *gin.Context) {
	ctx.JSON(200, c.resource)
}

func NewController() Interface.Controller {
	c := &Controller{}
	c.resource.load()

	return c
}
