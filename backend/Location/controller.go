package Location

import (
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	group *Interface.GroupNode
	model Model
}

// ----------------------------Interface.Controller Implementation--------------------------------

func NewController() Interface.Controller {
	c := &Controller{
		model: Model{},
	}
	return c
}

func (c *Controller) Name() string {
	return "location"
}

func (c *Controller) RequestGroup(cb Interface.NodeCallback) {
	c.group = cb(c.Name(), "api")
}
func (c *Controller) BuildRoutes() {
	// AMapToken.Instance()

	c.group.Router.GET("/ip", c.getLocationByIP)
}

//----------------------------Interface.Controller Implementation--------------------------------

// ------------------------------------------------------------
func (c *Controller) getLocationByIP(ctx *gin.Context) {
	ip := ctx.ClientIP() // get Client's IP first to get location

	loc := c.model.GetLocation(ip)
	if loc.Status != "1" {
		//c.JSON(400, gin.H{"status": "failed", "message": "fail to get location"})
		ctx.JSON(200, Data{
			Status:    "1",
			Info:      "OK",
			Infocode:  "10000",
			Province:  "北京市",
			City:      "北京市",
			Adcode:    "110000",
			Rectangle: "116.0119343,39.66127144;116.7829835,40.2164962",
		})
		return
	}

	ctx.JSON(200, loc)
}

// ------------------------------------------------------------
