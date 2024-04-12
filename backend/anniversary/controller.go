package anniversary

import (
	response "github.com/Thenecromance/OurStories/backend/Response"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	group *Interface.GroupNode
	model *model
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController() Interface.Controller {
	c := &Controller{
		model: newModel(),
	}
	return c
}

func (c *Controller) RequestGroup(cb Interface.NodeCallback) {
	c.group = cb(c.Name(), "/")
}

func (c *Controller) Name() string {
	return "anniversary"
}

func (c *Controller) BuildRoutes() {
	c.group.Router.GET("/:id", c.getAniversaryList)
	c.group.Router.POST("/:id", c.addNewDate)
	c.group.Router.PUT("/:id", c.updateDate)
	c.group.Router.DELETE("/:id", c.deleteDate)

}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getAniversaryList(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	resp.AddData("get list ")

}

func (c *Controller) addNewDate(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
	resp.AddData("addNewDate")
}

func (c *Controller) updateDate(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
	resp.AddData("updateDate ")
}
func (c *Controller) deleteDate(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
	resp.AddData("deleteDate")

}
