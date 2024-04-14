package main

import (
	"github.com/Thenecromance/OurStories/backend/Dashboard"
	"github.com/Thenecromance/OurStories/backend/Travel"
	"github.com/Thenecromance/OurStories/backend/UserV2"
	"github.com/Thenecromance/OurStories/backend/anniversary"
	"github.com/Thenecromance/OurStories/backend/api"
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

	svr.PreLoadMiddleWare("recovery", gin.Recovery())
	svr.PreLoadMiddleWare("log", gin.Logger())
	svr.PreLoadMiddleWare("blacklist", blacklist.NewMiddleWare())
}

func loadController(svr *server.Server) {

	userC := UserV2.NewController()
	svr.PreLoadMiddleWare("jwt", userC.Middleware())

	svr.Load(
		api.NewController(),
		Dashboard.NewController(),
		userC,
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
