package Manager

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"sync"
)

type Controller struct {
	mtx         sync.Mutex
	controllers []Interface.IController
}

func (c *Controller) RegisterController(controller ...Interface.IController) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for _, ctrl := range controller {
		log.Infof("Registering controller %s", ctrl.Name())
		//register the controller
		c.controllers = append(c.controllers, ctrl)
	}
}

func (c *Controller) Initialize() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	{
		log.Infof("initializing all controllers....")
		defer log.Info("All controllers initialized")
		for _, ctrl := range c.controllers {
			log.Infof("Initializing controller %s", ctrl.Name())
			ctrl.Initialize()
		}
	}

	{
		log.Info("Setting up routes for all controllers....")
		defer log.Info("All routes set up")
		for _, ctrl := range c.controllers {
			log.Infof("Setting up routes for controller %s", ctrl.Name())
			ctrl.SetupRoutes()
		}
	}

	log.Infof("initialized %d controllers", len(c.controllers))
}

func NewControllerMgr() *Controller {
	return &Controller{}
}
