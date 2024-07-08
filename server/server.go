package server

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/server/Manager"
	"github.com/Thenecromance/OurStories/server/setting"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

//
//import (
//	"github.com/Thenecromance/OurStories/Interface"
//	"github.com/Thenecromance/OurStories/server/resources"
//	"github.com/Thenecromance/OurStories/utility/log"
//	"github.com/gin-gonic/gin"
//)
//
//type Server struct {
//	gin       *gin.Engine
//	core      *core
//	resources *resources.Controller
//
//	command string
//}
//
//func (s *Server) RegisterRouter(routers ...Interface.IRoute) error {
//	return s.core.routerController.RegisterRouter(routers...)
//}
//
///*func (s *Server) RegisterMiddleWare(name string, handler gin.HandlerFunc) {
//	return
//	s.core.middleWareController.RegisterMiddleWare(name, handler)
//}*/
//
//func (s *Server) setTLS(tls Interface.ITLs) {
//	s.core.Tls = tls
//}
//
//func (s *Server) initialize() {
//	log.Infof("Initializing the server")
//
//	s.core.initializeCore(s.gin)
//
//	s.resources.ApplyTo(s.gin)
//
//	log.Infof("Server initialized")
//}
//
//func (s *Server) Run() {
//
//	defer log.Infof("Server is closing")
//	if s.core == nil {
//		s.core = newCore()
//	}
//
//	s.gin.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//
//	log.Info("Server is running")
//
//	s.core.Run()
//
//}
//
//func (s *Server) Close() error {
//	s.core.close()
//	return nil
//}
//
//func New() *Server {
//	svr := &Server{
//		gin:       gin.Default(),
//		core:      newCore(),
//		resources: resources.New(),
//	}
//	svr.initialize()
//
//	return svr
//}

type Server struct {
	// the base for the server
	core *core
	gin  *gin.Engine

	gs *setting.Gin // gin setting

	//=======================================
	controllerMgr *Manager.Controller
	//=======================================
}

// all the server stuff will initialize in this method
func (s *Server) initialize() {
	s.core.initialize(s.gin)

	s.setUpGinResource()

	s.controllerMgr.Initialize()
}

func (s *Server) setUpGinResource() {
	log.Info("Start to apply resources to gin engine")
	if s.gs.HtmlFiles != nil && len(s.gs.HtmlFiles) > 0 {
		//engine.LoadHTMLFiles(rc.cfg.HtmlFiles...)
		if len(s.gs.HtmlFiles) == 1 {
			s.gin.LoadHTMLGlob(s.gs.HtmlFiles[0])
		} else {
			s.gin.LoadHTMLFiles(s.gs.HtmlFiles...)
		}
	}

	if s.gs.NoMethod != "" {
		log.Infof("Setting NoMethod to %s", s.gs.NoMethod)
		s.gin.NoMethod(func(c *gin.Context) {
			c.File(s.gs.NoMethod)
		})
	}

	if s.gs.NoRoute != "" {
		log.Infof("Setting NoRoute to %s", s.gs.NoRoute)
		s.gin.NoRoute(func(c *gin.Context) {
			c.File(s.gs.NoRoute)
		})
	}

	if s.gs.ReMap != nil {
		for relativePath, root := range s.gs.ReMap {
			log.Infof("Mapping %s to %s", relativePath, root)
			s.gin.Static(relativePath, root)
		}
	}

	if s.gs.Redirects != nil && len(s.gs.Redirects) > 0 {
		for redirect, target := range s.gs.Redirects {
			log.Infof("Redirecting %s to %s", redirect, target)
			s.gin.GET(redirect, func(c *gin.Context) {
				c.HTML(200, "index.html", gin.H{})
			})
		}
	}

	log.Info("Resources applied to gin engine")
}

func (s *Server) RegisterController(controller ...Interface.IController) {
	s.controllerMgr.RegisterController(controller...)
}

// Run will Start the server
func (s *Server) Run() error {
	s.initialize()

	return s.core.run()
}

func (s *Server) Close() {
	defer s.core.shutdown() // core will shutdown at the end

}

func New() *Server {
	svr := &Server{
		core: newCore(),
		gin:  gin.Default(),
		gs:   setting.NewGinSetting(),

		controllerMgr: Manager.NewControllerMgr(),
	}

	return svr
}
