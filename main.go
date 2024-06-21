package main

import (
	"github.com/Thenecromance/OurStories/server"
	"github.com/Thenecromance/OurStories/services/controller"
	"github.com/Thenecromance/OurStories/services/repository"
	"github.com/Thenecromance/OurStories/services/services"
	"github.com/Thenecromance/OurStories/thirdParty/SQL"
)

// dependency injection
func newUserController() *controller.UserController {
	repo := repository.NewUserRepository(SQL.Get("user"))
	s := services.NewUserService(repo)
	return controller.NewUserController(s)

}

func newTravelContrller() *controller.TravelController {
	repo := repository.NewTravelRepository(SQL.Get("travel"))
	s := services.NewTravelService(repo)
	return controller.NewTravelController(s)
}

func main() {
	svr := server.New()

	uc := newUserController()
	svr.RegisterRouter(uc.GetRoutes()...)

	ec := controller.NewExampleController()
	svr.RegisterRouter(ec.GetRoutes()...)

	tc := newTravelContrller()
	svr.RegisterRouter(tc.GetRoutes()...)

	//svr.RegisterMiddleWare("auth", JWT.Middleware())

	svr.Run()
}
