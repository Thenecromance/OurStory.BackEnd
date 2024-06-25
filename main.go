package main

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/application/controller"
	"github.com/Thenecromance/OurStories/application/repository"
	"github.com/Thenecromance/OurStories/application/services"
	"github.com/Thenecromance/OurStories/server"
	"github.com/Thenecromance/OurStories/thirdParty/SQL"
)

// dependency injection
func newUserController() Interface.IController {
	repo := repository.NewUserRepository(SQL.Get("user"))
	s := services.NewUserService(repo)
	return controller.NewUserController(s)

}

func newTravelController() Interface.IController {
	repo := repository.NewTravelRepository(SQL.Get("travel"))
	s := services.NewTravelService(repo)
	return controller.NewTravelController(s)
}

func newRelationShipController() Interface.IController {
	repo := repository.NewRelationShipRepository(SQL.Get("user"))
	s := services.NewRelationShipService(repo)

	return controller.NewRelationshipController(s)
}

func main() {

	svr := server.New()

	uc := newUserController()
	svr.RegisterRouter(uc.GetRoutes()...)

	tc := newTravelController()
	svr.RegisterRouter(tc.GetRoutes()...)

	rc := newRelationShipController()
	svr.RegisterRouter(rc.GetRoutes()...)

	svr.Run()
}
