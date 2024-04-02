package main

import (
	"github.com/Thenecromance/OurStories/backend/Dashboard"
	"github.com/Thenecromance/OurStories/backend/Location"
	"github.com/Thenecromance/OurStories/backend/Travel"
	"github.com/Thenecromance/OurStories/backend/User"
	"github.com/Thenecromance/OurStories/backend/Weather"
	"github.com/Thenecromance/OurStories/backend/api"
	"github.com/Thenecromance/OurStories/base/SQL"
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/middleWare/Auth/gJWT"
	"github.com/Thenecromance/OurStories/server"
	"github.com/gin-gonic/gin"
	"time"
)

func initGinMode() {
	mode := Config.GetBool("Server", "ReleaseMode")
	if mode {
		gin.SetMode(gin.ReleaseMode)
	} else if mode {
		Config.SetBool("Server", "ReleaseMode", false)
	}

}

func AddMiddleWare(middleWare gin.HandlerFunc, controller ...Interface.Controller) []Interface.Controller {
	for _, c := range controller {
		c.PreLoadMiddleWare(middleWare)
	}
	return controller
}

func loadServerComponent() *server.Server {

	svr := server.New(server.RunningWithCA("./setting/puppyyu.cn_bundle.crt", "./setting/puppyyu.cn.key"))
	logger.Get().Info("using Middle ware")
	key := Config.GetStringWithDefault("JWT", "AuthKey", "putTheKeyHere")

	//dashboard controller for control the dashboard text values' change

	list := AddMiddleWare(
		gJWT.NewMiddleware(gJWT.WithExpireTime(time.Hour*24*15), gJWT.WithKey(key)),
		api.NewController(),
		Location.NewController(),
		Weather.NewController(),
		Travel.NewController())

	svr.LoadComponent(
		Dashboard.NewController(),
		User.NewController(),
	)
	svr.LoadComponent(list...)

	return svr
}

func main() {

	initGinMode()
	svr := loadServerComponent()
	SQL.Initialize()
	svr.Run(":8080")

}
