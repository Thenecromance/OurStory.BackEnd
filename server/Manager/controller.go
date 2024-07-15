package Manager

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"sync"
)

type ControllerMgr struct {
	mtx         sync.Mutex
	controllers []Interface.IController
}

func (c *ControllerMgr) RegisterController(controller ...Interface.IController) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for _, ctrl := range controller {
		log.Infof("[controller] Registering %s", ctrl.Name())
		//register the controller
		c.controllers = append(c.controllers, ctrl)
	}
}

func (c *ControllerMgr) Initialize() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	{
		log.Infof("initializing all controllers....")
		defer log.Info("All controllers initialized")
		for _, ctrl := range c.controllers {
			log.Infof("[controller] Initializing %s", ctrl.Name())
			ctrl.Initialize()
		}
	}

	{
		log.Info("Setting up routes for all controllers....")
		defer log.Info("All routes set up")
		for _, ctrl := range c.controllers {
			log.Infof("[controller] Setting up routes for %s", ctrl.Name())
			ctrl.SetupRoutes()
		}
	}

	log.Infof("initialized %d controllers", len(c.controllers))
}

func (c *ControllerMgr) GetAllControllers() []Interface.IController {
	return c.controllers
}

func NewControllerMgr() *ControllerMgr {
	return &ControllerMgr{}
}
