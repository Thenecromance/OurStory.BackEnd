package main

import (
	"fmt"
	"github.com/Thenecromance/OurStories/backend/Dashboard"
	"github.com/Thenecromance/OurStories/backend/Location"
	"github.com/Thenecromance/OurStories/backend/Travel"
	"github.com/Thenecromance/OurStories/backend/User"
	"github.com/Thenecromance/OurStories/backend/Weather"
	"github.com/Thenecromance/OurStories/backend/api"
	"github.com/Thenecromance/OurStories/base/SQL"
	Config "github.com/Thenecromance/OurStories/base/config"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/server"
	"github.com/gin-gonic/gin"
)

func loadServerComponent() *server.Server {

	svr := server.New()

	//dashboard controller for control the dashboard text values' change

	svr.LoadComponent(
		api.NewController(

			//Credit.NewController(),
		),
		Dashboard.NewController(),

		User.NewController(),
		Location.NewController(),
		Weather.NewController(),
		Travel.NewController(),
	)
	return svr
}

func initGinMode() {
	mode := Config.GetBool("Server", "ReleaseMode")
	if mode {
		gin.SetMode(gin.ReleaseMode)
	} else if mode {
		Config.SetBool("Server", "ReleaseMode", false)
	}

}

func nodeTest() {
	temp := Interface.NewNode("credit", "store")
	ctrl := Interface.NewRootNode()
	api := Interface.NewNode("/", "api")
	ctrl.Load(api,
		Interface.NewNode("api", "user"),
		Interface.NewNode("api", "location"),
		Interface.NewNode("api", "weather"),
		Interface.NewNode("api", "travel"),
		Interface.NewNode("api", "credit"),
		Interface.NewNode("/", "dashboard"),
		Interface.NewNode("dashboard", "topcard"),
		Interface.NewNode("dashboard", "sidenavbar"),
		Interface.NewNode("dashboard", "main"),
		Interface.NewNode("credit", "bank"),
		Interface.NewNode("a", "D"), // this node will be ignored
		temp,
	)

	ctrl.MakeAsTree()

	fmt.Println(ctrl.String())
	fmt.Println(temp.Path())
}

func main() {

	initGinMode()
	svr := loadServerComponent()
	SQL.Initialize()
	svr.Run(":8080")

}
