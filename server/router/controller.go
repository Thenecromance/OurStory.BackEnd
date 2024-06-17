package router

import "github.com/Thenecromance/OurStories/server/Interface"

type Controller struct {
}

func NewController() Interface.RouterController {
	return &Controller{}
}
