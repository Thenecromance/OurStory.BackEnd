package Weather

import (
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Interface.ControllerBase
	model Model
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{
		model: Model{},
	}
	c.LoadChildren(i...)
	return c
}

func (c *Controller) Name() string {
	return "weather"
}

func (c *Controller) SetRootGroup(group *gin.RouterGroup) {
	// parent group is  /api/
	c.ParentGroup = group
	//setup self group as /api/user
	c.Group = group.Group("/" + c.Name())
}

func (c *Controller) LoadChildren(children ...Interface.Controller) {
	c.Children = append(c.Children, children...)
	//setup children groups
	c.ChildrenSetGroup(c.Group)
}

// Use adds middleware to the Controller's group
func (c *Controller) Use(middleware ...gin.HandlerFunc) {
	c.Group.Use(middleware...)
}

func (c *Controller) BuildRoutes() {
	c.Group.GET("/", c.getWeather)
	c.ChildrenBuildRoutes()
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getWeather(ctx *gin.Context) {

	code := ctx.Query("city")
	if code == "" {
		ctx.JSON(400, gin.H{
			"error": "city code is required",
		})
		return
	}

	w := c.model.GetWeatherByCode(code)
	if w.Weather == "" {
		ctx.JSON(500, gin.H{
			"error": "fail to get weather info",
		})
		return
	}
	ctx.JSON(200, w)
	return
}
