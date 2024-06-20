package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExampleController struct {
	basicRoute Interface.Route
	restRoute  Interface.Route
}

func (ec *ExampleController) GetRoutes() []Interface.Route {
	return []Interface.Route{ec.basicRoute, ec.restRoute}
}

func (ec *ExampleController) setupRouters() {
	ec.basicRoute = route.NewRouter("/", http.MethodGet)
	{
		ec.basicRoute.SetHandler(ec.getExample)
	}

	ec.restRoute = route.NewREST("/api/example/:id")
	{
		ec.restRoute.SetHandler(
			ec.getRestHandler,
			ec.postRestHandler,
			ec.putRestHandler,
			ec.deleteRestHandler,
		)
	}
}

func (ec *ExampleController) getExample(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	resp.SetCode(response.OK).AddData("Get Example")
}

func (ec *ExampleController) getRestHandler(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	resp.SetCode(response.OK).AddData("Get Example")
}

func (ec *ExampleController) postRestHandler(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	resp.SetCode(response.OK).AddData("Get Example")
}

func (ec *ExampleController) putRestHandler(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	resp.SetCode(response.OK).AddData("Get Example")
}

func (ec *ExampleController) deleteRestHandler(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	resp.SetCode(response.OK).AddData("Get Example")
}

func NewExampleController() *ExampleController {
	c := &ExampleController{}
	c.setupRouters()
	return c
}
