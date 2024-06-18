package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin  *gin.Engine
	core *core
}

func (s *Server) initiliaze() {
	s.core.setupServer(s.gin)
}

func (s *Server) Run() {
	if s.core == nil {
		s.core = newCore()
	}
	s.initiliaze()
	s.core.Run()

}

func (s *Server) Close() error {
	return nil
}

func New() *Server {
	return &Server{
		gin:  gin.Default(),
		core: newCore(),
	}
}
