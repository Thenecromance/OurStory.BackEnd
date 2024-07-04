package main

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/SQL/MySQL"
	"github.com/Thenecromance/OurStories/application/controller"
	"github.com/Thenecromance/OurStories/application/repository"
	"github.com/Thenecromance/OurStories/application/services"
	"github.com/Thenecromance/OurStories/server"
)

// dependency injection
func newUserController() Interface.IController {
	repo := repository.NewUserRepository(MySQL.Get("user"))
	s := services.NewUserService(repo)

	return controller.NewUserController(s)
}

func newTravelController() Interface.IController {
	repo := repository.NewTravelRepository(MySQL.Get("travel"))
	s := services.NewTravelService(repo)

	return controller.NewTravelController(s)
}

func newRelationShipController() Interface.IController {
	repo := repository.NewRelationShipRepository(MySQL.Get("user"))
	s := services.NewRelationShipService(repo)

	return controller.NewRelationshipController(s)
}

func newAnniversaryController() Interface.IController {
	repo := repository.NewAnniversaryRepository(MySQL.Get("user"))
	s := services.NewAnniversaryService(repo)

	return controller.NewAnniversaryController(s)
}

func main() {

	svr := server.New()

	uc := newUserController()
	svr.RegisterRouter(uc.GetRoutes()...)

	tc := newTravelController()
	svr.RegisterRouter(tc.GetRoutes()...)

	rc := newRelationShipController()
	svr.RegisterRouter(rc.GetRoutes()...)

	ac := newAnniversaryController()
	svr.RegisterRouter(ac.GetRoutes()...)

	svr.Run()
}
