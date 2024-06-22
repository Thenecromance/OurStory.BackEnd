package main

import (
	"fmt"
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/server"
	"github.com/Thenecromance/OurStories/services/controller"
	"github.com/Thenecromance/OurStories/services/repository"
	"github.com/Thenecromance/OurStories/services/services"
	"github.com/Thenecromance/OurStories/services/services/Encryptor/Scrypt"
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

func main() {
	en := Scrypt.New()
	a, b := en.Hash("123456")
	println(a, b)
	fmt.Println(en.Verify("123456", a, b))

	return

	svr := server.New()

	uc := newUserController()
	svr.RegisterRouter(uc.GetRoutes()...)

	/*	ec := controller.NewExampleController()
		svr.RegisterRouter(ec.GetRoutes()...)*/

	tc := newTravelController()
	svr.RegisterRouter(tc.GetRoutes()...)

	//svr.RegisterMiddleWare("auth", JWT.Middleware())

	svr.Run()
}
