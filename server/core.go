package server

import (
	"github.com/Thenecromance/OurStories/server/setting"
	"github.com/Thenecromance/OurStories/utility/File"
	Log "github.com/Thenecromance/OurStories/utility/log"
	"log"
	"net/http"
)

//type core struct {
//	routerController Interface.IRouterController
//
//	Tls Interface.ITLs
//	cfg *coreSetting
//	svr *http.Server
//}
//
//func (c *core) Run() {
//
//	c.routerController.ApplyRouter()
//
//	if c.Tls != nil {
//		Log.Infof("Server is running on %s with ITLs.\ncertificate file %s\nKey file path:%s", c.cfg.Addr, c.Tls.GetCertificate(), c.Tls.GetKey())
//		err := c.svr.ListenAndServeTLS(c.Tls.GetCertificate(), c.Tls.GetKey())
//		if err != nil {
//			Log.Errorf("Error while running the server with ITLs: %s", err.Error())
//			return
//		}
//	} else {
//		//Log.Infof("Server is running on %s without ITLs. use http to request", c.cfg.Addr)
//
//		if c.cfg.Addr == ":8080" {
//			Log.Description("server is running , visit by : http://localhost:8080")
//		}
//
//		err := c.svr.ListenAndServe()
//		/*if err != nil {
//			Log.Errorf("Error while running the server: %s", err.Error())
//			return
//		}*/
//		if errors.Is(err, http.ErrServerClosed) {
//			Log.Infof("Server closed")
//		} else {
//			Log.Errorf("Error while running the server: %s", err.Error())
//
//		}
//	}
//}
//
//func (c *core) setupServer(handler http.Handler) {
//	Log.Description("Setting up the server")
//	c.svr = &http.Server{
//		Addr:                         c.cfg.Addr,
//		Handler:                      handler,
//		DisableGeneralOptionsHandler: c.cfg.DisableGeneralOptionsHandler,
//		TLSConfig:                    nil, // do it later
//		ReadTimeout:                  time.Duration(c.cfg.ReadTimeout) * time.Second,
//		WriteTimeout:                 time.Duration(c.cfg.WriteTimeout) * time.Second,
//		IdleTimeout:                  time.Duration(c.cfg.IdleTimeout) * time.Second,
//		MaxHeaderBytes:               c.cfg.MaxHeaderBytes,
//		ErrorLog:                     log.New(Log.Instance.GetWriter(), "Core", 0),
//
//		/*TLSNextProto:                 nil,
//		ConnState:                    nil,
//		BaseContext:                  nil,
//		ConnContext:                  nil,*/
//	}
//	Log.Description("Server setup done")
//
//}
//
//func (c *core) close() {
//	c.svr.Close()
//}
//
//func (c *core) initializeCore(g *gin.Engine) {
//	Log.Description("Initializing the core")
//	c.setupServer(g)
//	{
//		Log.Description("Initializing the route manager")
//		c.routerController = Manager.NewRouterManager(g)
//		Log.Description("IRoute manager initialized")
//	}
//
//	//{
//	//	Logger.Description("Registering routers to gin")
//	//	c.routerController.ApplyRouter()
//	//	Logger.Description("Routers registered to gin")
//	//}
//
//	Log.Description("Core initialized")
//}
//
//func newCore() *core {
//	c := &core{
//		cfg: new(coreSetting),
//	}
//	c.cfg.load()
//	return c
//}

type core struct {
	svr     *http.Server
	setting *setting.Core
}

func (c *core) shutdown() {
	Log.Info("server is shutting down")
	if err := c.svr.Shutdown(nil); err != nil {
		Log.Errorf("server shutdown failed: %s", err.Error())
	}
}

func (c *core) initialize(handler http.Handler) {
	Log.Info("initialize core....")
	defer Log.Info("core initialized")

	{
		Log.Info("load core setting....")
		defer Log.Info("core setting loaded")
		c.setting.Load()
	}

	{
		Log.Info("initialize http.Server....")
		defer Log.Info("http.Server initialized")
		c.svr = &http.Server{
			ErrorLog:       log.New(Log.Instance.GetWriter(), "Core", 0),
			Handler:        handler,
			Addr:           c.setting.Addr,
			ReadTimeout:    c.setting.ReadTimeout,
			WriteTimeout:   c.setting.WriteTimeout,
			IdleTimeout:    c.setting.IdleTimeout,
			MaxHeaderBytes: c.setting.MaxHeaderBytes,
		}
	}
}

// getCertPath get the certificate file path if file path not exist , return empty
func (c *core) getCertPath() string {
	if c.setting.CertPath != "" && File.Exists(c.setting.CertPath) {
		return c.setting.CertPath
	}
	return ""
}

func (c *core) getKeyPath() string {
	if c.setting.KeyPath != "" && File.Exists(c.setting.KeyPath) {
		return c.setting.KeyPath
	}
	return ""
}

func (c *core) enableTLS() bool {
	return c.getCertPath() != "" && c.getKeyPath() != ""
}

func (c *core) run() error {
	if c.enableTLS() {
		Log.Infof("server start to run on %s with TLS , please visit by https://%s", c.setting.Addr, c.setting.Addr)
		return c.svr.ListenAndServeTLS(
			c.getCertPath(),
			c.getKeyPath(),
		)
	} else {
		Log.Infof("server start to run on %s without TLS , please visit by http://%s", c.setting.Addr, c.setting.Addr)
		return c.svr.ListenAndServe()
	}
}

func newCore() *core {
	return &core{
		setting: new(setting.Core),
	}
}
