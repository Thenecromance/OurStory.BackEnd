package resources

import (
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	cfg *config
}

func (rc *Controller) ApplyTo(engine *gin.Engine) {
	log.Info("Start to apply resources to gin engine")
	if rc.cfg.HtmlFiles != nil && len(rc.cfg.HtmlFiles) > 0 {
		//engine.LoadHTMLFiles(rc.cfg.HtmlFiles...)
		if len(rc.cfg.HtmlFiles) == 1 {
			engine.LoadHTMLGlob(rc.cfg.HtmlFiles[0])
		} else {
			engine.LoadHTMLFiles(rc.cfg.HtmlFiles...)
		}
	}

	if rc.cfg.NoMethod != "" {
		log.Infof("Setting NoMethod to %s", rc.cfg.NoMethod)
		engine.NoMethod(func(c *gin.Context) {
			c.File(rc.cfg.NoMethod)
		})
	}

	if rc.cfg.NoRoute != "" {
		log.Infof("Setting NoRoute to %s", rc.cfg.NoRoute)
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

	log.Info("Resources applied to gin engine")
}

func New() *Controller {
	ctrl := &Controller{
		cfg: newConfig(),
	}
	return ctrl
}
