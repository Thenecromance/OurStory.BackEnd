package main

import (
	"github.com/Thenecromance/OurStories/server"
	"github.com/Thenecromance/OurStories/services/controller"
	"github.com/Thenecromance/OurStories/services/repository"
	"github.com/Thenecromance/OurStories/services/services"
	"github.com/Thenecromance/OurStories/thirdParty/SQL"
)

func newUserController() *controller.UserController {
	repo := repository.NewUserRepository(SQL.Get("user"))
	s := services.NewUserService(repo)
	return controller.NewUserController(s)

}

func main() {
	svr := server.New()

	/*	r := router.NewDefaultRouter()
		r.SetMethod("PUT")
		r.SetPath("/test")
		r.SetHandler(func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})*/
	uc := newUserController()
	svr.RegisterRouter(uc.GetRoutes()...)

	/*	r := router.NewREST("/rest")
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
		r2 := router.NewDefaultRouter()
		r2.SetMethod("GET")
		r2.SetPath("/test")
		r2.SetHandler(func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		svr.RegisterRouter(r)
		svr.RegisterRouter(r2)*/

	svr.Run()
}
