package server

import (
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/middleWare/TLS"
	"github.com/Thenecromance/OurStories/middleWare/Tracer"
	"github.com/gin-contrib/cors"
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

	root *Interface.RouteNode
	cfg  ginConfig
	g    *gin.Engine

	controllers []Interface.Controller
}

func (s *Server) initializeRouter() error {
	// load the routes node first. in this situation all the controllers didn't has any relationship between each other
	for _, controller := range s.controllers {
		s.root.Load(controller.GetNode())
	}
	// then based on the multi tree structure, we will build the tree
	err := s.root.MakeAsTree()
	if err != nil {
		logger.Get().Error(err)
		return err
	}
	// after the tree is built, we will build the routes group for each nodes(controllers)
	s.root.CreateNodeGroups()
	// finally we will build the routes for each controller
	for _, controller := range s.controllers {
		controller.ApplyMiddleWare()
		controller.BuildRoutes()
	}

	return nil
}

// Run start up the gin Server
func (s *Server) Run(addr string) {
	defer Config.CloseIni()
	s.UpdateFuncMap()
	s.initialize()

	if s.option.TLS {
		logger.Get().Info("server is running on TLS , cert file is ", s.option.CertFile, " key file is ", s.option.KeyFile, " address is ", addr)
		logger.Get().Info("please make sure the cert file and key file is correct and the address is correct")
		s.g.RunTLS(addr, s.option.CertFile, s.option.KeyFile)
	} else {
		logger.Get().Info("server is running on ", addr)
		s.g.Run(addr)
	}

}

func (s *Server) UpdateFuncMap() {
	s.g.SetFuncMap(funcMap)
	//logger.Get().Info("===========UpdateFuncMap===========")
	for key, _ := range funcMap {
		logger.Get().Info(key)
	}
	//logger.Get().Info("===========UpdateFuncMap===========")
}

func (s *Server) LoadComponent(controller ...Interface.Controller) {
	s.controllers = append(s.controllers, controller...)
}

func (s *Server) initialize() {
	s.cfg.load()
	s.cfg.apply(s.g)
	s.root = Interface.NewRootNode()
	s.root.RouterGroup = s.g.Group("/") // set up the root group as "/"

	//build each node's group
	if err := s.initializeRouter(); err != nil {
		return
	}

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
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                              // 允许所有来源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}             // 允许的请求方法
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"} // 允许的头部

	svr.g.Use(TLS.TlsHandler(8080), cors.New(config))

	for _, opt := range opts {
		opt(&svr.option)
	}
	return svr
}
