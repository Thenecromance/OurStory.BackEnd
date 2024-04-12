package server

import (
	"github.com/Thenecromance/OurStories/base/fileWatcher"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"log"
	"net/http"
	"time"
)

type core struct {
	option      CoreOption
	controllers []Interface.Controller
	svr         *http.Server

	cfg config
}

func (c *core) isTls() bool {
	return c.option.CertFile != "" && c.option.KeyFile != ""
}

func (c *core) initServer(handler http.Handler) {
	c.cfg.load()
	c.svr = &http.Server{
		Addr:    c.cfg.Addr, //
		Handler: handler,

		DisableGeneralOptionsHandler: c.cfg.DisableGeneralOptionsHandler,
		ReadTimeout:                  time.Duration(c.cfg.ReadTimeout) * time.Second,
		WriteTimeout:                 time.Duration(c.cfg.WriteTimeout) * time.Second,
		IdleTimeout:                  time.Duration(c.cfg.IdleTimeout) * time.Second,
		MaxHeaderBytes:               c.cfg.MaxHeaderBytes,

		ErrorLog: log.New(logger.GetWriter(), "[Server]", 0),
	}

	//build each routes
	for _, controller := range c.controllers {
		controller.BuildRoutes()
	}
}

// runServer starts the server
func (c *core) runServer() {
	defer fileWatcher.Close()
	defer c.svr.Close()

	if c.isTls() {
		logger.Get().Info("ListenAndServeTLS: ", c.cfg.Addr)
		err := c.svr.ListenAndServeTLS(c.option.CertFile, c.option.KeyFile)
		if err != nil {
			logger.Get().Error("ListenAndServeTLS error: ", err)
			return
		}

	} else {
		logger.Get().Info("ListenAndServe: ", c.cfg.Addr)
		err := c.svr.ListenAndServe()
		if err != nil {
			logger.Get().Error("ListenAndServe error: ", err)
			return
		}
	}

}

func (c *core) appendControllers(controllers ...Interface.Controller) {
	c.controllers = append(c.controllers, controllers...)
}

func newCore(opts ...Option) *core {
	c := &core{}

	for _, opt := range opts {
		opt(&c.option)
	}

	return c
}
