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
	group *Interface.GroupNode

	resource DynamicResource
	weather  Weather.Model
	location Location.Model
	sNav     SideNavBarModel
}

//----------------------------Interface.Controller Implementation--------------------------------

func NewController() Interface.Controller {
	c := &Controller{}
	c.sNav.Load()
	c.resource.load()

	return c
}

func (c *Controller) RequestGroup(cb Interface.NodeCallback) {
	c.group = cb(c.Name(), "/")
}

func (c *Controller) Name() string {
	return "agronDash"
}

func (c *Controller) BuildRoutes() {
	c.group.Router.GET("/title", c.getTitle)
	c.group.Router.GET("/topCards", c.getTopCard)
	c.group.Router.GET("/sideNavBar", c.getSideNavBar)
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
