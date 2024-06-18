package server

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	Log "github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type core struct {
	routerController     Interface.RouterController
	middleWareController Interface.MiddleWareController // all middlewares will be registered here which will be used by all routers

	cfg *config
	svr *http.Server
}

func (c *core) Run() {
	c.svr.ListenAndServe()
}

func (c *core) setupServer(handler http.Handler) {
	//c.cfg.load()

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
}

func (c *core) RegisterRouter(routerProxy Interface.RouterProxy) error {
	return c.routerController.RegisterRouter(routerProxy)
}

func (c *core) RegisterMiddleWare(name string, handler gin.HandlerFunc) {
	c.middleWareController.RegisterMiddleWare(name, handler)
}

func newCore() *core {
	c := &core{}
	c.cfg.load()
	return c
}
