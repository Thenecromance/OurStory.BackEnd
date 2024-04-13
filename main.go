package main

import (
	"github.com/Thenecromance/OurStories/backend/anniversary"
	"time"

	"github.com/Thenecromance/OurStories/backend/Dashboard"
	"github.com/Thenecromance/OurStories/backend/Travel"
	"github.com/Thenecromance/OurStories/backend/User"
	"github.com/Thenecromance/OurStories/backend/api"
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/middleWare/Auth/gJWT"
	blacklist "github.com/Thenecromance/OurStories/middleWare/BlackList"
	"github.com/Thenecromance/OurStories/middleWare/Tracer"
	"github.com/Thenecromance/OurStories/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func loadMiddleWare(svr *server.Server) {
	svr.PreLoadMiddleWare("tracer", Tracer.MiddleWare())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	svr.PreLoadMiddleWare("cors", cors.New(config)) // support all origins
	key := Config.GetStringWithDefault("JWT", "AuthKey", "putTheKeyHere")
	svr.PreLoadMiddleWare("jwt", gJWT.NewMiddleware(gJWT.WithExpireTime(time.Hour*24*15), gJWT.WithKey(key)))
	svr.PreLoadMiddleWare("recovery", gin.Recovery())
	svr.PreLoadMiddleWare("log", gin.Logger())
	svr.PreLoadMiddleWare("blacklist", blacklist.NewMiddleWare())
}

func loadController(svr *server.Server) {
	svr.Load(
		api.NewController(),
		Dashboard.NewController(),
		User.NewController(),
		Travel.NewController(),
		anniversary.NewController(),
	)
}

func main() {

	svr := server.New(
		// server.RunningWithCA("setting/cert.crt", "setting/key.key")
	)
	loadMiddleWare(svr)
	loadController(svr)
	svr.Run()
}

//1625500800
//834681600
