package blacklist

import (
	"github.com/Thenecromance/OurStories/base/logger"
	"github.com/gin-gonic/gin"
)

var blacklist = make(map[string]string)

func NewMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		if _, ok := blacklist[c.ClientIP()]; ok {
			logger.Get().Error("blacklist ip:" + c.ClientIP())
			c.Abort()
		}
	}
}
