package main

import (
	"github.com/Thenecromance/OurStories/backend/Dashboard"
	"github.com/Thenecromance/OurStories/backend/Location"
	"github.com/Thenecromance/OurStories/backend/User"
	"github.com/Thenecromance/OurStories/backend/Weather"
	"github.com/Thenecromance/OurStories/backend/api"
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/server"
	"github.com/gin-gonic/gin"
)

func loadServerComponent() *server.Server {
	gin.SetMode(gin.ReleaseMode)
	svr := server.New()

	api := api.NewController()
	api.SetRootGroup(svr.Group())
	api.LoadChildren(
		User.NewController(),
		Location.NewController(),
		Weather.NewController(),
	)

	//dashboard controller for control the dashboard text values' change
	dash := Dashboard.NewController()
	dash.SetRootGroup(svr.Group())

	svr.LoadComponent(api,
		dash,
	)
	return svr
}

func loadDashboardComponent() {

	////ArgonDash.LoadSSR(ArgonDash.NewTitle())
	//ArgonDash.LoadAPI(ArgonDashControl.GetTopCardController())
	//ArgonDash.LoadAPI(ArgonDashControl.GetSideNavBarController())
	////ArgonDashControl.LoadAPI(ArgonDashControl.GetSideNavBarController())
}

func main() {
	svr := loadServerComponent()
	loadDashboardComponent()
	SQL.Initialize()

	svr.Run(":8080")
}
