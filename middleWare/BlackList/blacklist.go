package blacklist

import (
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

var blacklist = make(map[string]string)

func NewMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		if _, ok := blacklist[c.ClientIP()]; ok {
			log.Error("blacklist ip:" + c.ClientIP())
			c.Abort()
		}
	}
}
