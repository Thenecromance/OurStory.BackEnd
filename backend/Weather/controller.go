package Weather

import (
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	group *Interface.GroupNode
	model Model
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{
		model: Model{},
	}
	return c
}

func (c *Controller) Name() string {
	return "weather"
}

func (c *Controller) RequestGroup(cb Interface.NodeCallback) {
	c.group = cb(c.Name(), "api")
}

func (c *Controller) BuildRoutes() {
	c.group.Router.GET("/", c.getWeather)
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
