package server

import (
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin  *gin.Engine
	core *core
}

func (s *Server) initiliaze() {

}

func (s *Server) RegisterRouter(routerProxy Interface.RouterProxy) error {
	return s.core.RegisterRouter(routerProxy)
}

func (s *Server) Run() {
	s.initiliaze()
	s.core.Run()
	
}

func (s *Server) Close() error {
	return nil
}

func New() *Server {
	return &Server{
		gin: gin.Default(),
	}
}
