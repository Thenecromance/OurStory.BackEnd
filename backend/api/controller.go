package api

import Interface "github.com/Thenecromance/OurStories/interface"

type controller struct {
	group *Interface.GroupNode
}

func (ctrl *controller) Name() string {
	return "api"
}

func (ctrl *controller) BuildRoutes() {

}

func (ctrl *controller) RequestGroup(cb Interface.NodeCallback) {
	ctrl.group = cb(ctrl.Name(), "/")
}

func NewController(i ...Interface.Controller) Interface.Controller {
	ctrl := &controller{}
	return ctrl
}
