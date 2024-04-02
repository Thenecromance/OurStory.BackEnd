package Bank

import (
	"github.com/Thenecromance/OurStories/backend"
	"github.com/Thenecromance/OurStories/backend/AMapToken"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	group *Interface.GroupNode

	Model
}

// ----------------------------Interface.Controller Implementation--------------------------------

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{}

	return c
}

func (c *Controller) Name() string {
	return "bank"
}

func (c *Controller) RequestGroup(cb Interface.NodeCallback) {
	c.group = cb(c.Name(), "/")
}

func (c *Controller) BuildRoutes() {
	AMapToken.Instance()
	c.group.Router.GET("/", c.getUserCredits) // get current user's credits
	//c.POST("/", c.getLocationByIP)
	c.group.Router.PUT("/", c.updateCredits) // when a user buys something, update credits
	c.group.Router.DELETE("/", c.refund)     // when a user refunds something, update credits
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
