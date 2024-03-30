package Bank

import (
	"github.com/Thenecromance/OurStories/backend"
	"github.com/Thenecromance/OurStories/backend/AMapToken"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Interface.ControllerBase

	Model
}

// ----------------------------Interface.Controller Implementation--------------------------------

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{}
	c.RouteNode = Interface.NewNode("credit", c.Name())
	c.LoadChildren(i...)
	return c
}

func (c *Controller) Name() string {
	return "bank"
}

func (c *Controller) LoadChildren(children ...Interface.Controller) {
	c.Children = append(c.Children, children...)
	//setup children groups
	//c.ChildrenSetGroup(c.Group)
}

// Use adds middleware to the Controller's group
func (c *Controller) AddMiddleWare(middleware ...gin.HandlerFunc) {
	c.AddMiddleWare(middleware...)
}

func (c *Controller) BuildRoutes() {
	AMapToken.Instance()
	c.GET("/", c.getUserCredits) // get current user's credits
	//c.POST("/", c.getLocationByIP)
	c.PUT("/", c.updateCredits) // when a user buys something, update credits
	c.DELETE("/", c.refund)     // when a user refunds something, update credits
	c.ChildrenBuildRoutes()
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getUserCredits(ctx *gin.Context) {
	backend.ResponseUnImplemented(ctx)
}

func (c *Controller) updateCredits(ctx *gin.Context) {
	backend.ResponseUnImplemented(ctx)
}

func (c *Controller) refund(ctx *gin.Context) {
	backend.ResponseUnImplemented(ctx)
}
