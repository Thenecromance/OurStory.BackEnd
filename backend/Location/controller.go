package Location

import (
	"github.com/Thenecromance/OurStories/backend/AMapToken"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Interface.ControllerBase
	model Model
}

// ----------------------------Interface.Controller Implementation--------------------------------

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{
		model: Model{},
	}
	c.RouteNode = Interface.NewNode("api", c.Name())
	c.LoadChildren(i...)
	return c
}

func (c *Controller) Name() string {
	return "location"
}

//func (c *Controller) SetRootGroup(group *gin.RouterGroup) {
//	// parent group is  /api/
//	c.ParentGroup = group
//	//setup self group as /api/user
//	c.Group = group.Group("/" + c.Name())
//}

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
	c.GET("/ip", c.getLocationByIP)
	c.ChildrenBuildRoutes()
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
