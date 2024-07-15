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

	s := services.NewUserService(repo)

	svr.RegisterRepository(repo)
	svr.RegisterController(controller.NewUserController(s))
}

func newTravelController(svr *server.Server) {
	repo := repository.NewTravelRepository(MySQL.Default())

	s := services.NewTravelService(repo)

	ctrl := controller.NewTravelController(s)

	svr.RegisterRepository(repo)
	svr.RegisterController(ctrl)
}

func newRelationShipController(svr *server.Server) {
	repo := repository.NewRelationShipRepository(MySQL.Default())

	s := services.NewRelationShipService(repo)

	ctrl := controller.NewRelationshipController(s)

	svr.RegisterRepository(repo)
	svr.RegisterController(ctrl)
}

func newAnniversaryController(svr *server.Server) {
	repo := repository.NewAnniversaryRepository(MySQL.Default())

	s := services.NewAnniversaryService(repo)

	ctrl := controller.NewAnniversaryController(s)

	svr.RegisterRepository(repo)
	svr.RegisterController(ctrl)
}

func newShopController(svr *server.Server) {
	sRepo := repository.NewShopRepository(MySQL.Default())
	cRepo := repository.NewCartRepository(MySQL.Default())

	s := services.NewShopService(sRepo, cRepo)

	ctrl := controller.NewShopController(s)

	svr.RegisterRepository(sRepo, cRepo)
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
