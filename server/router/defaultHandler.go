package router

import "github.com/gin-gonic/gin"

func defaultFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func defaultHandler() *gin.HandlersChain {
	return &gin.HandlersChain{
		defaultFunc,
	}

}
