package server

import (
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type Server struct {
	core *core
	gin  *ginControl
}

func (s *Server) initialize() {
	if s.core == nil || s.gin == nil {
		logger.Get().Error("before initialize server ,please make sure the core and gin is not nil")
		return
	}

	s.gin.initialize()

	logger.Get().Info("initialize the core")
	// add gin.Engine to the core server
	s.core.initServer(s.gin.root)
}

// Load controller to the server, but it will not start the services
func (s *Server) Load(controller ...Interface.Controller) {
	s.core.appendControllers(controller...)

	for _, c := range controller {
		c.RequestGroup(s.gin.node)
	}
}

func (s *Server) PreLoadMiddleWare(name string, middleware gin.HandlerFunc) {
	s.gin.PreLoadMiddleWare(name, middleware)
}

func (s *Server) Run() {
	s.initialize()

	s.core.runServer()
}

func New(opts ...Option) *Server {
	s := &Server{
		core: newCore(opts...),
		gin:  newGinControl(),
	}
	return s
}
