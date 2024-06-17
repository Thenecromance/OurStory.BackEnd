package server

import (
	"github.com/Thenecromance/OurStories/server/Interface"
)

type core struct {
	routerController      Interface.RouterController
	middlerWareController Interface.GlobalMiddleWareController // all middlewares will be registered here which will be used by all routers
}

func NewCore(routerController Interface.RouterController, middlerWareController Interface.GlobalMiddleWareController) *core {
	return &core{
		routerController:      routerController,
		middlerWareController: middlerWareController,
	}
}
