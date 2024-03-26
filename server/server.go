package server

import (
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/middleWare/Tracer"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
)

var (
	funcMap template.FuncMap = template.FuncMap{}
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
	s.initialize()

	s.g.Run(addr)
}

func AppendFuncMap(functionMap template.FuncMap) {
	/*s.funcMap[key] = function*/
	for key, function := range functionMap {
		funcMap[key] = function
	}
}

func (s *Server) UpdateFuncMap() {
	s.g.SetFuncMap(funcMap)
	logger.Get().Info("===========UpdateFuncMap===========")
	for key, _ := range funcMap {
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
		g: gin.Default(),
	}
	svr.g.Use(
		Tracer.MiddleWare(),
		/*gJWT.NewMiddleware(
		gJWT.WithExpireTime(3600),
		gJWT.WithKey("")),*/
	)

	for _, opt := range opts {
		opt(&svr.option)
	}
	return svr
}
