//package main
//
//import (
//	"github.com/Thenecromance/OurStories/backend/Dashboard"
//	"github.com/Thenecromance/OurStories/backend/Location"
//	"github.com/Thenecromance/OurStories/backend/Travel"
//	"github.com/Thenecromance/OurStories/backend/User"
//	"github.com/Thenecromance/OurStories/backend/Weather"
//	"github.com/Thenecromance/OurStories/backend/api"
//	"github.com/Thenecromance/OurStories/base/SQL"
//	Config "github.com/Thenecromance/OurStories/base/config"
//	"github.com/Thenecromance/OurStories/base/logger"
//	Interface "github.com/Thenecromance/OurStories/interface"
//	"github.com/Thenecromance/OurStories/middleWare/Auth/gJWT"
//	"github.com/Thenecromance/OurStories/server"
//	"github.com/gin-gonic/gin"
//	"time"
//)
//
//func initGinMode() {
//	mode := Config.GetBool("Server", "ReleaseMode")
//	if mode {
//		gin.SetMode(gin.ReleaseMode)
//	} else if mode {
//		Config.SetBool("Server", "ReleaseMode", false)
//	}
//
//}
//
//func AddMiddleWare(middleWare gin.HandlerFunc, controller ...Interface.Controller) []Interface.Controller {
//	for _, c := range controller {
//		c.PreLoadMiddleWare(middleWare)
//	}
//	return controller
//}
//
//func loadServerComponent() *server.Server {
//
//	svr := server.New(server.RunningWithCA("./setting/puppyyu.cn_bundle.crt", "./setting/puppyyu.cn.key"))
//	logger.Get().Info("using Middle ware")
//	key := Config.GetStringWithDefault("JWT", "AuthKey", "putTheKeyHere")
//
//	//dashboard controller for control the dashboard text values' change
//
//	list := AddMiddleWare(
//		gJWT.NewMiddleware(gJWT.WithExpireTime(time.Hour*24*15), gJWT.WithKey(key)),
//		api.NewController(),
//		Location.NewController(),
//		Weather.NewController(),
//		Travel.NewController())
//
//	svr.LoadComponent(
//		Dashboard.NewController(),
//		User.NewController(),
//	)
//	svr.LoadComponent(list...)
//
//	return svr
//}
//
//func main() {
//	initGinMode()
//	svr := loadServerComponent()
//	SQL.Initialize()
//	svr.Run(":8080")
//
//}

package main

import (
	"github.com/Thenecromance/OurStories/backend/Dashboard"
	"github.com/Thenecromance/OurStories/backend/Location"
	"github.com/Thenecromance/OurStories/backend/Travel"
	"github.com/Thenecromance/OurStories/backend/User"
	"github.com/Thenecromance/OurStories/backend/Weather"
	"github.com/Thenecromance/OurStories/backend/api"
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/middleWare/Auth/gJWT"
	"github.com/Thenecromance/OurStories/middleWare/Tracer"
	"github.com/Thenecromance/OurStories/server"
	"github.com/gin-contrib/cors"
	"time"
)

func loadMiddleWare(svr *server.Server) {
	svr.PreLoadMiddleWare("tracer", Tracer.MiddleWare())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                              // 允许所有来源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}             // 允许的请求方法
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"} // 允许的头部
	svr.PreLoadMiddleWare("cors", cors.New(config))

	svr.PreLoadMiddleWare("jwt", gJWT.NewMiddleware(gJWT.WithExpireTime(time.Hour*24*15), gJWT.WithKey("putTheKeyHere")))
	//svr.PreLoadMiddleWare(TLS.TlsHandler(8080))
}

func loadController(svr *server.Server) {
	svr.Load(
		api.NewController(),
		Dashboard.NewController(),
		User.NewController(),
		Location.NewController(),
		Weather.NewController(),
		Travel.NewController(),
	)
}

func main() {
	svr := server.New()
	SQL.Initialize()
	loadMiddleWare(svr)
	loadController(svr)
	svr.Run()

}
