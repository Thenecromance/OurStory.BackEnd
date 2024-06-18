package router

import "github.com/gin-gonic/gin"

//func defaultFunc(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"message": "pong",
//	})
//}
//
//func defaultHandler() *gin.HandlersChain {
//	return &gin.HandlersChain{
//		defaultFunc,
//	}
//
//}

func defaultFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "service is not available now.",
	})
}

func defaultHandler() gin.HandlerFunc {
	return defaultFunc
}
func defaultRESTHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		defaultFunc,
		defaultFunc,
		defaultFunc,
		defaultFunc,
	}
}

func defaultMiddleWare() gin.HandlersChain {
	return gin.HandlersChain{}
}
