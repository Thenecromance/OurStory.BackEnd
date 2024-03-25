package server

import (
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
)

func init() {
	os.Mkdir("setting", 0755)
}

type Server struct {
	option ServerOption

	cfg ginConfig

	g    *gin.Engine
	root *gin.RouterGroup

	controllers []Interface.Controller

	funcMap template.FuncMap
}

// Gin returns the gin engine
func (s *Server) Gin() *gin.Engine {
	return s.g
}

func (s *Server) Group() *gin.RouterGroup {
	if s.root == nil {
		s.root = s.g.Group("/")
	}
	return s.root
}

// Run start up the gin Server
func (s *Server) Run(addr string) {
	defer Config.CloseIni()

	for _, controller := range s.controllers {
		controller.BuildRoutes()
	}

	s.UpdateFuncMap()

	s.g.Run(addr)
}

func (s *Server) AppendFuncMap(functionMap template.FuncMap) {
	/*s.funcMap[key] = function*/
	for key, function := range functionMap {
		s.funcMap[key] = function
	}
}

func (s *Server) UpdateFuncMap() {
	s.g.SetFuncMap(s.funcMap)
	logger.Get().Info("===========UpdateFuncMap===========")
	for key, _ := range s.funcMap {
		logger.Get().Info(key)
	}
	logger.Get().Info("===========UpdateFuncMap===========")
}

func (s *Server) LoadComponent(controller ...Interface.Controller) {
	s.controllers = append(s.controllers, controller...)
}

func (s *Server) initialize() {
	s.cfg.load()
	s.cfg.apply(s.g)
}

func New(opts ...Option) *Server {

	svr := &Server{
		g:       gin.Default(),
		funcMap: template.FuncMap{},
	}
	svr.initialize()

	for _, opt := range opts {
		opt(&svr.option)
	}
	return svr
}
