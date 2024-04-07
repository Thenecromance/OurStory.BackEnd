package amapservice

import Interface "github.com/Thenecromance/OurStories/interface"

type controller struct {
	TokenInfo string `json:"token"`
}

func (ctrl *controller) BuildRoutes() {}

func (ctrl *controller) Name() string {
	return "api"
}

func (ctrl *controller) RequestGroup(cb Interface.NodeCallback) {}
func New() Interface.Controller {

	ctrl := &controller{}

	return ctrl
}
