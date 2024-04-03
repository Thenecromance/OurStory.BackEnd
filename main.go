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
	"github.com/Thenecromance/OurStories/middleWare/Auth/gJWT"
	"github.com/Thenecromance/OurStories/middleWare/Tracer"
	"github.com/Thenecromance/OurStories/server"
	"github.com/gin-contrib/cors"
	"time"
)

func loadMiddleWare(svr *server.Server) {
	svr.PreLoadMiddleWare("tracer", Tracer.MiddleWare())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // 允许所有来源
	//config.AllowOrigins = []string{"*"}
	svr.PreLoadMiddleWare("cors", cors.New(config))
	key := Config.GetStringWithDefault("JWT", "AuthKey", "putTheKeyHere")
	svr.PreLoadMiddleWare("jwt", gJWT.NewMiddleware(gJWT.WithExpireTime(time.Hour*24*15), gJWT.WithKey(key)))
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
