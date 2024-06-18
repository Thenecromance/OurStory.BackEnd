package main

import (
	"github.com/Thenecromance/OurStories/router"
	"github.com/Thenecromance/OurStories/server"
	"github.com/gin-gonic/gin"
)

func main() {
	svr := server.New()

	/*	r := router.NewRouter()
		r.SetMethod("PUT")
		r.SetPath("/test")
		r.SetHandler(func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})*/

	r := router.NewREST()
	r.SetPath("/rest")
	r.SetHandler(func(c *gin.Context) {
		c.JSON(200, gin.H{"Type": "GET"})
	}, func(c *gin.Context) {
		c.JSON(200, gin.H{"Type": "POST"})
	},
		func(c *gin.Context) {
			c.JSON(200, gin.H{"Type": "PUT"})
		},
		func(c *gin.Context) {
			c.JSON(200, gin.H{"Type": "DELETE"})
		})
	r2 := router.NewRouter()
	r2.SetMethod("GET")
	r2.SetPath("/test")
	r2.SetHandler(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	svr.RegisterRouter(r)
	svr.RegisterRouter(r2)

	svr.Run()
}
