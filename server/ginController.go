package server

import (
	"encoding/json"
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/base/logger"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

type resourcePath struct {
	RelativePath string `json:"relativePath"`
	AbsolutePath string `json:"absolutePath"`
}

type ginControl struct {
	root *gin.Engine

	GroupMap       map[string]*Interface.GroupNode `json:"Group"`
	MiddleWarePool map[string]gin.HandlerFunc      `json:"-"`

	TemplatePath string         `json:"templatePath"`
	Statics      []resourcePath `json:"statics"`
	Redirect     []string       `json:"redirect"`
	NoRoute      string         `json:"noRoute"`
	NoMethod     string         `json:"noMethod"`
}

func (g *ginControl) applyShadow(shadow *ginControl) {

	for key, value := range shadow.GroupMap {
		logger.Get().Infof("apply config to [%s]", key)
		if g.GroupMap[key] == nil {
			logger.Get().Errorf("node [%s] not exists", key)
			continue
		}
		g.GroupMap[key].Parent = value.Parent
		g.GroupMap[key].MiddleWare = value.MiddleWare
	}

	g.TemplatePath = shadow.TemplatePath
	g.Statics = shadow.Statics
	g.Redirect = shadow.Redirect
	g.NoRoute = shadow.NoRoute
	g.NoMethod = shadow.NoMethod
}

func (g *ginControl) LoadConfig() {

	_, err := os.Stat("setting/ginConfig.json")
	if err != nil {
		g.Save()
		logger.Get().Error("config not exists ", err)
		return
	}
	raw, err := os.ReadFile("setting/ginConfig.json")
	if err != nil {
		return
	}

	shadow := &ginControl{}

	err = json.Unmarshal(raw, shadow)
	if err != nil {
		logger.Get().Error("LoadConfig error: ", err)
		return
	}
	g.applyShadow(shadow)
}

func (g *ginControl) Save() {

	bytes, err := json.MarshalIndent(g, "", "    ")
	if err != nil {
		logger.Get().Error("Save error: ", err)
		return
	}
	err = os.WriteFile("setting/ginConfig.json", bytes, 0644)
	if err != nil {
		logger.Get().Error("Save error: ", err)
		return
	}
}

func (g *ginControl) PreLoadMiddleWare(name string, middleware gin.HandlerFunc) {
	g.MiddleWarePool[name] = middleware
}

// controller can use this function to add a new node to the gin engine
func (g *ginControl) node(path string, parent string, middleWare ...string) *Interface.GroupNode {
	g.GroupMap[path] = &Interface.GroupNode{
		Parent:     parent,
		MiddleWare: middleWare,
	}
	return g.GroupMap[path]
}

func (g *ginControl) findParentNode(path string) *Interface.GroupNode {
	if g.GroupMap[path].Router == nil {
		parent := g.findParentNode(g.GroupMap[path].Parent) //recursively find the parent node

		//logger.Get().Debug("Start to build node : ", path, " with parent node : ", parent.Router.BasePath())

		curNode := g.GroupMap[path]
		curNode.Router = parent.Router.Group(path)
		logger.Get().Info("create group: ", curNode.Router.BasePath())
		if curNode.MiddleWare != nil {
			for _, mwName := range g.GroupMap[path].MiddleWare {
				logger.Get().Info("Add middleware to [", path, "] middleware name: ", mwName)
				g.GroupMap[path].Router.Use(g.MiddleWarePool[mwName])
			}
		}

	}
	return g.GroupMap[path]
}

// after this,all group's should be initialized
func (g *ginControl) initializeRouterGroups() {
	{
		logger.Get().Debug("Start to build root node")
		g.GroupMap["/"].Router = g.root.Group("/")
		for _, mwName := range g.GroupMap["/"].MiddleWare {
			logger.Get().Info("Add middleware to [", "/", "] middleware name: ", mwName)
			g.GroupMap["/"].Router.Use(g.MiddleWarePool[mwName])
		}
	}

	g.Save()

	for path, _ := range g.GroupMap {

		if g.GroupMap[path].Router != nil {
			continue
		} else {
			//logger.Get().Debug("Start to build node: ", path)
			g.findParentNode(path)

		}
	}
}

func (g *ginControl) initialize() {
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true // 允许所有来源
	g.root.Use(gin.Logger(), cors.New(cfg))
	g.node("/", "")
	g.LoadConfig()

	//initialize the root group

	g.initializeRouterGroups()

	if len(g.TemplatePath) > 0 {
		g.root.LoadHTMLGlob(g.TemplatePath)
	}
	if g.Statics != nil {
		for _, static := range g.Statics {
			g.root.Static(static.RelativePath, static.AbsolutePath)
		}
	}

	if g.Redirect != nil && len(g.Redirect) != 0 {
		for _, redirect := range g.Redirect {
			logger.Get().Infof("redirect path %s to index.html", redirect)
			g.root.GET(redirect, func(c *gin.Context) {
				c.HTML(200, "index.html", gin.H{})
			})
		}
	}

	if len(g.NoRoute) != 0 {
		g.root.NoRoute(func(c *gin.Context) {
			c.File(g.NoRoute)
		})
	}

	if len(g.NoMethod) == 0 {
		g.root.NoMethod(func(c *gin.Context) {
			c.File(g.NoMethod)
		})
	}

}

func setupGinMode() {
	mode := Config.GetBool("Server", "ReleaseMode")
	if mode {
		gin.SetMode(gin.ReleaseMode)
	} else if mode {
		Config.SetBool("Server", "ReleaseMode", false)
	}
}

func newGinControl() *ginControl {
	setupGinMode()
	return &ginControl{
		root:           gin.New(),
		GroupMap:       make(map[string]*Interface.GroupNode),
		MiddleWarePool: make(map[string]gin.HandlerFunc),
	}
}

//0xc0003e36e0
