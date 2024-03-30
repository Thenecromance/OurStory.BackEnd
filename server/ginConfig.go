package server

import (
	"encoding/json"
	"github.com/Thenecromance/OurStories/base/logger"
	"github.com/gin-gonic/gin"
	"os"
)

type pair struct {
	RelativePath string `json:"relativePath"`
	AbsolutePath string `json:"absolutePath"`
}

type ginConfig struct {
	TemplatePath string   `json:"templatePath"`
	Statics      []pair   `json:"statics"`
	Redirect     []string `json:"redirect"`
	NoRoute      string   `json:"noRoute"`
	NoMethod     string   `json:"noMethod"`
}

const (
	configPath = "setting/gin.json"
)

/*func (cfg *ginConfig) defaultSetting() {
	cfg.TemplatePath = "dist/*.html"
	cfg.Statics = []pair{
		{
			RelativePath: "css",
			AbsolutePath: "dist/css",
		},
		{
			RelativePath: "js",
			AbsolutePath: "dist/js",
		},
		{
			RelativePath: "img",
			AbsolutePath: "dist/img",
		},
		{
			RelativePath: "fonts",
			AbsolutePath: "dist/fonts",
		},
		{
			RelativePath: "favicon.ico",
			AbsolutePath: "/favicon.ico",
		},
	}
	cfg.Redirect = []string{}
}*/

func (cfg *ginConfig) fileExists() bool {
	_, err := os.Stat(configPath)
	return err == nil
}

func (cfg *ginConfig) load() {
	logger.Get().Debug("load gin's config")

	if !cfg.fileExists() {
		logger.Get().Warnf("gin's config file not found, create a new one, setting it by yourself")

		cfg.save()
		return
	}

	buf, err := os.ReadFile(configPath)
	if err != nil {
		logger.Get().Info(err)
		return
	}

	err = json.Unmarshal(buf, cfg)
	if err != nil {
		return
	}
	cfg.save()

}

func (cfg *ginConfig) save() {
	buf, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return
	}

	err = os.WriteFile(configPath, buf, 0644)
	if err != nil {
		return
	}
}

// apply will load the template path, statics, redirect, noRoute, and noMethod to the gin engine
func (cfg *ginConfig) apply(_gin *gin.Engine) {
	if len(cfg.TemplatePath) > 0 {
		_gin.LoadHTMLGlob(cfg.TemplatePath)
	}
	if cfg.Statics != nil {
		for _, static := range cfg.Statics {
			_gin.Static(static.RelativePath, static.AbsolutePath)
		}
	}

	if cfg.Redirect != nil && len(cfg.Redirect) != 0 {
		for _, redirect := range cfg.Redirect {
			logger.Get().Infof("redirect path %s to index.html", redirect)
			_gin.GET(redirect, func(c *gin.Context) {
				c.HTML(200, "index.html", gin.H{})
			})
		}
	}

	if len(cfg.NoRoute) != 0 {
		_gin.NoRoute(func(c *gin.Context) {
			c.File(cfg.NoRoute)
		})
	}

	if len(cfg.NoMethod) == 0 {
		_gin.NoMethod(func(c *gin.Context) {
			c.File(cfg.NoMethod)
		})
	}

}
