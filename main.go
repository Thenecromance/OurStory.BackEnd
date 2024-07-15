package main

import (
	"github.com/Thenecromance/OurStories/SQL/MySQL"
	"github.com/Thenecromance/OurStories/application/controller"
	"github.com/Thenecromance/OurStories/application/repository"
	"github.com/Thenecromance/OurStories/application/services"
	"github.com/Thenecromance/OurStories/server"
)

// dependency injection
func newUserController(svr *server.Server) {
	repo := repository.NewUserRepository(MySQL.Default())
	svr.RegisterRepository(repo)
	s := services.NewUserService(repo)

	svr.RegisterController(controller.NewUserController(s))
}

func newTravelController(svr *server.Server) {
	repo := repository.NewTravelRepository(MySQL.Default())
	svr.RegisterRepository(repo)
	s := services.NewTravelService(repo)

	ctrl := controller.NewTravelController(s)
	svr.RegisterController(ctrl)
}

func newRelationShipController(svr *server.Server) {
	repo := repository.NewRelationShipRepository(MySQL.Default())
	svr.RegisterRepository(repo)
	s := services.NewRelationShipService(repo)

	ctrl := controller.NewRelationshipController(s)
	svr.RegisterController(ctrl)
}

func newAnniversaryController(svr *server.Server) {
	repo := repository.NewAnniversaryRepository(MySQL.Default())
	svr.RegisterRepository(repo)
	s := services.NewAnniversaryService(repo)

	ctrl := controller.NewAnniversaryController(s)
	svr.RegisterController(ctrl)
}

func newShopController(svr *server.Server) {
	sRepo := repository.NewShopRepository(MySQL.Default())
	cRepo := repository.NewCartRepository(MySQL.Default())
	svr.RegisterRepository(sRepo, cRepo)
	s := services.NewShopService(sRepo, cRepo)

	ctrl := controller.NewShopController(s)
	svr.RegisterController(ctrl)
}

func main() {

	svr := server.New()

	newUserController(svr)

	newTravelController(svr)

	newRelationShipController(svr)

	newAnniversaryController(svr)

	newShopController(svr)

	if err := svr.Run(); err != nil {
		panic(err)
	}
}
