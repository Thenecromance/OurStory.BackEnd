package server

import (
	"github.com/Thenecromance/OurStories/Interface"
	Config "github.com/Thenecromance/OurStories/utility/config"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin  *gin.Engine
	core *core
}

func (s *Server) RegisterRouter(routers ...Interface.Route) error {
	return s.core.routerController.RegisterRouter(routers...)
}

func (s *Server) RegisterMiddleWare(name string, handler gin.HandlerFunc) {
	s.core.middleWareController.RegisterMiddleWare(name, handler)
}

func (s *Server) setTLS(tls Interface.TLS) {
	s.core.Tls = tls
}

func (s *Server) initialize() {
	log.Infof("Initializing the server")
	s.core.initializeCore(s.gin)
	log.Infof("Server initialized")
}

func (s *Server) Run() {
	Config.Flush()
	defer Config.CloseIni()
	if s.core == nil {
		s.core = newCore()
	}
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
