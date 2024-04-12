package anniversary

import (
	response "github.com/Thenecromance/OurStories/backend/Response"
	"github.com/Thenecromance/OurStories/base/logger"
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

	result := c.model.GetAnniversaryList()
	if result == nil {
		return
	}

	resp.AddData(result).SetCode(response.SUCCESS)
}

func (c *Controller) addNewDate(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
	userId := ctx.Param("id")

	type Input struct {
		Owner        string `form:"owner" json:"owner,omitempty"`
		Year         int    `form:"year" json:"year,omitempty"`
		Month        int    `form:"month" json:"month,omitempty"`
		Day          int    `form:"day" json:"day,omitempty"`
		Title        string `form:"title" json:"title,omitempty"`
		Info         string `form:"info" json:"info,omitempty"`
		ShareToOther bool   `form:"shared" json:"shared"`
	}
	var input Input
	if err := ctx.ShouldBind(&input); err != nil {
		logger.Get().Errorf("failed to bind object to receive %s", err)
		resp.AddData("add failed").SetCode(response.FAIL)
		return
	}

	if input.Owner != userId {
		logger.Get().Warnf("a genius is trying to update fake shit, input is %s , but the target user is %s", input.Owner, userId)

		resp.AddData("add failed").SetCode(response.FAIL)
		return
	}

	logger.Get().Info(input)
	resp.SetCode(response.SUCCESS).AddData(input)
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
