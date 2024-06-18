package server

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/server/Manager"
	Log "github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type core struct {
	routerController     Interface.RouterController
	middleWareController Interface.MiddleWareController // all middlewares will be registered here which will be used by all routers

	Tls Interface.TLS
	cfg *config
	svr *http.Server
}

func (c *core) Run() {

	c.routerController.ApplyRouter()

	if c.Tls != nil {
		Log.Infof("Server is running on %s with TLS.\ncertificate file %s\nKey file path:%s", c.cfg.Addr, c.Tls.GetCertificate(), c.Tls.GetKey())
		err := c.svr.ListenAndServeTLS(c.Tls.GetCertificate(), c.Tls.GetKey())
		if err != nil {
			Log.Errorf("Error while running the server with TLS: %s", err.Error())
			return
		}
	} else {
		Log.Infof("Server is running on %s without TLS. use http to request", c.cfg.Addr)
		err := c.svr.ListenAndServe()
		if err != nil {
			Log.Errorf("Error while running the server: %s", err.Error())
			return
		}
	}
}

func (c *core) setupServer(handler http.Handler) {
	Log.Info("Setting up the server")
	c.svr = &http.Server{
		Addr:                         c.cfg.Addr,
		Handler:                      handler,
		DisableGeneralOptionsHandler: c.cfg.DisableGeneralOptionsHandler,
		TLSConfig:                    nil, // do it later
		ReadTimeout:                  time.Duration(c.cfg.ReadTimeout) * time.Second,
		WriteTimeout:                 time.Duration(c.cfg.WriteTimeout) * time.Second,
		IdleTimeout:                  time.Duration(c.cfg.IdleTimeout) * time.Second,
		MaxHeaderBytes:               c.cfg.MaxHeaderBytes,
		ErrorLog:                     log.New(Log.Instance.GetWriter(), "Core", 0),

		/*TLSNextProto:                 nil,
		ConnState:                    nil,
		BaseContext:                  nil,
		ConnContext:                  nil,*/
	}
	Log.Info("Server setup done")
}

func (c *core) initializeCore(g *gin.Engine) {
	Log.Info("Initializing the core")
	c.setupServer(g)
	{
		Log.Info("Initializing the router manager")
		c.routerController = Manager.NewRouterManager(g)
		Log.Info("Router manager initialized")
	}

	//{
	//	Log.Info("Registering routers to gin")
	//	c.routerController.ApplyRouter()
	//	Log.Info("Routers registered to gin")
	//}

	Log.Info("Core initialized")
}

func newCore() *core {
	c := &core{
		cfg: new(config),
	}
	c.cfg.load()
	return c
}
