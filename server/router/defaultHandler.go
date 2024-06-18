package router

import (
	"github.com/Thenecromance/OurStories/server/response"
	"github.com/gin-gonic/gin"
)

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
	resp := response.New()
	defer response.Send(c, resp)

	resp.Code = response.NotAcceptable
	resp.Meta.Count = 0
	resp.Data = gin.H{
		"system": "service not found",
	}
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
