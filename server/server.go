package server

import (
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/middleWare/Tracer"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
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

	fsWatcher *fsnotify.Watcher
	root      *Interface.RouteNode
	cfg       ginConfig
	g         *gin.Engine

	controllers []Interface.Controller
}

func (s *Server) initializeFileWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Get().Error(err)
		return
	}
	s.fsWatcher = watcher
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
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
		controller.BuildRoutes()
	}

	return nil
}

// Run start up the gin Server
func (s *Server) Run(addr string) {
	defer Config.CloseIni()

	/*	for _, controller := range s.controllers {
		controller.CreateNodeGroups()
	}*/

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

	if err := s.initializeRouter(); err != nil {
		return
	}

	s.initializeFileWatcher()
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
	config.AllowMethods = []string{"GET", "POST"}                              // 允许的请求方法
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"} // 允许的头部

	svr.g.Use(cors.New(config))

	for _, opt := range opts {
		opt(&svr.option)
	}
	return svr
}
