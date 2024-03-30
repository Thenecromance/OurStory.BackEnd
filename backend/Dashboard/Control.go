package Dashboard

import (
	"github.com/Thenecromance/OurStories/backend"
	"github.com/Thenecromance/OurStories/backend/Location"
	"github.com/Thenecromance/OurStories/backend/Weather"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
	"strings"
)

type Controller struct {
	Interface.ControllerBase
	//model Model

	resource DynamicResource
	weather  Weather.Model
	location Location.Model
	sNav     SideNavBarModel
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController(i ...Interface.Controller) Interface.Controller {
	c := &Controller{}
	c.sNav.Load()
	c.resource.load()
	c.RouteNode = Interface.NewNode("/", c.Name())
	c.LoadChildren(i...)
	return c
}

func (c *Controller) Name() string {
	return "agronDash"
}

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
func (c *Controller) getLocationWeather(ip string) (item topCardItem) {
	logger.Get().Info(ip)
	if strings.HasSuffix(ip, "::1") {
		todayWeather := c.weather.GetWeatherByCode("110105")
		item.Title = "北京"
		item.Value = todayWeather.Weather
		item.Description = todayWeather.Temperature + "°C"
		item.ShowIcon.Component = "qi qi-100"
		item.ShowIcon.Background = "bg-primary"
		item.ShowIcon.Shape = "circle"
		return
	}

	loc := c.location.GetLocation(ip)
	todayWeather := c.weather.GetWeatherByCode(loc.Adcode)
	item.Title = loc.City
	item.Value = todayWeather.Weather
	item.Description = todayWeather.Temperature + "°C"
	item.ShowIcon.Component = "qi qi-100"
	item.ShowIcon.Background = "bg-primary"
	item.ShowIcon.Shape = "circle"

	return
}

func (c *Controller) getSideNavBar(ctx *gin.Context) {
	backend.Resp(ctx, c.sNav.Navs)
}
