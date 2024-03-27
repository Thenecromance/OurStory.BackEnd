package Dashboard

import (
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/server"
	"github.com/gin-gonic/gin"
	"html/template"
)

type Controller struct {
	Interface.ControllerBase
	//model Model

	resource DynamicResource
}

//----------------------------Interface.Controller Implementation--------------------------------

/*func NewControllerWithGroup(i ...Interface.Controller) Interface.Controller {
	c := &Controller{}
	c.resource.load()

	server.AppendFuncMap(template.FuncMap{
		"GetTitle": c.GetTitle,
	})

	//c.SetRootGroup(group)
	c.LoadChildren(i...)
	return c
}*/

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{}
	c.resource.load()
	c.RouteNode = Interface.NewNode("/", c.Name())
	server.AppendFuncMap(template.FuncMap{
		"GetTitle": c.GetTitle,
	})
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
	c.ChildrenBuildRoutes()
}

func (c *Controller) GetTitle() string {
	return c.resource.Title
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getTitle(ctx *gin.Context) {
	ctx.JSON(200, c.resource)
}
