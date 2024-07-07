package server

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/server/resources"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin       *gin.Engine
	core      *core
	resources *resources.Controller
}

func (s *Server) RegisterRouter(routers ...Interface.IRoute) error {
	return s.core.routerController.RegisterRouter(routers...)
}

/*func (s *Server) RegisterMiddleWare(name string, handler gin.HandlerFunc) {
	return
	s.core.middleWareController.RegisterMiddleWare(name, handler)
}*/

func (s *Server) setTLS(tls Interface.ITLs) {
	s.core.Tls = tls
}

func (s *Server) initialize() {
	log.Infof("Initializing the server")

	s.core.initializeCore(s.gin)

	s.resources.ApplyTo(s.gin)

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
		gin:       gin.Default(),
		core:      newCore(),
		resources: resources.New(),
	}
	svr.initialize()

	return svr
}
