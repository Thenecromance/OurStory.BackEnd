package demo

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/server/router"
	"github.com/gin-gonic/gin"
)

type ServiceDemo struct {
	ModelDemo
	ViewDemo
	ControllerDemo
}

type ControllerDemo struct {
	Router Interface.Router
}

type ModelDemo struct {
}

type ViewDemo struct {
}

func test() {
	demo := ControllerDemo{}
	demo.Router = router.NewREST()

	demo.Router.SetPath("/test")
	demo.Router.SetHandler([]gin.HandlerFunc{nil,
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		}, nil, nil})
}
