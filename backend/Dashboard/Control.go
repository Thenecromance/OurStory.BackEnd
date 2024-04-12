package Dashboard

import (
	"github.com/Thenecromance/OurStories/backend/Dashboard/SideNavBar"
	response "github.com/Thenecromance/OurStories/backend/Response"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	group    *Interface.GroupNode
	snbModel *SideNavBar.Model
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController() Interface.Controller {
	c := &Controller{
		snbModel: SideNavBar.New(),
	}
	return c
}

func (c *Controller) RequestGroup(cb Interface.NodeCallback) {
	c.group = cb(c.Name(), "/")
}

func (c *Controller) Name() string {
	return "dashboard"
}

func (c *Controller) BuildRoutes() {
	//c.group.Router.GET("/topCards", c.getTopCard)
	//c.group.Router.GET("/sideNavBar", c.getSideNavBar)

	c.group.Router.GET("/navBar", c.getNavBar)
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getNavBar(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
	resp.AddData(c.snbModel.Items(0)).SetCode(response.SUCCESS)
	return
}
