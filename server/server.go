package server

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin  *gin.Engine
	core *core
}

func (s *Server) RegisterRouter(routers ...Interface.IRoute) error {
	return s.core.routerController.RegisterRouter(routers...)
}

func (s *Server) RegisterMiddleWare(name string, handler gin.HandlerFunc) {
	return
	s.core.middleWareController.RegisterMiddleWare(name, handler)
}

func (s *Server) setTLS(tls Interface.ITLs) {
	s.core.Tls = tls
}

func (s *Server) setupResource() {
	s.gin.LoadHTMLGlob("dist/*.html")

	//s.gin.Handle("GET", "/", func(c *gin.Context) {
	//	/*	c.File("./dist/index.html")*/
	//})

	s.gin.Static("assets", "dist/assets")
	s.gin.Static("/favicon.svg", "dist/favicon.svg")

	s.gin.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})
	s.gin.NoMethod(func(c *gin.Context) {
		c.File("./dist/index.html")
	})
}

func (s *Server) initialize() {
	log.Infof("Initializing the server")
	s.setupResource()
	s.core.initializeCore(s.gin)

	log.Infof("Server initialized")
}

func (s *Server) Run() {

	defer log.Infof("Server is closing")
	if s.core == nil {
		s.core = newCore()
	}
	log.Info("Server is running")
	s.core.Run()

}

func (s *Server) Close() error {
	return nil
}

func New() *Server {
	svr := &Server{
		gin:  gin.Default(),
		core: newCore(),
	}
	svr.initialize()

	return svr
}
