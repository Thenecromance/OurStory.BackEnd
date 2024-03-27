package main

import (
	"github.com/Thenecromance/OurStories/backend/Credit"
	"github.com/Thenecromance/OurStories/backend/Dashboard"
	"github.com/Thenecromance/OurStories/backend/Location"
	"github.com/Thenecromance/OurStories/backend/Travel"
	"github.com/Thenecromance/OurStories/backend/User"
	"github.com/Thenecromance/OurStories/backend/Weather"
	"github.com/Thenecromance/OurStories/backend/api"
	"github.com/Thenecromance/OurStories/base/SQL"
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/server"
	"github.com/gin-gonic/gin"
)

func loadServerComponent() *server.Server {

	svr := server.New()

	//dashboard controller for control the dashboard text values' change

	svr.LoadComponent(
		api.NewControllerWithGroup(svr.Group(),
			User.NewController(),
			Location.NewController(),
			Weather.NewController(),
			Travel.NewController(),
			Credit.NewController(),
		),
		Dashboard.NewControllerWithGroup(svr.Group()),
	)
	return svr
}

func loadDashboardComponent() {

	////ArgonDash.LoadSSR(ArgonDash.NewTitle())
	//ArgonDash.LoadAPI(ArgonDashControl.GetTopCardController())
	//ArgonDash.LoadAPI(ArgonDashControl.GetSideNavBarController())
	////ArgonDashControl.LoadAPI(ArgonDashControl.GetSideNavBarController())
}

func initGinMode() {
	mode := Config.GetBool("Server", "ReleaseMode")
	if mode {
		gin.SetMode(gin.ReleaseMode)
	} else if mode {
		Config.SetBool("Server", "ReleaseMode", false)
	}

}

func main() {

	initGinMode()
	svr := loadServerComponent()
	loadDashboardComponent()
	SQL.Initialize()
	svr.Run(":8080")

}
