package anniversary

import (
	"github.com/Thenecromance/OurStories/application/model/anniversary"
	response "github.com/Thenecromance/OurStories/backend/Response"
	"github.com/Thenecromance/OurStories/backend/anniversary/data"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
	"time"
)

type Controller struct {
	group *Interface.GroupNode
	model *anniversary.Model
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController() Interface.Controller {
	c := &Controller{
		model: anniversary.NewModel(),
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
	c.group.Router.PUT("/:id", c.updateAnniversary)
	c.group.Router.DELETE("/:id", c.deleteDate)

}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getAnniversaryList(ctx *gin.Context) {
	log.Debug("sss")
	resp := response.New(ctx)
	defer resp.Send()

	log.Debug("start to request Anniversary...")
	result := c.model.GetAnniversaryList()
	log.Debug("get list complete!")
	if result == nil {
		return
	}

	resp.AddData(result).SetCode(response.SUCCESS)
}

// InputCommon for support client update data with common date format
type InputCommon struct {
	Owner        string `form:"owner"   json:"owner,omitempty"    binding:"required"`
	Year         int    `form:"year"    json:"year,omitempty"     binding:"required"`
	Month        int    `form:"month"   json:"month,omitempty"    binding:"required"`
	Day          int    `form:"day"     json:"day,omitempty"      binding:"required"`
	Title        string `form:"title"   json:"title,omitempty"    binding:"required"`
	Info         string `form:"info"    json:"info,omitempty"     `
	ShareToOther bool   `form:"shared"  json:"shared" `
}

// InputUnix support client update data with unix date format
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

	err := c.model.AddAnniversary(ani)
	if err != nil {
		log.Debug(err)
		return
	}
}
func (c *Controller) processUnix(unix InputUnix) {
	log.Debug(unix.TimeStamp)
	ani := data.Anniversary{
		Owner:     unix.Owner,
		Title:     unix.Title,
		Info:      unix.Info,
		TimeStamp: unix.TimeStamp,
	}
	err := c.model.AddAnniversary(ani)
	if err != nil {
		log.Debug(err)
		return
	}
}

func (c *Controller) addNewDate(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()
	userId := ctx.Param("id")

	var input InputCommon
	var inputUnix InputUnix
	var errCommon, errUnix error
	if errCommon = ctx.ShouldBind(&input); errCommon == nil {
		log.Debug("common")
		if input.Owner != userId {
			log.Warnf("a genius is trying to update fake shit, input is %s , but the target user is %s", input.Owner, userId)
			resp.AddData("add failed").SetCode(response.FAIL)
			return
		}
		c.processCommon(input)
		resp.SetCode(response.SUCCESS).AddData("Done")

	} else if errUnix = ctx.ShouldBind(&inputUnix); errUnix == nil {
		log.Debug("Unix")
		if inputUnix.Owner != userId {
			log.Warnf("a genius is trying to update fake shit, input is %s , but the target user is %s", inputUnix.Owner, userId)
			resp.AddData("add failed").SetCode(response.FAIL)
			return
		}
		c.processUnix(inputUnix)
		resp.SetCode(response.SUCCESS).AddData("Done")
	} else {
		log.Error(errCommon)
		log.Error(errUnix)
		log.Debug("None")
		return
	}
}

func (c *Controller) updateAnniversary(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	//c.models.GetAnniversaryById()

	resp.AddData("updateAnniversary ")
}
func (c *Controller) deleteDate(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	resp.AddData("deleteDate")
}
