package Dashboard

import (
	"github.com/Thenecromance/OurStories/backend"
	"github.com/Thenecromance/OurStories/backend/Location"
	"github.com/Thenecromance/OurStories/backend/Weather"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Interface.ControllerBase
	//model Model

	resource DynamicResource
	weather  Weather.Model
	location Location.Model
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{}
	c.resource.load()
	c.RouteNode = Interface.NewNode("/", c.Name())
	c.LoadChildren(i...)
	return c
}

func (c *Controller) Name() string {
	return "agronDash"
}

/*func (c *Controller) SetRootGroup(group *gin.RouterGroup) {
	// parent group is  /api/
	c.ParentGroup = group
	//setup self group as /api/user
	c.Group = group.Group("/" + c.Name())
}*/

func (c *Controller) LoadChildren(children ...Interface.Controller) {
	c.Children = append(c.Children, children...)
	//setup children groups
	//c.ChildrenSetGroup(c.Group)
}

// Use adds middleware to the Controller's group
func (c *Controller) Use(middleware ...gin.HandlerFunc) {
	c.Use(middleware...)
}

func (c *Controller) BuildRoutes() {
	c.GET("/title", c.getTitle)
	c.GET("/topCards", c.getTopCard)
	c.GET("/sideNavBar", c.getSideNavBar)
	c.ChildrenBuildRoutes()
}

func (c *Controller) GetTitle() string {
	return c.resource.Title
}

//----------------------------Interface.Controller Implementation--------------------------------

func (c *Controller) getTitle(ctx *gin.Context) {
	ctx.JSON(200, c.resource)
}

func (c *Controller) getTopCard(ctx *gin.Context) {
	var cardsInfo []topCardItem
	cardsInfo = append(cardsInfo, c.getLocationWeather(ctx.ClientIP()))
	backend.Resp(ctx, cardsInfo)
}

// getLocationWeather get the location and weather info
func (c *Controller) getLocationWeather(ip string) topCardItem {
	logger.Get().Info(ip)
	loc := c.location.GetLocation("111.197.34.158")
	todayWeather := c.weather.GetWeatherByCode(loc.Adcode)

	return topCardItem{
		Title:       loc.City,
		Value:       todayWeather.Weather,
		Description: todayWeather.Temperature + "°C",
		ShowIcon: Icon{
			Component:  "",
			Background: "",
			Shape:      "",
		},
	}

}

func (c *Controller) getSideNavBar(ctx *gin.Context) {

}
