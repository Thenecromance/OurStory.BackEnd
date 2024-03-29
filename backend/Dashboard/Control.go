package Dashboard

import (
	"github.com/Thenecromance/OurStories/backend"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Interface.ControllerBase
	//model Model

	resource DynamicResource
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{}
	c.resource.load()
	c.RouteNode = Interface.NewNode("/", c.Name())
	c.LoadChildren(i...)
	return c
}

func (c *Controller) Name() string {
	return "agronDash"
}

/*func (c *Controller) SetRootGroup(group *gin.RouterGroup) {
	// parent group is  /api/
	c.ParentGroup = group
	//setup self group as /api/user
	c.Group = group.Group("/" + c.Name())
}*/

func (c *Controller) LoadChildren(children ...Interface.Controller) {
	c.Children = append(c.Children, children...)
	//setup children groups
	//c.ChildrenSetGroup(c.Group)
}

// Use adds middleware to the Controller's group
func (c *Controller) Use(middleware ...gin.HandlerFunc) {
	c.Use(middleware...)
}

func (c *Controller) BuildRoutes() {
	c.GET("/title", c.getTitle)
	c.GET("/topCard", c.getTopCard)
	c.ChildrenBuildRoutes()
}

func (c *Controller) GetTitle() string {
	return c.resource.Title
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getTitle(ctx *gin.Context) {
	ctx.JSON(200, c.resource)
}

func (c *Controller) getTopCard(ctx *gin.Context) {
	var cardsInfo []topCardItem
	backend.Resp(ctx, cardsInfo)
}
