package anniversary

import (
	response "github.com/Thenecromance/OurStories/backend/Response"
	"github.com/Thenecromance/OurStories/backend/anniversary/data"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
	"time"
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
	c.group.Router.GET("/:id", c.getAnniversaryList)
	c.group.Router.POST("/:id", c.addNewDate)
	c.group.Router.PUT("/:id", c.updateDate)
	c.group.Router.DELETE("/:id", c.deleteDate)

}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getAnniversaryList(ctx *gin.Context) {
	logger.Get().Debug("sss")
	resp := response.New(ctx)
	defer resp.Send()

	logger.Get().Debug("start to request Anniversary...")
	result := c.model.GetAnniversaryList()
	logger.Get().Debug("get list complete!")
	if result == nil {
		return
	}

	resp.AddData(result).SetCode(response.SUCCESS)
}

type InputCommon struct {
	Owner        string `form:"owner"   json:"owner,omitempty"    binding:"required"`
	Year         int    `form:"year"    json:"year,omitempty"     binding:"required"`
	Month        int    `form:"month"   json:"month,omitempty"    binding:"required"`
	Day          int    `form:"day"     json:"day,omitempty"      binding:"required"`
	Title        string `form:"title"   json:"title,omitempty"    binding:"required"`
	Info         string `form:"info"    json:"info,omitempty"     `
	ShareToOther bool   `form:"shared"  json:"shared" `
}
type InputUnix struct {
	Owner        string `form:"owner"       json:"owner,omitempty"      binding:"required"`
	TimeStamp    int64  `form:"time_stamp"  json:"time_stamp,omitempty" binding:"required"`
	Title        string `form:"title"       json:"title,omitempty"      binding:"required"`
	Info         string `form:"info"        json:"info,omitempty"      `
	ShareToOther bool   `form:"shared"      json:"shared" `
}

func (c *Controller) processCommon(common InputCommon) {
	ani := data.Anniversary{
		Owner:     common.Owner,
		TimeStamp: time.Date(common.Year, time.Month(common.Month), common.Day, 0, 0, 0, 0, time.Local).Unix(),
		Title:     common.Title,
		Info:      common.Info,
	}

	c.model.AddAnniversary(ani)
}
func (c *Controller) processUnix(unix InputUnix) {

	logger.Get().Debug(unix.TimeStamp)
	ani := data.Anniversary{
		Owner:     unix.Owner,
		Title:     unix.Title,
		Info:      unix.Info,
		TimeStamp: unix.TimeStamp,
	}
	c.model.AddAnniversary(ani)
}

func (c *Controller) addNewDate(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
	userId := ctx.Param("id")

	var input InputCommon
	var inputUnix InputUnix
	var errCommon, errUnix error
	if errCommon = ctx.ShouldBind(&input); errCommon == nil {
		logger.Get().Debug("common")
		if input.Owner != userId {
			logger.Get().Warnf("a genius is trying to update fake shit, input is %s , but the target user is %s", input.Owner, userId)
			resp.AddData("add failed").SetCode(response.FAIL)
			return
		}
		c.processCommon(input)
		resp.SetCode(response.SUCCESS).AddData("Done")

	} else if errUnix = ctx.ShouldBind(&inputUnix); errUnix == nil {
		logger.Get().Debug("Unix")
		if inputUnix.Owner != userId {
			logger.Get().Warnf("a genius is trying to update fake shit, input is %s , but the target user is %s", inputUnix.Owner, userId)
			resp.AddData("add failed").SetCode(response.FAIL)
			return
		}
		c.processUnix(inputUnix)
		resp.SetCode(response.SUCCESS).AddData("Done")
	} else {
		logger.Get().Error(errCommon)
		logger.Get().Error(errUnix)
		logger.Get().Debug("None")
		return
	}
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
