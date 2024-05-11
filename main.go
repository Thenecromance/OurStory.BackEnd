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
	svr.RegisterMiddleWare("tracer", Tracer.MiddleWare())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	svr.RegisterMiddleWare("cors", cors.New(config)) // support all origins

	svr.RegisterMiddleWare("recovery", gin.Recovery())
	svr.RegisterMiddleWare("log", gin.Logger())
	svr.RegisterMiddleWare("blacklist", blacklist.NewMiddleWare())
}

func loadController(svr *server.Server) {

	user := UserV2.NewController()
	svr.RegisterMiddleWare("jwt", user.Middleware())

	svr.LoadController(
		api.NewController(),
		Dashboard.NewController(),
		user,
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
