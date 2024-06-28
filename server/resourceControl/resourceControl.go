package resourceControl

import (
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type ResourceControl struct {
	cfg *config
}

func (rc *ResourceControl) Apply(engine *gin.Engine) {

	if rc.cfg.HtmlFiles != nil && len(rc.cfg.HtmlFiles) > 0 {
		//engine.LoadHTMLFiles(rc.cfg.HtmlFiles...)
		if len(rc.cfg.HtmlFiles) == 1 {
			engine.LoadHTMLGlob(rc.cfg.HtmlFiles[0])
		} else {
			engine.LoadHTMLFiles(rc.cfg.HtmlFiles...)
		}
	}

	/*	engine.GET("/", func(c *gin.Context) {
		c.File("dist/index.html")
	})*/

	if rc.cfg.NoMethod != "" {
		engine.NoMethod(func(c *gin.Context) {
			c.File(rc.cfg.NoMethod)
		})
	}

	if rc.cfg.NoRoute != "" {
		engine.NoRoute(func(c *gin.Context) {
			c.File(rc.cfg.NoRoute)
		})
	}

	if rc.cfg.ReMap != nil {
		for relativePath, root := range rc.cfg.ReMap {
			log.Infof("Mapping %s to %s", relativePath, root)
			engine.Static(relativePath, root)
		}
	}

	if rc.cfg.Redirects != nil && len(rc.cfg.Redirects) > 0 {
		for redirect, target := range rc.cfg.Redirects {
			log.Infof("Redirecting %s to %s", redirect, target)
			engine.GET(redirect, func(c *gin.Context) {
				c.HTML(200, "index.html", gin.H{})
			})
		}
	}

}

func New() *ResourceControl {
	ctrl := &ResourceControl{
		cfg: newConfig(),
	}
	return ctrl
}
